package lints

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

import (
	"testing"
)

func TestEtsiTypeAsQcStmt(t *testing.T) {
	m := map[string]LintStatus{
		"QcStmtEtsiQcTypeAsQcStmtCert10.pem": Error,
		"QcStmtEtsiTaggedValueCert20.pem":    Error,
		"QcStmtEtsiValidCert03.pem":          Pass,
		"QcStmtEtsiEsealValidCert02.pem":     Pass,
		"QcStmtEtsiTwoQcTypesCert15.pem":     Pass,
		"QcStmtEtsiNoQcStatmentsCert22.pem":  NA,
		"QcStmtEtsiValidCert24.pem":          Pass,
	}
	for inputPath, expected := range m {
		inputPath = "../testlint/testCerts/" + inputPath
		out := Lints["e_qcstatem_etsi_type_as_statem"].Execute(ReadCertificate(inputPath))

		if out.Status != expected {
			t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
		}
	}
}
