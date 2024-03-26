// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: confluent/meta.proto

package confluent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doc    string            `protobuf:"bytes,1,opt,name=doc,proto3" json:"doc,omitempty"`
	Params map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags   []string          `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_confluent_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_confluent_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_confluent_meta_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetDoc() string {
	if x != nil {
		return x.Doc
	}
	return ""
}

func (x *Meta) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Meta) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var file_confluent_meta_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*Meta)(nil),
		Field:         1088,
		Name:          "confluent.file_meta",
		Tag:           "bytes,1088,opt,name=file_meta",
		Filename:      "confluent/meta.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Meta)(nil),
		Field:         1088,
		Name:          "confluent.message_meta",
		Tag:           "bytes,1088,opt,name=message_meta",
		Filename:      "confluent/meta.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Meta)(nil),
		Field:         1088,
		Name:          "confluent.field_meta",
		Tag:           "bytes,1088,opt,name=field_meta",
		Filename:      "confluent/meta.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*Meta)(nil),
		Field:         1088,
		Name:          "confluent.enum_meta",
		Tag:           "bytes,1088,opt,name=enum_meta",
		Filename:      "confluent/meta.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*Meta)(nil),
		Field:         1088,
		Name:          "confluent.enum_value_meta",
		Tag:           "bytes,1088,opt,name=enum_value_meta",
		Filename:      "confluent/meta.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional confluent.Meta file_meta = 1088;
	E_FileMeta = &file_confluent_meta_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional confluent.Meta message_meta = 1088;
	E_MessageMeta = &file_confluent_meta_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional confluent.Meta field_meta = 1088;
	E_FieldMeta = &file_confluent_meta_proto_extTypes[2]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional confluent.Meta enum_meta = 1088;
	E_EnumMeta = &file_confluent_meta_proto_extTypes[3]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional confluent.Meta enum_value_meta = 1088;
	E_EnumValueMeta = &file_confluent_meta_proto_extTypes[4]
)

var File_confluent_meta_proto protoreflect.FileDescriptor

var file_confluent_meta_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e,
	0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x63, 0x12, 0x33,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x3a, 0x4b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc0, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x3a,
	0x54, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xc0, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x3a, 0x4e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xc0, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x6c, 0x75, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x65, 0x74, 0x61, 0x3a, 0x4b, 0x0a, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xc0, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x65, 0x6e, 0x75, 0x6d, 0x4d, 0x65,
	0x74, 0x61, 0x3a, 0x5b, 0x0a, 0x0f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc0, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x0d, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x42,
	0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_confluent_meta_proto_rawDescOnce sync.Once
	file_confluent_meta_proto_rawDescData = file_confluent_meta_proto_rawDesc
)

func file_confluent_meta_proto_rawDescGZIP() []byte {
	file_confluent_meta_proto_rawDescOnce.Do(func() {
		file_confluent_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_confluent_meta_proto_rawDescData)
	})
	return file_confluent_meta_proto_rawDescData
}

var file_confluent_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_confluent_meta_proto_goTypes = []interface{}{
	(*Meta)(nil),                          // 0: confluent.Meta
	nil,                                   // 1: confluent.Meta.ParamsEntry
	(*descriptorpb.FileOptions)(nil),      // 2: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil),   // 3: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),     // 4: google.protobuf.FieldOptions
	(*descriptorpb.EnumOptions)(nil),      // 5: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 6: google.protobuf.EnumValueOptions
}
var file_confluent_meta_proto_depIdxs = []int32{
	1,  // 0: confluent.Meta.params:type_name -> confluent.Meta.ParamsEntry
	2,  // 1: confluent.file_meta:extendee -> google.protobuf.FileOptions
	3,  // 2: confluent.message_meta:extendee -> google.protobuf.MessageOptions
	4,  // 3: confluent.field_meta:extendee -> google.protobuf.FieldOptions
	5,  // 4: confluent.enum_meta:extendee -> google.protobuf.EnumOptions
	6,  // 5: confluent.enum_value_meta:extendee -> google.protobuf.EnumValueOptions
	0,  // 6: confluent.file_meta:type_name -> confluent.Meta
	0,  // 7: confluent.message_meta:type_name -> confluent.Meta
	0,  // 8: confluent.field_meta:type_name -> confluent.Meta
	0,  // 9: confluent.enum_meta:type_name -> confluent.Meta
	0,  // 10: confluent.enum_value_meta:type_name -> confluent.Meta
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	6,  // [6:11] is the sub-list for extension type_name
	1,  // [1:6] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_confluent_meta_proto_init() }
func file_confluent_meta_proto_init() {
	if File_confluent_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_confluent_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_confluent_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_confluent_meta_proto_goTypes,
		DependencyIndexes: file_confluent_meta_proto_depIdxs,
		MessageInfos:      file_confluent_meta_proto_msgTypes,
		ExtensionInfos:    file_confluent_meta_proto_extTypes,
	}.Build()
	File_confluent_meta_proto = out.File
	file_confluent_meta_proto_rawDesc = nil
	file_confluent_meta_proto_goTypes = nil
	file_confluent_meta_proto_depIdxs = nil
}
