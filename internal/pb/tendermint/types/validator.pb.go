// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tendermint/types/validator.proto

package types

import (
	_ "github.com/cosmos/gogoproto/gogoproto"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
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

type ValidatorSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validators       []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	Proposer         *Validator   `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	TotalVotingPower int64        `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
}

func (x *ValidatorSet) Reset() {
	*x = ValidatorSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorSet) ProtoMessage() {}

func (x *ValidatorSet) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorSet.ProtoReflect.Descriptor instead.
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return file_tendermint_types_validator_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorSet) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *ValidatorSet) GetProposer() *Validator {
	if x != nil {
		return x.Proposer
	}
	return nil
}

func (x *ValidatorSet) GetTotalVotingPower() int64 {
	if x != nil {
		return x.TotalVotingPower
	}
	return 0
}

type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address          []byte            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey           *crypto.PublicKey `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VotingPower      int64             `protobuf:"varint,3,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	ProposerPriority int64             `protobuf:"varint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_validator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_validator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_tendermint_types_validator_proto_rawDescGZIP(), []int{1}
}

func (x *Validator) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Validator) GetPubKey() *crypto.PublicKey {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Validator) GetVotingPower() int64 {
	if x != nil {
		return x.VotingPower
	}
	return 0
}

func (x *Validator) GetProposerPriority() int64 {
	if x != nil {
		return x.ProposerPriority
	}
	return 0
}

type SimpleValidator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey      *crypto.PublicKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VotingPower int64             `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (x *SimpleValidator) Reset() {
	*x = SimpleValidator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_validator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleValidator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleValidator) ProtoMessage() {}

func (x *SimpleValidator) ProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_validator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleValidator.ProtoReflect.Descriptor instead.
func (*SimpleValidator) Descriptor() ([]byte, []int) {
	return file_tendermint_types_validator_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleValidator) GetPubKey() *crypto.PublicKey {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *SimpleValidator) GetVotingPower() int64 {
	if x != nil {
		return x.VotingPower
	}
	return 0
}

var File_tendermint_types_validator_proto protoreflect.FileDescriptor

var file_tendermint_types_validator_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6b, 0x65,
	0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x22, 0xb2, 0x01,
	0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x6b, 0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tendermint_types_validator_proto_rawDescOnce sync.Once
	file_tendermint_types_validator_proto_rawDescData = file_tendermint_types_validator_proto_rawDesc
)

func file_tendermint_types_validator_proto_rawDescGZIP() []byte {
	file_tendermint_types_validator_proto_rawDescOnce.Do(func() {
		file_tendermint_types_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_types_validator_proto_rawDescData)
	})
	return file_tendermint_types_validator_proto_rawDescData
}

var file_tendermint_types_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tendermint_types_validator_proto_goTypes = []interface{}{
	(*ValidatorSet)(nil),     // 0: tendermint.types.ValidatorSet
	(*Validator)(nil),        // 1: tendermint.types.Validator
	(*SimpleValidator)(nil),  // 2: tendermint.types.SimpleValidator
	(*crypto.PublicKey)(nil), // 3: tendermint.crypto.PublicKey
}
var file_tendermint_types_validator_proto_depIdxs = []int32{
	1, // 0: tendermint.types.ValidatorSet.validators:type_name -> tendermint.types.Validator
	1, // 1: tendermint.types.ValidatorSet.proposer:type_name -> tendermint.types.Validator
	3, // 2: tendermint.types.Validator.pub_key:type_name -> tendermint.crypto.PublicKey
	3, // 3: tendermint.types.SimpleValidator.pub_key:type_name -> tendermint.crypto.PublicKey
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tendermint_types_validator_proto_init() }
func file_tendermint_types_validator_proto_init() {
	if File_tendermint_types_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_types_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorSet); i {
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
		file_tendermint_types_validator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
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
		file_tendermint_types_validator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleValidator); i {
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
			RawDescriptor: file_tendermint_types_validator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_types_validator_proto_goTypes,
		DependencyIndexes: file_tendermint_types_validator_proto_depIdxs,
		MessageInfos:      file_tendermint_types_validator_proto_msgTypes,
	}.Build()
	File_tendermint_types_validator_proto = out.File
	file_tendermint_types_validator_proto_rawDesc = nil
	file_tendermint_types_validator_proto_goTypes = nil
	file_tendermint_types_validator_proto_depIdxs = nil
}
