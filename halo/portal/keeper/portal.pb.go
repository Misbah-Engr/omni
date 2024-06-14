// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: halo/portal/keeper/portal.proto

package keeper

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
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

// Block groups a set of Msgs adding height and offset.
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // Auto-incremented ID (BlockHeight, BlockOffset)
	CreatedHeight uint64 `protobuf:"varint,2,opt,name=created_height,json=createdHeight,proto3" json:"created_height,omitempty"` // Height this block was created at.
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halo_portal_keeper_portal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_halo_portal_keeper_portal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_halo_portal_keeper_portal_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Block) GetCreatedHeight() uint64 {
	if x != nil {
		return x.CreatedHeight
	}
	return 0
}

// Msg represents a single cross-chain message emitted by the consensus chain portal.
type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                         // Auto-incremented ID
	BlockId      uint64 `protobuf:"varint,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`                // Block ID to which this msg belongs
	MsgType      uint32 `protobuf:"varint,3,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`                // Message type (ValidatorSet, Withdrawal, etc.)
	MsgTypeId    uint64 `protobuf:"varint,4,opt,name=msg_type_id,json=msgTypeId,proto3" json:"msg_type_id,omitempty"`        // ID of the type referred to be MsgType
	DestChainId  uint64 `protobuf:"varint,5,opt,name=dest_chain_id,json=destChainId,proto3" json:"dest_chain_id,omitempty"`  // Destination chain ID
	ShardId      uint64 `protobuf:"varint,6,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`                // Shard of the message
	StreamOffset uint64 `protobuf:"varint,7,opt,name=stream_offset,json=streamOffset,proto3" json:"stream_offset,omitempty"` // Offset of the message in the stream
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halo_portal_keeper_portal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_halo_portal_keeper_portal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_halo_portal_keeper_portal_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Msg) GetBlockId() uint64 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

func (x *Msg) GetMsgType() uint32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *Msg) GetMsgTypeId() uint64 {
	if x != nil {
		return x.MsgTypeId
	}
	return 0
}

func (x *Msg) GetDestChainId() uint64 {
	if x != nil {
		return x.DestChainId
	}
	return 0
}

func (x *Msg) GetShardId() uint64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *Msg) GetStreamOffset() uint64 {
	if x != nil {
		return x.StreamOffset
	}
	return 0
}

type Offset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                        // Auto-incremented ID
	DestChainId uint64 `protobuf:"varint,2,opt,name=dest_chain_id,json=destChainId,proto3" json:"dest_chain_id,omitempty"` // Destination chain ID
	ShardId     uint64 `protobuf:"varint,3,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`               // Shard ID
	Offset      uint64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`                                // Offset of the last message in the stream
}

func (x *Offset) Reset() {
	*x = Offset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halo_portal_keeper_portal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offset) ProtoMessage() {}

func (x *Offset) ProtoReflect() protoreflect.Message {
	mi := &file_halo_portal_keeper_portal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offset.ProtoReflect.Descriptor instead.
func (*Offset) Descriptor() ([]byte, []int) {
	return file_halo_portal_keeper_portal_proto_rawDescGZIP(), []int{2}
}

func (x *Offset) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Offset) GetDestChainId() uint64 {
	if x != nil {
		return x.DestChainId
	}
	return 0
}

func (x *Offset) GetShardId() uint64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *Offset) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_halo_portal_keeper_portal_proto protoreflect.FileDescriptor

var file_halo_portal_keeper_portal_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x68, 0x61, 0x6c, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66,
	0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x26,
	0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x20, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x10, 0x02, 0x18, 0x01, 0x18, 0x01, 0x22, 0xef, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x3a, 0x1e, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x18,
	0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x69, 0x64, 0x10, 0x02, 0x18, 0x02, 0x22, 0x9f, 0x01, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x3a, 0x2e, 0xf2, 0x9e, 0xd3, 0x8e,
	0x03, 0x28, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x16, 0x64, 0x65,
	0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x2c, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x10, 0x02, 0x18, 0x01, 0x18, 0x03, 0x42, 0xc0, 0x01, 0x0a, 0x16, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x42, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x6d,
	0x6e, 0x69, 0x2f, 0x68, 0x61, 0x6c, 0x6f, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x48, 0x50, 0x4b, 0xaa, 0x02, 0x12, 0x48, 0x61,
	0x6c, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0xca, 0x02, 0x12, 0x48, 0x61, 0x6c, 0x6f, 0x5c, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x4b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0xe2, 0x02, 0x1e, 0x48, 0x61, 0x6c, 0x6f, 0x5c, 0x50, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x5c, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x48, 0x61, 0x6c, 0x6f, 0x3a, 0x3a, 0x50,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x3a, 0x3a, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_halo_portal_keeper_portal_proto_rawDescOnce sync.Once
	file_halo_portal_keeper_portal_proto_rawDescData = file_halo_portal_keeper_portal_proto_rawDesc
)

func file_halo_portal_keeper_portal_proto_rawDescGZIP() []byte {
	file_halo_portal_keeper_portal_proto_rawDescOnce.Do(func() {
		file_halo_portal_keeper_portal_proto_rawDescData = protoimpl.X.CompressGZIP(file_halo_portal_keeper_portal_proto_rawDescData)
	})
	return file_halo_portal_keeper_portal_proto_rawDescData
}

var file_halo_portal_keeper_portal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_halo_portal_keeper_portal_proto_goTypes = []any{
	(*Block)(nil),  // 0: halo.portal.keeper.Block
	(*Msg)(nil),    // 1: halo.portal.keeper.Msg
	(*Offset)(nil), // 2: halo.portal.keeper.Offset
}
var file_halo_portal_keeper_portal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_halo_portal_keeper_portal_proto_init() }
func file_halo_portal_keeper_portal_proto_init() {
	if File_halo_portal_keeper_portal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_halo_portal_keeper_portal_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Block); i {
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
		file_halo_portal_keeper_portal_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Msg); i {
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
		file_halo_portal_keeper_portal_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Offset); i {
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
			RawDescriptor: file_halo_portal_keeper_portal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_halo_portal_keeper_portal_proto_goTypes,
		DependencyIndexes: file_halo_portal_keeper_portal_proto_depIdxs,
		MessageInfos:      file_halo_portal_keeper_portal_proto_msgTypes,
	}.Build()
	File_halo_portal_keeper_portal_proto = out.File
	file_halo_portal_keeper_portal_proto_rawDesc = nil
	file_halo_portal_keeper_portal_proto_goTypes = nil
	file_halo_portal_keeper_portal_proto_depIdxs = nil
}
