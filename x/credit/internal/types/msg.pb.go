// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/credit/internal/types/msg.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// The message to add credits.
type MsgAdd struct {
	// The msg's signer.
	Signer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty" validate:"required,accaddress"`
	// The address to add the credits.
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" validate:"required,accaddress"`
	// amount of credits to add.
	Amount               github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount,omitempty" hash:"name:1" validate:"required,bigint"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *MsgAdd) Reset()         { *m = MsgAdd{} }
func (m *MsgAdd) String() string { return proto.CompactTextString(m) }
func (*MsgAdd) ProtoMessage()    {}
func (*MsgAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_def7da7b61d064e7, []int{0}
}
func (m *MsgAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgAdd.Unmarshal(m, b)
}
func (m *MsgAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgAdd.Marshal(b, m, deterministic)
}
func (m *MsgAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAdd.Merge(m, src)
}
func (m *MsgAdd) XXX_Size() int {
	return xxx_messageInfo_MsgAdd.Size(m)
}
func (m *MsgAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAdd.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAdd proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAdd)(nil), "mesg.credit.types.MsgAdd")
}

func init() { proto.RegisterFile("x/credit/internal/types/msg.proto", fileDescriptor_def7da7b61d064e7) }

var fileDescriptor_def7da7b61d064e7 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x69, 0x91, 0x82, 0x88, 0x58, 0xc8, 0x54, 0x31, 0xb4, 0xc1, 0x03, 0xca, 0x40, 0x62,
	0x21, 0x26, 0xb2, 0xa0, 0x64, 0x63, 0x60, 0x29, 0x12, 0x03, 0x9b, 0x63, 0x5f, 0x1d, 0xab, 0xf5,
	0xb9, 0xd8, 0x0e, 0x82, 0x6f, 0xc8, 0x07, 0x60, 0x8e, 0xc4, 0x57, 0xe8, 0xc8, 0x84, 0xf2, 0x07,
	0x89, 0x01, 0x24, 0x16, 0x26, 0xfb, 0xe9, 0xdd, 0xbb, 0xdf, 0x49, 0x2f, 0x3c, 0x7d, 0xa6, 0xdc,
	0x82, 0x50, 0x9e, 0x2a, 0xf4, 0x60, 0x91, 0x6d, 0xa8, 0x7f, 0xd9, 0x82, 0xa3, 0xda, 0xc9, 0x6c,
	0x6b, 0x8d, 0x37, 0xd1, 0xb1, 0x06, 0x27, 0xb3, 0x61, 0x2a, 0xeb, 0xcd, 0x13, 0x22, 0x8d, 0x34,
	0xb4, 0xb7, 0xab, 0x66, 0x45, 0x3b, 0xd5, 0x8b, 0xfe, 0x37, 0xc4, 0xc8, 0xdb, 0x34, 0x0c, 0x6e,
	0x9d, 0x2c, 0x84, 0x88, 0xd6, 0x61, 0xe0, 0x94, 0x44, 0xb0, 0xb3, 0x49, 0x3c, 0x49, 0x8e, 0xca,
	0xbb, 0x5d, 0xbb, 0x98, 0x3f, 0xb1, 0x8d, 0x12, 0xcc, 0x43, 0x4e, 0x2c, 0x3c, 0x36, 0xca, 0x82,
	0x38, 0x67, 0x9c, 0x33, 0x21, 0x2c, 0x38, 0x47, 0x3e, 0xda, 0x45, 0x2a, 0x95, 0xaf, 0x9b, 0x2a,
	0xe3, 0x46, 0x53, 0x6e, 0x9c, 0x36, 0x6e, 0x7c, 0x52, 0x27, 0xd6, 0xc3, 0x95, 0x59, 0xc1, 0x79,
	0x31, 0x24, 0x96, 0x23, 0x22, 0xd2, 0xe1, 0xc1, 0xb8, 0x64, 0x36, 0xfd, 0x3f, 0xda, 0x17, 0x23,
	0xc2, 0x30, 0x60, 0xda, 0x34, 0xe8, 0x67, 0xfb, 0xf1, 0x24, 0x39, 0x2c, 0xef, 0x77, 0xed, 0x22,
	0xa9, 0x99, 0xab, 0x73, 0x82, 0x4c, 0x43, 0x7e, 0x41, 0xe2, 0x1f, 0xd8, 0x95, 0x92, 0x0a, 0x7d,
	0xc7, 0x3d, 0xfb, 0x03, 0xf7, 0x06, 0xfd, 0x72, 0xa4, 0x94, 0xd7, 0xaf, 0xef, 0xf3, 0xbd, 0x87,
	0xab, 0x6f, 0xa9, 0xae, 0x9e, 0x74, 0x65, 0x1a, 0x14, 0xcc, 0x2b, 0x83, 0x14, 0x50, 0x2a, 0x04,
	0xfa, 0x4b, 0xb1, 0x55, 0xd0, 0xd7, 0x73, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x87, 0x9b, 0xd1,
	0xd8, 0xfa, 0x01, 0x00, 0x00,
}
