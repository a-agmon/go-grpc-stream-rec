// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: rcservice.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_rcservice_proto protoreflect.FileDescriptor

var file_rcservice_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x72, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x0f, 0x64, 0x69, 0x73,
	0x74, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x55, 0x0a,
	0x09, 0x52, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x72, 0x63, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x2d, 0x61, 0x67, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x63, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_rcservice_proto_goTypes = []interface{}{
	(*RecommendRequest)(nil),  // 0: rcserver.RecommendRequest
	(*RecommendResponse)(nil), // 1: rcserver.RecommendResponse
}
var file_rcservice_proto_depIdxs = []int32{
	0, // 0: rcserver.RcService.Recommend:input_type -> rcserver.RecommendRequest
	1, // 1: rcserver.RcService.Recommend:output_type -> rcserver.RecommendResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rcservice_proto_init() }
func file_rcservice_proto_init() {
	if File_rcservice_proto != nil {
		return
	}
	file_distscore_proto_init()
	file_recommend_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rcservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rcservice_proto_goTypes,
		DependencyIndexes: file_rcservice_proto_depIdxs,
	}.Build()
	File_rcservice_proto = out.File
	file_rcservice_proto_rawDesc = nil
	file_rcservice_proto_goTypes = nil
	file_rcservice_proto_depIdxs = nil
}