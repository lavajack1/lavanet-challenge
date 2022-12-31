// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: cosmos/query/v1/query.proto

package query

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

var file_cosmos_query_v1_query_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         11110001,
		Name:          "cosmos.query.v1.module_query_safe",
		Tag:           "varint,11110001,opt,name=module_query_safe",
		Filename:      "cosmos/query/v1/query.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// module_query_safe is set to true when the query is safe to be called from
	// within the state machine, for example from another module's Keeper, via
	// ADR-033 calls or from CosmWasm contracts.
	// Concretely, it means that the query is:
	// 1. deterministic: given a block height, returns the exact same response
	// upon multiple calls; and doesn't introduce any state-machine-breaking
	// changes across SDK patch version.
	// 2. consumes gas correctly.
	//
	// If you are a module developer and want to add this annotation to one of
	// your own queries, please make sure that the corresponding query:
	// 1. is deterministic and won't introduce state-machine-breaking changes
	// without a coordinated upgrade path,
	// 2. has its gas tracked, to avoid the attack vector where no gas is
	// accounted for on potentially high-computation queries.
	//
	// For queries that potentially consume a large amount of gas (for example
	// those with pagination, if the pagination field is incorrectly set), we
	// also recommend adding Protobuf comments to warn module developers
	// consuming these queries.
	//
	// When set to true, the query can safely be called
	//
	// optional bool module_query_safe = 11110001;
	E_ModuleQuerySafe = &file_cosmos_query_v1_query_proto_extTypes[0]
)

var File_cosmos_query_v1_query_proto protoreflect.FileDescriptor

var file_cosmos_query_v1_query_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3a, 0x4d, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x73, 0x61, 0x66, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf1, 0x8c, 0xa6, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x61, 0x66, 0x65, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_cosmos_query_v1_query_proto_goTypes = []interface{}{
	(*descriptorpb.MethodOptions)(nil), // 0: google.protobuf.MethodOptions
}
var file_cosmos_query_v1_query_proto_depIdxs = []int32{
	0, // 0: cosmos.query.v1.module_query_safe:extendee -> google.protobuf.MethodOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_query_v1_query_proto_init() }
func file_cosmos_query_v1_query_proto_init() {
	if File_cosmos_query_v1_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_query_v1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_query_v1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_query_v1_query_proto_depIdxs,
		ExtensionInfos:    file_cosmos_query_v1_query_proto_extTypes,
	}.Build()
	File_cosmos_query_v1_query_proto = out.File
	file_cosmos_query_v1_query_proto_rawDesc = nil
	file_cosmos_query_v1_query_proto_goTypes = nil
	file_cosmos_query_v1_query_proto_depIdxs = nil
}
