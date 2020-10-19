/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zmap/zlint/util"
)

const (
	// ICANN_GTLD_JSON is the URL for the ICANN gTLD JSON registry (version 2).
	// This registry does not contain ccTLDs but does carry full gTLD information
	// needed to determine validity periods.
	// See https://www.icann.org/resources/pages/registries/registries-en for more
	// information.
	ICANN_GTLD_JSON = "https://www.icann.org/resources/registries/gtlds/v2/gtlds.json"
	// ICANN_TLDS is the URL for the ICANN list of valid top-level domains
	// maintained by the IANA. It contains both ccTLDs and gTLDs but does not
	// carry sufficient granularity to determine validity periods.
	// See https://www.icann.org/resources/pages/tlds-2012-02-25-en for more
	// information.
	ICANN_TLDS = "https://data.iana.org/TLD/tlds-alpha-by-domain.txt"
)

var (
	// httpClient is a http.Client instance configured with timeouts.
	httpClient = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   15 * time.Second,
				KeepAlive: 15 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	// gTLDMapTemplate is a template that produces a Golang source code file in
	// the "util" package containing a single member variable, a map of strings to
	// `util.GTLDPeriod` objects called `tldMap`.
	gTLDMapTemplate = template.Must(template.New("gTLDMapTemplate").Parse(
		`// Code generated by go generate; DO NOT EDIT.
// This file was generated by zlint-gtld-update.

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package util

var tldMap = map[string]GTLDPeriod{
{{- range .GTLDs }}
	"{{ .GTLD }}": {
		GTLD: "{{ .GTLD }}",
		DelegationDate: "{{ .DelegationDate }}",
		RemovalDate: "{{ .RemovalDate }}",
	},
{{- end }}
	// .onion is a special case and not a general gTLD. However, it is allowed in
	// some circumstances in the web PKI so the Zlint gtldMap includes it with
	// a delegationDate based on the CABF ballot to allow EV issuance for .onion
	// domains: https://cabforum.org/2015/02/18/ballot-144-validation-rules-dot-onion-names/
	"onion": {
		GTLD: "onion",
		DelegationDate: "2015-02-18",
		RemovalDate: "",
	},
}
`))
)

// getData fetches the response body bytes from an HTTP get to the provider url,
// or returns an error.
func getData(url string) ([]byte, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch data from %q : %s",
			url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code fetching data "+
			"from %q : expected status %d got %d",
			url, http.StatusOK, resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unexpected error reading response "+
			"body from %q : %s",
			url, err)
	}
	return respBody, nil
}

// getTLDData fetches the ICANN_TLDS list and uses the information to build
// and return a list of util.GTLDPeriod objects (or an error if anything fails).
// Since this data source only contains TLD names and not any information
// about delegation/removal all of the returned `util.GTLDPeriod` objects will
// have the DelegationDate "1985-01-01" (matching the `.com` delegation date)
// and no RemovalDate.
func getTLDData() ([]util.GTLDPeriod, error) {
	respBody, err := getData(ICANN_TLDS)
	if err != nil {
		return nil, fmt.Errorf("error getting ICANN TLD list : %s", err)
	}
	tlds := strings.Split(string(respBody), "\n")

	var results []util.GTLDPeriod
	for _, tld := range tlds {
		// Skip empty lines and the header comment line
		if strings.TrimSpace(tld) == "" || strings.HasPrefix(tld, "#") {
			continue
		}
		results = append(results, util.GTLDPeriod{
			GTLD: strings.ToLower(tld),
			// The TLD list doesn't indicate when any of the TLDs were delegated so
			// assume these TLDs were all delegated at the same time as "com".
			DelegationDate: "1985-01-01",
		})
	}
	return results, nil
}

// getGTLDData fetches the ICANN_GTLD_JSON and parses it into a list of
// util.GTLDPeriod objects, or returns an error. The gTLDEntries are returned
// as-is and may contain entries that were never delegated from the root DNS.
func getGTLDData() ([]util.GTLDPeriod, error) {
	respBody, err := getData(ICANN_GTLD_JSON)
	if err != nil {
		return nil, fmt.Errorf("error getting ICANN gTLD JSON : %s", err)
	}

	var results struct {
		GTLDs []util.GTLDPeriod
	}
	if err := json.Unmarshal(respBody, &results); err != nil {
		return nil, fmt.Errorf("unexpected error unmarshaling ICANN gTLD JSON response "+
			"body from %q : %s",
			ICANN_GTLD_JSON, err)
	}
	return results.GTLDs, nil
}

