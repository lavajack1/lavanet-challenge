// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: amino/amino.proto

package amino

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_amino_amino_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         11110001,
		Name:          "amino.name",
		Tag:           "bytes,11110001,opt,name=name",
		Filename:      "amino/amino.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         11110002,
		Name:          "amino.message_encoding",
		Tag:           "bytes,11110002,opt,name=message_encoding",
		Filename:      "amino/amino.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         11110003,
		Name:          "amino.encoding",
		Tag:           "bytes,11110003,opt,name=encoding",
		Filename:      "amino/amino.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         11110004,
		Name:          "amino.field_name",
		Tag:           "bytes,11110004,opt,name=field_name",
		Filename:      "amino/amino.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         11110005,
		Name:          "amino.dont_omitempty",
		Tag:           "varint,11110005,opt,name=dont_omitempty",
		Filename:      "amino/amino.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// name is the string used when registering a concrete
	// type into the Amino type registry, via the Amino codec's
	// `RegisterConcrete()` method. This string MUST be at most 39
	// characters long, or else the message will be rejected by the
	// Ledger hardware device.
	//
	// optional string name = 11110001;
	E_Name = &file_amino_amino_proto_extTypes[0]
	// encoding describes the encoding format used by Amino for the given
	// message. The field type is chosen to be a string for
	// flexibility, but it should ideally be short and expected to be
	// machine-readable, for example "base64" or "utf8_json". We
	// highly recommend to use underscores for word separation instead of spaces.
	//
	// If left empty, then the Amino encoding is expected to be the same as the
	// Protobuf one.
	//
	// This annotation should not be confused with the `encoding`
	// one which operates on the field level.
	//
	// optional string message_encoding = 11110002;
	E_MessageEncoding = &file_amino_amino_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// encoding describes the encoding format used by Amino for
	// the given field. The field type is chosen to be a string for
	// flexibility, but it should ideally be short and expected to be
	// machine-readable, for example "base64" or "utf8_json". We
	// highly recommend to use underscores for word separation instead of spaces.
	//
	// If left empty, then the Amino encoding is expected to be the same as the
	// Protobuf one.
	//
	// This annotation should not be confused with the
	// `message_encoding` one which operates on the message level.
	//
	// optional string encoding = 11110003;
	E_Encoding = &file_amino_amino_proto_extTypes[2]
	// field_name sets a different field name (i.e. key name) in
	// the amino JSON object for the given field.
	//
	// Example:
	//
	// message Foo {
	//   string bar = 1 [(amino.field_name) = "baz"];
	// }
	//
	// Then the Amino encoding of Foo will be:
	// `{"baz":"some value"}`
	//
	// optional string field_name = 11110004;
	E_FieldName = &file_amino_amino_proto_extTypes[3]
	// dont_omitempty sets the field in the JSON object even if
	// its value is empty, i.e. equal to the Golang zero value. To learn what
	// the zero values are, see https://go.dev/ref/spec#The_zero_value.
	//
	// Fields default to `omitempty`, which is the default behavior when this
	// annotation is unset. When set to true, then the field value in the
	// JSON object will be set, i.e. not `undefined`.
	//
	// Example:
	//
	// message Foo {
	//   string bar = 1;
	//   string baz = 2 [(amino.dont_omitempty) = true];
	// }
	//
	// f := Foo{};
	// out := AminoJSONEncoder(&f);
	// out == {"baz":""}
	//
	// optional bool dont_omitempty = 11110005;
	E_DontOmitempty = &file_amino_amino_proto_extTypes[4]
)

var File_amino_amino_proto protoreflect.FileDescriptor

var file_amino_amino_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x36, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf1, 0x8c, 0xa6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x4d, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf2, 0x8c, 0xa6, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x3a, 0x3c, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf3,
	0x8c, 0xa6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x3a, 0x3f, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf4,
	0x8c, 0xa6, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x3a, 0x47, 0x0a, 0x0e, 0x64, 0x6f, 0x6e, 0x74, 0x5f, 0x6f, 0x6d, 0x69, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf5, 0x8c, 0xa6, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x6f,
	0x6e, 0x74, 0x4f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x74, 0x78, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_amino_amino_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
}
var file_amino_amino_proto_depIdxs = []int32{
	0, // 0: amino.name:extendee -> google.protobuf.MessageOptions
	0, // 1: amino.message_encoding:extendee -> google.protobuf.MessageOptions
	1, // 2: amino.encoding:extendee -> google.protobuf.FieldOptions
	1, // 3: amino.field_name:extendee -> google.protobuf.FieldOptions
	1, // 4: amino.dont_omitempty:extendee -> google.protobuf.FieldOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_amino_amino_proto_init() }
func file_amino_amino_proto_init() {
	if File_amino_amino_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_amino_amino_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_amino_amino_proto_goTypes,
		DependencyIndexes: file_amino_amino_proto_depIdxs,
		ExtensionInfos:    file_amino_amino_proto_extTypes,
	}.Build()
	File_amino_amino_proto = out.File
	file_amino_amino_proto_rawDesc = nil
	file_amino_amino_proto_goTypes = nil
	file_amino_amino_proto_depIdxs = nil
}
