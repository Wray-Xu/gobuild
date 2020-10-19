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
// source: google/ads/googleads/v5/services/video_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v5/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [VideoService.GetVideo][google.ads.googleads.v5.services.VideoService.GetVideo].
type GetVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the video to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetVideoRequest) Reset() {
	*x = GetVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v5_services_video_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoRequest) ProtoMessage() {}

func (x *GetVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v5_services_video_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoRequest.ProtoReflect.Descriptor instead.
func (*GetVideoRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v5_services_video_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetVideoRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v5_services_video_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v5_services_video_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x35, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x35, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xd7, 0x01, 0x0a, 0x0c, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa9, 0x01, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f,
	0x76, 0x35, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x1b, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x42, 0xf8, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x11, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x35, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x35, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x35, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v5_services_video_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v5_services_video_service_proto_rawDescData = file_google_ads_googleads_v5_services_video_service_proto_rawDesc
)

func file_google_ads_googleads_v5_services_video_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v5_services_video_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v5_services_video_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v5_services_video_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v5_services_video_service_proto_rawDescData
}

var file_google_ads_googleads_v5_services_video_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v5_services_video_service_proto_goTypes = []interface{}{
	(*GetVideoRequest)(nil), // 0: google.ads.googleads.v5.services.GetVideoRequest
	(*resources.Video)(nil), // 1: google.ads.googleads.v5.resources.Video
}
var file_google_ads_googleads_v5_services_video_service_proto_depIdxs = []int32{
	0, // 0: google.ads.googleads.v5.services.VideoService.GetVideo:input_type -> google.ads.googleads.v5.services.GetVideoRequest
	1, // 1: google.ads.googleads.v5.services.VideoService.GetVideo:output_type -> google.ads.googleads.v5.resources.Video
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v5_services_video_service_proto_init() }
func file_google_ads_googleads_v5_services_video_service_proto_init() {
	if File_google_ads_googleads_v5_services_video_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v5_services_video_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoRequest); i {
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
			RawDescriptor: file_google_ads_googleads_v5_services_video_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v5_services_video_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v5_services_video_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v5_services_video_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v5_services_video_service_proto = out.File
	file_google_ads_googleads_v5_services_video_service_proto_rawDesc = nil
	file_google_ads_googleads_v5_services_video_service_proto_goTypes = nil
	file_google_ads_googleads_v5_services_video_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoServiceClient interface {
	// Returns the requested video in full detail.
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*resources.Video, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*resources.Video, error) {
	out := new(resources.Video)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v5.services.VideoService/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
type VideoServiceServer interface {
	// Returns the requested video in full detail.
	GetVideo(context.Context, *GetVideoRequest) (*resources.Video, error)
}

// UnimplementedVideoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (*UnimplementedVideoServiceServer) GetVideo(context.Context, *GetVideoRequest) (*resources.Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}

func RegisterVideoServiceServer(s *grpc.Server, srv VideoServiceServer) {
	s.RegisterService(&_VideoService_serviceDesc, srv)
}

func _VideoService_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v5.services.VideoService/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v5.services.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideo",
			Handler:    _VideoService_GetVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v5/services/video_service.proto",
}
