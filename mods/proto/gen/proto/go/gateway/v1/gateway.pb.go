// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gateway/v1/gateway.proto

package gatewayv1

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

type PacketType int32

const (
	PacketType_PACKET_TYPE_UNKNOWN PacketType = 0
	// new stream transmission
	PacketType_PACKET_TYPE_START PacketType = 1
	// stream data
	PacketType_PACKET_TYPE_DATA PacketType = 2
	// stream transmission is complete
	PacketType_PACKET_TYPE_END PacketType = 3
	// stream transmission failure
	PacketType_PACKET_TYPE_ERROR PacketType = 4
	// mix-net message filler
	PacketType_PACKET_TYPE_DUMMY PacketType = 5
)

// Enum value maps for PacketType.
var (
	PacketType_name = map[int32]string{
		0: "PACKET_TYPE_UNKNOWN",
		1: "PACKET_TYPE_START",
		2: "PACKET_TYPE_DATA",
		3: "PACKET_TYPE_END",
		4: "PACKET_TYPE_ERROR",
		5: "PACKET_TYPE_DUMMY",
	}
	PacketType_value = map[string]int32{
		"PACKET_TYPE_UNKNOWN": 0,
		"PACKET_TYPE_START":   1,
		"PACKET_TYPE_DATA":    2,
		"PACKET_TYPE_END":     3,
		"PACKET_TYPE_ERROR":   4,
		"PACKET_TYPE_DUMMY":   5,
	}
)

func (x PacketType) Enum() *PacketType {
	p := new(PacketType)
	*p = x
	return p
}

func (x PacketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketType) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_v1_gateway_proto_enumTypes[0].Descriptor()
}

func (PacketType) Type() protoreflect.EnumType {
	return &file_gateway_v1_gateway_proto_enumTypes[0]
}

func (x PacketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketType.Descriptor instead.
func (PacketType) EnumDescriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{0}
}

// Gateway data streams are transmitted within mix-net messages as packets.
// Packets track sequence to facilitate reassembly upon exit of the mix-net.
type PacketHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// https://protobuf.dev/programming-guides/encoding/#bools-and-enums
	// Enums are encoded as if they were int32, which is variable size.
	Type PacketType `protobuf:"varint,1,opt,name=type,proto3,enum=gateway.v1.PacketType" json:"type,omitempty"`
	// unique identifier for the data stream
	StreamId uint64 `protobuf:"varint,2,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	// the order of the packet within the stream
	Sequence uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *PacketHeader) Reset() {
	*x = PacketHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketHeader) ProtoMessage() {}

func (x *PacketHeader) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketHeader.ProtoReflect.Descriptor instead.
func (*PacketHeader) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *PacketHeader) GetType() PacketType {
	if x != nil {
		return x.Type
	}
	return PacketType_PACKET_TYPE_UNKNOWN
}

func (x *PacketHeader) GetStreamId() uint64 {
	if x != nil {
		return x.StreamId
	}
	return 0
}

func (x *PacketHeader) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

var File_gateway_v1_gateway_proto protoreflect.FileDescriptor

var file_gateway_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x73, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x2a, 0x95, 0x01, 0x0a, 0x0a,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41,
	0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x41,
	0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x55, 0x4d, 0x4d,
	0x59, 0x10, 0x05, 0x42, 0xaf, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x33, 0x31, 0x33, 0x33, 0x33, 0x33, 0x33, 0x37, 0x2f, 0x62, 0x6d, 0x72, 0x6e,
	0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0a, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x16, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_v1_gateway_proto_rawDescOnce sync.Once
	file_gateway_v1_gateway_proto_rawDescData = file_gateway_v1_gateway_proto_rawDesc
)

func file_gateway_v1_gateway_proto_rawDescGZIP() []byte {
	file_gateway_v1_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_v1_gateway_proto_rawDescData)
	})
	return file_gateway_v1_gateway_proto_rawDescData
}

var file_gateway_v1_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gateway_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gateway_v1_gateway_proto_goTypes = []interface{}{
	(PacketType)(0),      // 0: gateway.v1.PacketType
	(*PacketHeader)(nil), // 1: gateway.v1.PacketHeader
}
var file_gateway_v1_gateway_proto_depIdxs = []int32{
	0, // 0: gateway.v1.PacketHeader.type:type_name -> gateway.v1.PacketType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gateway_v1_gateway_proto_init() }
func file_gateway_v1_gateway_proto_init() {
	if File_gateway_v1_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_v1_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketHeader); i {
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
			RawDescriptor: file_gateway_v1_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_v1_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_v1_gateway_proto_depIdxs,
		EnumInfos:         file_gateway_v1_gateway_proto_enumTypes,
		MessageInfos:      file_gateway_v1_gateway_proto_msgTypes,
	}.Build()
	File_gateway_v1_gateway_proto = out.File
	file_gateway_v1_gateway_proto_rawDesc = nil
	file_gateway_v1_gateway_proto_goTypes = nil
	file_gateway_v1_gateway_proto_depIdxs = nil
}
