// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: recommend.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RecommendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *RecommendRequest) Reset() {
	*x = RecommendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendRequest) ProtoMessage() {}

func (x *RecommendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recommend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendRequest.ProtoReflect.Descriptor instead.
func (*RecommendRequest) Descriptor() ([]byte, []int) {
	return file_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *RecommendRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type RecommendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *RecommendResponse) Reset() {
	*x = RecommendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendResponse) ProtoMessage() {}

func (x *RecommendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recommend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendResponse.ProtoReflect.Descriptor instead.
func (*RecommendResponse) Descriptor() ([]byte, []int) {
	return file_recommend_proto_rawDescGZIP(), []int{1}
}

func (x *RecommendResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type RecommendManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []string `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *RecommendManyRequest) Reset() {
	*x = RecommendManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendManyRequest) ProtoMessage() {}

func (x *RecommendManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recommend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendManyRequest.ProtoReflect.Descriptor instead.
func (*RecommendManyRequest) Descriptor() ([]byte, []int) {
	return file_recommend_proto_rawDescGZIP(), []int{2}
}

func (x *RecommendManyRequest) GetContent() []string {
	if x != nil {
		return x.Content
	}
	return nil
}

type RecommendManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []string `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *RecommendManyResponse) Reset() {
	*x = RecommendManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendManyResponse) ProtoMessage() {}

func (x *RecommendManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recommend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendManyResponse.ProtoReflect.Descriptor instead.
func (*RecommendManyResponse) Descriptor() ([]byte, []int) {
	return file_recommend_proto_rawDescGZIP(), []int{3}
}

func (x *RecommendManyResponse) GetContent() []string {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_recommend_proto protoreflect.FileDescriptor

var file_recommend_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x72, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x10, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x15, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x2d, 0x61, 0x67,
	0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recommend_proto_rawDescOnce sync.Once
	file_recommend_proto_rawDescData = file_recommend_proto_rawDesc
)

func file_recommend_proto_rawDescGZIP() []byte {
	file_recommend_proto_rawDescOnce.Do(func() {
		file_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(file_recommend_proto_rawDescData)
	})
	return file_recommend_proto_rawDescData
}

var file_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_recommend_proto_goTypes = []interface{}{
	(*RecommendRequest)(nil),      // 0: rcserver.RecommendRequest
	(*RecommendResponse)(nil),     // 1: rcserver.RecommendResponse
	(*RecommendManyRequest)(nil),  // 2: rcserver.RecommendManyRequest
	(*RecommendManyResponse)(nil), // 3: rcserver.RecommendManyResponse
}
var file_recommend_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_recommend_proto_init() }
func file_recommend_proto_init() {
	if File_recommend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recommend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendRequest); i {
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
		file_recommend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendResponse); i {
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
		file_recommend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendManyRequest); i {
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
		file_recommend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendManyResponse); i {
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
			RawDescriptor: file_recommend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_recommend_proto_goTypes,
		DependencyIndexes: file_recommend_proto_depIdxs,
		MessageInfos:      file_recommend_proto_msgTypes,
	}.Build()
	File_recommend_proto = out.File
	file_recommend_proto_rawDesc = nil
	file_recommend_proto_goTypes = nil
	file_recommend_proto_depIdxs = nil
}
