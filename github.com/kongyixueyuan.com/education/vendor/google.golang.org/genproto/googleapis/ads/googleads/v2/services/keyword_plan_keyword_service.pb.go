// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/ads/googleads/v2/services/keyword_plan_keyword_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for [KeywordPlanKeywordService.GetKeywordPlanKeyword][google.ads.googleads.v2.services.KeywordPlanKeywordService.GetKeywordPlanKeyword].
type GetKeywordPlanKeywordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the ad group keyword to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetKeywordPlanKeywordRequest) Reset() {
	*x = GetKeywordPlanKeywordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeywordPlanKeywordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeywordPlanKeywordRequest) ProtoMessage() {}

func (x *GetKeywordPlanKeywordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeywordPlanKeywordRequest.ProtoReflect.Descriptor instead.
func (*GetKeywordPlanKeywordRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetKeywordPlanKeywordRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for [KeywordPlanKeywordService.MutateKeywordPlanKeywords][google.ads.googleads.v2.services.KeywordPlanKeywordService.MutateKeywordPlanKeywords].
type MutateKeywordPlanKeywordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose Keyword Plan keywords are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual Keyword Plan keywords.
	Operations []*KeywordPlanKeywordOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateKeywordPlanKeywordsRequest) Reset() {
	*x = MutateKeywordPlanKeywordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateKeywordPlanKeywordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateKeywordPlanKeywordsRequest) ProtoMessage() {}

func (x *MutateKeywordPlanKeywordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateKeywordPlanKeywordsRequest.ProtoReflect.Descriptor instead.
func (*MutateKeywordPlanKeywordsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateKeywordPlanKeywordsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateKeywordPlanKeywordsRequest) GetOperations() []*KeywordPlanKeywordOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateKeywordPlanKeywordsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateKeywordPlanKeywordsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a Keyword Plan keyword.
type KeywordPlanKeywordOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The FieldMask that determines which resource fields are modified in an
	// update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*KeywordPlanKeywordOperation_Create
	//	*KeywordPlanKeywordOperation_Update
	//	*KeywordPlanKeywordOperation_Remove
	Operation isKeywordPlanKeywordOperation_Operation `protobuf_oneof:"operation"`
}

func (x *KeywordPlanKeywordOperation) Reset() {
	*x = KeywordPlanKeywordOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanKeywordOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanKeywordOperation) ProtoMessage() {}

func (x *KeywordPlanKeywordOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanKeywordOperation.ProtoReflect.Descriptor instead.
func (*KeywordPlanKeywordOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescGZIP(), []int{2}
}

func (x *KeywordPlanKeywordOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *KeywordPlanKeywordOperation) GetOperation() isKeywordPlanKeywordOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *KeywordPlanKeywordOperation) GetCreate() *resources.KeywordPlanKeyword {
	if x, ok := x.GetOperation().(*KeywordPlanKeywordOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *KeywordPlanKeywordOperation) GetUpdate() *resources.KeywordPlanKeyword {
	if x, ok := x.GetOperation().(*KeywordPlanKeywordOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *KeywordPlanKeywordOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*KeywordPlanKeywordOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isKeywordPlanKeywordOperation_Operation interface {
	isKeywordPlanKeywordOperation_Operation()
}

type KeywordPlanKeywordOperation_Create struct {
	// Create operation: No resource name is expected for the new Keyword Plan
	// ad group keyword.
	Create *resources.KeywordPlanKeyword `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanKeywordOperation_Update struct {
	// Update operation: The Keyword Plan keyword is expected to have a valid
	// resource name.
	Update *resources.KeywordPlanKeyword `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanKeywordOperation_Remove struct {
	// Remove operation: A resource name for the removed Keyword Plan keyword is
	// expected, in this format:
	//
	// `customers/{customer_id}/keywordPlanKeywords/{kp_ad_group_keyword_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanKeywordOperation_Create) isKeywordPlanKeywordOperation_Operation() {}

func (*KeywordPlanKeywordOperation_Update) isKeywordPlanKeywordOperation_Operation() {}

func (*KeywordPlanKeywordOperation_Remove) isKeywordPlanKeywordOperation_Operation() {}

// Response message for a Keyword Plan keyword mutate.
type MutateKeywordPlanKeywordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateKeywordPlanKeywordResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateKeywordPlanKeywordsResponse) Reset() {
	*x = MutateKeywordPlanKeywordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateKeywordPlanKeywordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateKeywordPlanKeywordsResponse) ProtoMessage() {}

func (x *MutateKeywordPlanKeywordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateKeywordPlanKeywordsResponse.ProtoReflect.Descriptor instead.
func (*MutateKeywordPlanKeywordsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateKeywordPlanKeywordsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateKeywordPlanKeywordsResponse) GetResults() []*MutateKeywordPlanKeywordResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the Keyword Plan keyword mutate.
type MutateKeywordPlanKeywordResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateKeywordPlanKeywordResult) Reset() {
	*x = MutateKeywordPlanKeywordResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateKeywordPlanKeywordResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateKeywordPlanKeywordResult) ProtoMessage() {}

func (x *MutateKeywordPlanKeywordResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateKeywordPlanKeywordResult.ProtoReflect.Descriptor instead.
func (*MutateKeywordPlanKeywordResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateKeywordPlanKeywordResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xfa, 0x01,
	0x0a, 0x20, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50,
	0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x62, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xa3, 0x02, 0x0a, 0x1b, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x4f, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xc7, 0x01, 0x0a, 0x21, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5a,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x1e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0x9d, 0x04, 0x0a, 0x19, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61,
	0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xdd, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c,
	0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x12, 0x35, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x2a, 0x7d, 0xda,
	0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x82, 0x02, 0x0a, 0x19, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x42, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c,
	0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x22, 0x38,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1b, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x42, 0x85, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1e, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x20, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca,
	0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescData = file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDesc
)

func file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDescData
}

var file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_goTypes = []interface{}{
	(*GetKeywordPlanKeywordRequest)(nil),      // 0: google.ads.googleads.v2.services.GetKeywordPlanKeywordRequest
	(*MutateKeywordPlanKeywordsRequest)(nil),  // 1: google.ads.googleads.v2.services.MutateKeywordPlanKeywordsRequest
	(*KeywordPlanKeywordOperation)(nil),       // 2: google.ads.googleads.v2.services.KeywordPlanKeywordOperation
	(*MutateKeywordPlanKeywordsResponse)(nil), // 3: google.ads.googleads.v2.services.MutateKeywordPlanKeywordsResponse
	(*MutateKeywordPlanKeywordResult)(nil),    // 4: google.ads.googleads.v2.services.MutateKeywordPlanKeywordResult
	(*fieldmaskpb.FieldMask)(nil),             // 5: google.protobuf.FieldMask
	(*resources.KeywordPlanKeyword)(nil),      // 6: google.ads.googleads.v2.resources.KeywordPlanKeyword
	(*status.Status)(nil),                     // 7: google.rpc.Status
}
var file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v2.services.MutateKeywordPlanKeywordsRequest.operations:type_name -> google.ads.googleads.v2.services.KeywordPlanKeywordOperation
	5, // 1: google.ads.googleads.v2.services.KeywordPlanKeywordOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 2: google.ads.googleads.v2.services.KeywordPlanKeywordOperation.create:type_name -> google.ads.googleads.v2.resources.KeywordPlanKeyword
	6, // 3: google.ads.googleads.v2.services.KeywordPlanKeywordOperation.update:type_name -> google.ads.googleads.v2.resources.KeywordPlanKeyword
	7, // 4: google.ads.googleads.v2.services.MutateKeywordPlanKeywordsResponse.partial_failure_error:type_name -> google.rpc.Status
	4, // 5: google.ads.googleads.v2.services.MutateKeywordPlanKeywordsResponse.results:type_name -> google.ads.googleads.v2.services.MutateKeywordPlanKeywordResult
	0, // 6: google.ads.googleads.v2.services.KeywordPlanKeywordService.GetKeywordPlanKeyword:input_type -> google.ads.googleads.v2.services.GetKeywordPlanKeywordRequest
	1, // 7: google.ads.googleads.v2.services.KeywordPlanKeywordService.MutateKeywordPlanKeywords:input_type -> google.ads.googleads.v2.services.MutateKeywordPlanKeywordsRequest
	6, // 8: google.ads.googleads.v2.services.KeywordPlanKeywordService.GetKeywordPlanKeyword:output_type -> google.ads.googleads.v2.resources.KeywordPlanKeyword
	3, // 9: google.ads.googleads.v2.services.KeywordPlanKeywordService.MutateKeywordPlanKeywords:output_type -> google.ads.googleads.v2.services.MutateKeywordPlanKeywordsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_init() }
func file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_init() {
	if File_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeywordPlanKeywordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateKeywordPlanKeywordsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPlanKeywordOperation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateKeywordPlanKeywordsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateKeywordPlanKeywordResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*KeywordPlanKeywordOperation_Create)(nil),
		(*KeywordPlanKeywordOperation_Update)(nil),
		(*KeywordPlanKeywordOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto = out.File
	file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_rawDesc = nil
	file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_goTypes = nil
	file_google_ads_googleads_v2_services_keyword_plan_keyword_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordPlanKeywordServiceClient is the client API for KeywordPlanKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanKeywordServiceClient interface {
	// Returns the requested Keyword Plan keyword in full detail.
	GetKeywordPlanKeyword(ctx context.Context, in *GetKeywordPlanKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanKeyword, error)
	// Creates, updates, or removes Keyword Plan keywords. Operation statuses are
	// returned.
	MutateKeywordPlanKeywords(ctx context.Context, in *MutateKeywordPlanKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanKeywordsResponse, error)
}

type keywordPlanKeywordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanKeywordServiceClient(cc grpc.ClientConnInterface) KeywordPlanKeywordServiceClient {
	return &keywordPlanKeywordServiceClient{cc}
}

func (c *keywordPlanKeywordServiceClient) GetKeywordPlanKeyword(ctx context.Context, in *GetKeywordPlanKeywordRequest, opts ...grpc.CallOption) (*resources.KeywordPlanKeyword, error) {
	out := new(resources.KeywordPlanKeyword)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordPlanKeywordService/GetKeywordPlanKeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanKeywordServiceClient) MutateKeywordPlanKeywords(ctx context.Context, in *MutateKeywordPlanKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanKeywordsResponse, error) {
	out := new(MutateKeywordPlanKeywordsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordPlanKeywordService/MutateKeywordPlanKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanKeywordServiceServer is the server API for KeywordPlanKeywordService service.
type KeywordPlanKeywordServiceServer interface {
	// Returns the requested Keyword Plan keyword in full detail.
	GetKeywordPlanKeyword(context.Context, *GetKeywordPlanKeywordRequest) (*resources.KeywordPlanKeyword, error)
	// Creates, updates, or removes Keyword Plan keywords. Operation statuses are
	// returned.
	MutateKeywordPlanKeywords(context.Context, *MutateKeywordPlanKeywordsRequest) (*MutateKeywordPlanKeywordsResponse, error)
}

// UnimplementedKeywordPlanKeywordServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanKeywordServiceServer struct {
}

func (*UnimplementedKeywordPlanKeywordServiceServer) GetKeywordPlanKeyword(context.Context, *GetKeywordPlanKeywordRequest) (*resources.KeywordPlanKeyword, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetKeywordPlanKeyword not implemented")
}
func (*UnimplementedKeywordPlanKeywordServiceServer) MutateKeywordPlanKeywords(context.Context, *MutateKeywordPlanKeywordsRequest) (*MutateKeywordPlanKeywordsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateKeywordPlanKeywords not implemented")
}

func RegisterKeywordPlanKeywordServiceServer(s *grpc.Server, srv KeywordPlanKeywordServiceServer) {
	s.RegisterService(&_KeywordPlanKeywordService_serviceDesc, srv)
}

func _KeywordPlanKeywordService_GetKeywordPlanKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanKeywordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanKeywordServiceServer).GetKeywordPlanKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordPlanKeywordService/GetKeywordPlanKeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanKeywordServiceServer).GetKeywordPlanKeyword(ctx, req.(*GetKeywordPlanKeywordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanKeywordService_MutateKeywordPlanKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanKeywordServiceServer).MutateKeywordPlanKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordPlanKeywordService/MutateKeywordPlanKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanKeywordServiceServer).MutateKeywordPlanKeywords(ctx, req.(*MutateKeywordPlanKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanKeywordService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.KeywordPlanKeywordService",
	HandlerType: (*KeywordPlanKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanKeyword",
			Handler:    _KeywordPlanKeywordService_GetKeywordPlanKeyword_Handler,
		},
		{
			MethodName: "MutateKeywordPlanKeywords",
			Handler:    _KeywordPlanKeywordService_MutateKeywordPlanKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/keyword_plan_keyword_service.proto",
}
