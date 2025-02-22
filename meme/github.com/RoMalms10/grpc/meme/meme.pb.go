// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: meme/meme.proto

package meme

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenerateMemeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TemplateName  string                 `protobuf:"bytes,1,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	TopText       string                 `protobuf:"bytes,2,opt,name=top_text,json=topText,proto3" json:"top_text,omitempty"`
	BottomText    string                 `protobuf:"bytes,3,opt,name=bottom_text,json=bottomText,proto3" json:"bottom_text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateMemeRequest) Reset() {
	*x = GenerateMemeRequest{}
	mi := &file_meme_meme_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateMemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateMemeRequest) ProtoMessage() {}

func (x *GenerateMemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_meme_meme_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateMemeRequest.ProtoReflect.Descriptor instead.
func (*GenerateMemeRequest) Descriptor() ([]byte, []int) {
	return file_meme_meme_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateMemeRequest) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *GenerateMemeRequest) GetTopText() string {
	if x != nil {
		return x.TopText
	}
	return ""
}

func (x *GenerateMemeRequest) GetBottomText() string {
	if x != nil {
		return x.BottomText
	}
	return ""
}

type GenerateMemeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MemeUrl       string                 `protobuf:"bytes,1,opt,name=meme_url,json=memeUrl,proto3" json:"meme_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateMemeResponse) Reset() {
	*x = GenerateMemeResponse{}
	mi := &file_meme_meme_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateMemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateMemeResponse) ProtoMessage() {}

func (x *GenerateMemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_meme_meme_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateMemeResponse.ProtoReflect.Descriptor instead.
func (*GenerateMemeResponse) Descriptor() ([]byte, []int) {
	return file_meme_meme_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateMemeResponse) GetMemeUrl() string {
	if x != nil {
		return x.MemeUrl
	}
	return ""
}

var File_meme_meme_proto protoreflect.FileDescriptor

var file_meme_meme_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x65, 0x2f, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x65, 0x6d, 0x65, 0x22, 0x76, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x22,
	0x31, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x65, 0x55,
	0x72, 0x6c, 0x32, 0x54, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x65, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d,
	0x65, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x4d, 0x61, 0x6c, 0x6d, 0x73, 0x31, 0x30,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_meme_meme_proto_rawDescOnce sync.Once
	file_meme_meme_proto_rawDescData []byte
)

func file_meme_meme_proto_rawDescGZIP() []byte {
	file_meme_meme_proto_rawDescOnce.Do(func() {
		file_meme_meme_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_meme_meme_proto_rawDesc), len(file_meme_meme_proto_rawDesc)))
	})
	return file_meme_meme_proto_rawDescData
}

var file_meme_meme_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_meme_meme_proto_goTypes = []any{
	(*GenerateMemeRequest)(nil),  // 0: meme.GenerateMemeRequest
	(*GenerateMemeResponse)(nil), // 1: meme.GenerateMemeResponse
}
var file_meme_meme_proto_depIdxs = []int32{
	0, // 0: meme.MemeService.GenerateMeme:input_type -> meme.GenerateMemeRequest
	1, // 1: meme.MemeService.GenerateMeme:output_type -> meme.GenerateMemeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meme_meme_proto_init() }
func file_meme_meme_proto_init() {
	if File_meme_meme_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_meme_meme_proto_rawDesc), len(file_meme_meme_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meme_meme_proto_goTypes,
		DependencyIndexes: file_meme_meme_proto_depIdxs,
		MessageInfos:      file_meme_meme_proto_msgTypes,
	}.Build()
	File_meme_meme_proto = out.File
	file_meme_meme_proto_goTypes = nil
	file_meme_meme_proto_depIdxs = nil
}
