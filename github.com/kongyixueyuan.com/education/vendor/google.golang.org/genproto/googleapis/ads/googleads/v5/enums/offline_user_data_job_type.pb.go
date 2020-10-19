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
// source: google/ads/googleads/v5/enums/offline_user_data_job_type.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// The type of an offline user data job.
type OfflineUserDataJobTypeEnum_OfflineUserDataJobType int32

const (
	// Not specified.
	OfflineUserDataJobTypeEnum_UNSPECIFIED OfflineUserDataJobTypeEnum_OfflineUserDataJobType = 0
	// Used for return value only. Represents value unknown in this version.
	OfflineUserDataJobTypeEnum_UNKNOWN OfflineUserDataJobTypeEnum_OfflineUserDataJobType = 1
	// Store Sales Direct data for self service.
	OfflineUserDataJobTypeEnum_STORE_SALES_UPLOAD_FIRST_PARTY OfflineUserDataJobTypeEnum_OfflineUserDataJobType = 2
	// Store Sales Direct data for third party.
	OfflineUserDataJobTypeEnum_STORE_SALES_UPLOAD_THIRD_PARTY OfflineUserDataJobTypeEnum_OfflineUserDataJobType = 3
	// Customer Match user list data.
	OfflineUserDataJobTypeEnum_CUSTOMER_MATCH_USER_LIST OfflineUserDataJobTypeEnum_OfflineUserDataJobType = 4
)

// Enum value maps for OfflineUserDataJobTypeEnum_OfflineUserDataJobType.
var (
	OfflineUserDataJobTypeEnum_OfflineUserDataJobType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "STORE_SALES_UPLOAD_FIRST_PARTY",
		3: "STORE_SALES_UPLOAD_THIRD_PARTY",
		4: "CUSTOMER_MATCH_USER_LIST",
	}
	OfflineUserDataJobTypeEnum_OfflineUserDataJobType_value = map[string]int32{
		"UNSPECIFIED":                    0,
		"UNKNOWN":                        1,
		"STORE_SALES_UPLOAD_FIRST_PARTY": 2,
		"STORE_SALES_UPLOAD_THIRD_PARTY": 3,
		"CUSTOMER_MATCH_USER_LIST":       4,
	}
)

func (x OfflineUserDataJobTypeEnum_OfflineUserDataJobType) Enum() *OfflineUserDataJobTypeEnum_OfflineUserDataJobType {
	p := new(OfflineUserDataJobTypeEnum_OfflineUserDataJobType)
	*p = x
	return p
}

func (x OfflineUserDataJobTypeEnum_OfflineUserDataJobType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OfflineUserDataJobTypeEnum_OfflineUserDataJobType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_enumTypes[0].Descriptor()
}

func (OfflineUserDataJobTypeEnum_OfflineUserDataJobType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_enumTypes[0]
}

func (x OfflineUserDataJobTypeEnum_OfflineUserDataJobType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OfflineUserDataJobTypeEnum_OfflineUserDataJobType.Descriptor instead.
func (OfflineUserDataJobTypeEnum_OfflineUserDataJobType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing types of an offline user data job.
type OfflineUserDataJobTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OfflineUserDataJobTypeEnum) Reset() {
	*x = OfflineUserDataJobTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineUserDataJobTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineUserDataJobTypeEnum) ProtoMessage() {}

func (x *OfflineUserDataJobTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineUserDataJobTypeEnum.ProtoReflect.Descriptor instead.
func (*OfflineUserDataJobTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v5_enums_offline_user_data_job_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01,
	0x0a, 0x1a, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x9c, 0x01, 0x0a,
	0x16, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53,
	0x41, 0x4c, 0x45, 0x53, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x49, 0x52, 0x53,
	0x54, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x54, 0x4f,
	0x52, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f,
	0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x10, 0x03, 0x12, 0x1c, 0x0a,
	0x18, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x04, 0x42, 0xf0, 0x01, 0x0a, 0x21,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x1b, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x35, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x35, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x35, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescData = file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDesc
)

func file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDescData
}

var file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_goTypes = []interface{}{
	(OfflineUserDataJobTypeEnum_OfflineUserDataJobType)(0), // 0: google.ads.googleads.v5.enums.OfflineUserDataJobTypeEnum.OfflineUserDataJobType
	(*OfflineUserDataJobTypeEnum)(nil),                     // 1: google.ads.googleads.v5.enums.OfflineUserDataJobTypeEnum
}
var file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_init() }
func file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_init() {
	if File_google_ads_googleads_v5_enums_offline_user_data_job_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineUserDataJobTypeEnum); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v5_enums_offline_user_data_job_type_proto = out.File
	file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_rawDesc = nil
	file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_goTypes = nil
	file_google_ads_googleads_v5_enums_offline_user_data_job_type_proto_depIdxs = nil
}