// delegatedGTLDs filters the provided list of GTLDPeriods removing any entries
// that were never delegated from the root DNS.
func delegatedGTLDs(entries []util.GTLDPeriod) []util.GTLDPeriod {
	var results []util.GTLDPeriod
	for _, gTLD := range entries {
		if gTLD.DelegationDate == "" {
			continue
		}
		results = append(results, gTLD)
	}
	return results
}

// validateGTLDs checks that all entries have a valid parseable DelegationDate
// string, and if not-empty, a valid parseable RemovalDate string. This function
// assumes an entry with an empty DelegationDate is an error. Use
// `delegatedGTLDs` to filter out entries that were never delegated before
// validating.
func validateGTLDs(entries []util.GTLDPeriod) error {
	for _, gTLD := range entries {
		// All entries should have a valid delegation date
		if _, err := time.Parse(util.GTLDPeriodDateFormat, gTLD.DelegationDate); err != nil {
			return err
		}
		// a gTLD that has not been removed has an empty RemovalDate and that's OK
		if _, err := time.Parse(util.GTLDPeriodDateFormat, gTLD.RemovalDate); gTLD.RemovalDate != "" && err != nil {
			return err
		}
	}
	return nil
}

// renderGTLDMap fetches the ICANN gTLD data, filters out undelegated entries,
// validates the remaining entries have parseable dates, and renders the
// gTLDMapTemplate to the provided writer using the validated entries (or
// returns an error if any of the aforementioned steps fail). It then fetches
// the ICANN TLD data, and uses it to populate any missing entries for ccTLDs.
// These entries will have a default delegationDate because the data source is
// not specific enough to provide one. The produced output text is a Golang
// source code file in the `util` package that contains a single map variable
// containing GTLDPeriod objects created with the ICANN data.
func renderGTLDMap(writer io.Writer) error {
	// Get all of ICANN's gTLDs including ones that haven't been delegated.
	allGTLDs, err := getGTLDData()
	if err != nil {
		return err
	}

	// Filter out the non-delegated gTLD entries
	delegatedGTLDs := delegatedGTLDs(allGTLDs)

	// Validate that all of the delegated gTLDs have correct dates
	if err := validateGTLDs(delegatedGTLDs); err != nil {
		return err
	}

	// Get all of the TLDs. This data source doesn't provide delegationDates and
	// so we only want to use it to populate missing entries in `delegatedGTLDs`,
	// not to replace any existing entries that have more specific information
	// about the validity period for the TLD.
	allTLDs, err := getTLDData()
	if err != nil {
		return err
	}

	tldMap := make(map[string]util.GTLDPeriod)

	// Deduplicate delegatedGTLDs into the tldMap first
	for _, tld := range delegatedGTLDs {
		tldMap[tld.GTLD] = tld
	}

	// Then populate any missing entries from the allTLDs list
	for _, tld := range allTLDs {
		if _, found := tldMap[tld.GTLD]; !found {
			tldMap[tld.GTLD] = tld
		}
	}

	templateData := struct {
		GTLDs map[string]util.GTLDPeriod
	}{
		GTLDs: tldMap,
	}

	// Render the gTLD map to a buffer with the delegated gTLD data
	var buf bytes.Buffer
	if err := gTLDMapTemplate.Execute(&buf, templateData); err != nil {
		return err
	}

	// format the buffer so it won't trip up the `gofmt_test.go` checks
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	// Write the formatted buffer to the writer
	_, err = writer.Write(formatted)
	if err != nil {
		return err
	}
	return nil
}

// init sets up command line flags
func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	log.SetLevel(log.InfoLevel)
}

// main handles rendering a gTLD map to either standard out (when no argument is
// provided) or to the provided filename. If an error occurs it is printed to
// standard err and the program terminates with a non-zero exit status.
func main() {
	errQuit := func(err error) {
		fmt.Fprintf(os.Stderr, "error updating gTLD map: %s\n", err)
		os.Exit(1)
	}

	// Default to writing to standard out
	writer := os.Stdout
	if flag.NArg() > 0 {
		// If a filename is specified as a command line flag then open it (creating
		// if needed), truncate the existing contents, and use the file as the
		// writer instead of standard out
		filename := flag.Args()[0]
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0664)
		if err != nil {
			errQuit(err)
		}
		defer f.Close()
		writer = f
	}

	if err := renderGTLDMap(writer); err != nil {
		errQuit(err)
	}
}
