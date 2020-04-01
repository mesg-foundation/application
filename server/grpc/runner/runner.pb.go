// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/grpc/runner/runner.proto

package runner

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	execution "github.com/mesg-foundation/engine/execution"
	github_com_mesg_foundation_engine_hash "github.com/mesg-foundation/engine/hash"
	types "github.com/mesg-foundation/engine/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// RegisterRequest is the request of the endpoint Register.
type RegisterRequest struct {
	// payload to use to register this runner. it should be the content of the env variable MESG_REGISTER_PAYLOAD.
	Payload              string   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{0}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

// // RegisterRequest is the request of the endpoint Register.
type RegisterRequestPayload struct {
	// signature of the value to verify authenticity.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty" validate:"required"`
	// value contains what the engine actually needs to register the runner.
	Value                RegisterRequestPayload_Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value" validate:"required"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RegisterRequestPayload) Reset()         { *m = RegisterRequestPayload{} }
func (m *RegisterRequestPayload) String() string { return proto.CompactTextString(m) }
func (*RegisterRequestPayload) ProtoMessage()    {}
func (*RegisterRequestPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{1}
}
func (m *RegisterRequestPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequestPayload.Unmarshal(m, b)
}
func (m *RegisterRequestPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequestPayload.Marshal(b, m, deterministic)
}
func (m *RegisterRequestPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequestPayload.Merge(m, src)
}
func (m *RegisterRequestPayload) XXX_Size() int {
	return xxx_messageInfo_RegisterRequestPayload.Size(m)
}
func (m *RegisterRequestPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequestPayload.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequestPayload proto.InternalMessageInfo

func (m *RegisterRequestPayload) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *RegisterRequestPayload) GetValue() RegisterRequestPayload_Value {
	if m != nil {
		return m.Value
	}
	return RegisterRequestPayload_Value{}
}

// Value message
type RegisterRequestPayload_Value struct {
	// Service's hash to start the runner with.
	ServiceHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=serviceHash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"serviceHash,omitempty" validate:"required,hash"`
	// Hash of the environmental variables to start the runner with.
	EnvHash              github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,2,opt,name=envHash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"envHash,omitempty" validate:"omitempty,hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *RegisterRequestPayload_Value) Reset()         { *m = RegisterRequestPayload_Value{} }
func (m *RegisterRequestPayload_Value) String() string { return proto.CompactTextString(m) }
func (*RegisterRequestPayload_Value) ProtoMessage()    {}
func (*RegisterRequestPayload_Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{1, 0}
}
func (m *RegisterRequestPayload_Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequestPayload_Value.Unmarshal(m, b)
}
func (m *RegisterRequestPayload_Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequestPayload_Value.Marshal(b, m, deterministic)
}
func (m *RegisterRequestPayload_Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequestPayload_Value.Merge(m, src)
}
func (m *RegisterRequestPayload_Value) XXX_Size() int {
	return xxx_messageInfo_RegisterRequestPayload_Value.Size(m)
}
func (m *RegisterRequestPayload_Value) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequestPayload_Value.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequestPayload_Value proto.InternalMessageInfo

func (m *RegisterRequestPayload_Value) GetServiceHash() github_com_mesg_foundation_engine_hash.Hash {
	if m != nil {
		return m.ServiceHash
	}
	return nil
}

func (m *RegisterRequestPayload_Value) GetEnvHash() github_com_mesg_foundation_engine_hash.Hash {
	if m != nil {
		return m.EnvHash
	}
	return nil
}

// RegisterResponse is the response of the endpoint Register.
type RegisterResponse struct {
	// token to use with the other endpoints of this API.
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{2}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// ExecutionRequest is the request of the endpoint Execution.
type ExecutionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionRequest) Reset()         { *m = ExecutionRequest{} }
func (m *ExecutionRequest) String() string { return proto.CompactTextString(m) }
func (*ExecutionRequest) ProtoMessage()    {}
func (*ExecutionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{3}
}
func (m *ExecutionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionRequest.Unmarshal(m, b)
}
func (m *ExecutionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionRequest.Marshal(b, m, deterministic)
}
func (m *ExecutionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionRequest.Merge(m, src)
}
func (m *ExecutionRequest) XXX_Size() int {
	return xxx_messageInfo_ExecutionRequest.Size(m)
}
func (m *ExecutionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionRequest proto.InternalMessageInfo

// EventRequest is the request of the endpoint Result.
type ResultRequest struct {
	// Execution's hash of the executed execution.
	ExecutionHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=executionHash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"executionHash,omitempty" validate:"required,hash"`
	// Execution's result.
	//
	// Types that are valid to be assigned to Result:
	//	*ResultRequest_Outputs
	//	*ResultRequest_Error
	Result               isResultRequest_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ResultRequest) Reset()         { *m = ResultRequest{} }
func (m *ResultRequest) String() string { return proto.CompactTextString(m) }
func (*ResultRequest) ProtoMessage()    {}
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{4}
}
func (m *ResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultRequest.Unmarshal(m, b)
}
func (m *ResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultRequest.Marshal(b, m, deterministic)
}
func (m *ResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultRequest.Merge(m, src)
}
func (m *ResultRequest) XXX_Size() int {
	return xxx_messageInfo_ResultRequest.Size(m)
}
func (m *ResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResultRequest proto.InternalMessageInfo

type isResultRequest_Result interface {
	isResultRequest_Result()
}

type ResultRequest_Outputs struct {
	Outputs *types.Struct `protobuf:"bytes,2,opt,name=outputs,proto3,oneof" json:"outputs,omitempty"`
}
type ResultRequest_Error struct {
	Error string `protobuf:"bytes,3,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (*ResultRequest_Outputs) isResultRequest_Result() {}
func (*ResultRequest_Error) isResultRequest_Result()   {}

func (m *ResultRequest) GetResult() isResultRequest_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ResultRequest) GetExecutionHash() github_com_mesg_foundation_engine_hash.Hash {
	if m != nil {
		return m.ExecutionHash
	}
	return nil
}

func (m *ResultRequest) GetOutputs() *types.Struct {
	if x, ok := m.GetResult().(*ResultRequest_Outputs); ok {
		return x.Outputs
	}
	return nil
}

func (m *ResultRequest) GetError() string {
	if x, ok := m.GetResult().(*ResultRequest_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResultRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResultRequest_Outputs)(nil),
		(*ResultRequest_Error)(nil),
	}
}

// ResultResponse is the response of the endpoint Result.
type ResultResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultResponse) Reset()         { *m = ResultResponse{} }
func (m *ResultResponse) String() string { return proto.CompactTextString(m) }
func (*ResultResponse) ProtoMessage()    {}
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{5}
}
func (m *ResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultResponse.Unmarshal(m, b)
}
func (m *ResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultResponse.Marshal(b, m, deterministic)
}
func (m *ResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultResponse.Merge(m, src)
}
func (m *ResultResponse) XXX_Size() int {
	return xxx_messageInfo_ResultResponse.Size(m)
}
func (m *ResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResultResponse proto.InternalMessageInfo

// EventRequest is the request of the endpoint Event.
type EventRequest struct {
	// Event's key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" validate:"required,printascii"`
	// Event's data
	Data                 *types.Struct `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{6}
}
func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EventRequest) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

// EventResponse is the response of the endpoint Event.
type EventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1650fde6416bd437, []int{7}
}
func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "mesg.grpc.runner.RegisterRequest")
	proto.RegisterType((*RegisterRequestPayload)(nil), "mesg.grpc.runner.RegisterRequestPayload")
	proto.RegisterType((*RegisterRequestPayload_Value)(nil), "mesg.grpc.runner.RegisterRequestPayload.Value")
	proto.RegisterType((*RegisterResponse)(nil), "mesg.grpc.runner.RegisterResponse")
	proto.RegisterType((*ExecutionRequest)(nil), "mesg.grpc.runner.ExecutionRequest")
	proto.RegisterType((*ResultRequest)(nil), "mesg.grpc.runner.ResultRequest")
	proto.RegisterType((*ResultResponse)(nil), "mesg.grpc.runner.ResultResponse")
	proto.RegisterType((*EventRequest)(nil), "mesg.grpc.runner.EventRequest")
	proto.RegisterType((*EventResponse)(nil), "mesg.grpc.runner.EventResponse")
}

func init() { proto.RegisterFile("server/grpc/runner/runner.proto", fileDescriptor_1650fde6416bd437) }

var fileDescriptor_1650fde6416bd437 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x4d, 0xda, 0x2f, 0x69, 0x7b, 0xdb, 0x7e, 0x8d, 0x06, 0x28, 0xc1, 0x95, 0xea, 0x32, 0xab,
	0x22, 0xc0, 0x86, 0x42, 0x85, 0xc4, 0x32, 0xa2, 0x52, 0xc5, 0x8f, 0x84, 0x5c, 0x89, 0x05, 0xbb,
	0xa9, 0x73, 0xeb, 0x58, 0x4d, 0x66, 0x9c, 0xf9, 0x89, 0xc8, 0x73, 0xf0, 0x4a, 0x2c, 0x78, 0x03,
	0x76, 0x7e, 0x88, 0x2c, 0x59, 0x21, 0xcf, 0xd8, 0x49, 0x08, 0x49, 0x2b, 0x21, 0xb1, 0x4a, 0x26,
	0xe7, 0xcc, 0x3d, 0xf7, 0x9e, 0x7b, 0x32, 0xe0, 0x2b, 0x94, 0x23, 0x94, 0x61, 0x22, 0xb3, 0x38,
	0x94, 0x86, 0x73, 0x94, 0xe5, 0x47, 0x90, 0x49, 0xa1, 0x05, 0x69, 0x0d, 0x50, 0x25, 0x41, 0x01,
	0x07, 0xee, 0x77, 0x8f, 0x26, 0x22, 0x11, 0xa1, 0x45, 0x2f, 0xcd, 0x55, 0x58, 0x9c, 0xec, 0xc1,
	0x7e, 0x73, 0xb7, 0xbc, 0xc3, 0x29, 0xac, 0xc7, 0x19, 0xaa, 0x10, 0xbf, 0x60, 0x6c, 0x74, 0x2a,
	0x78, 0x89, 0x1f, 0x2c, 0xe0, 0x4a, 0x4b, 0x13, 0x6b, 0x07, 0xd2, 0x37, 0xb0, 0x17, 0x61, 0x92,
	0x2a, 0x8d, 0x32, 0xc2, 0xa1, 0x41, 0xa5, 0xc9, 0x73, 0xd8, 0xc8, 0xd8, 0xb8, 0x2f, 0x58, 0xb7,
	0x5d, 0x3f, 0xaa, 0x1f, 0x6f, 0x75, 0xee, 0x4f, 0x72, 0xff, 0xce, 0x88, 0xf5, 0xd3, 0x2e, 0xd3,
	0xf8, 0x9a, 0x4a, 0x1c, 0x9a, 0x54, 0x62, 0x97, 0x46, 0x15, 0x8f, 0x7e, 0x5d, 0x87, 0xfd, 0x85,
	0x32, 0x1f, 0x1d, 0x44, 0x4e, 0x61, 0x4b, 0xa5, 0x09, 0x67, 0xda, 0x48, 0xb4, 0xf5, 0x76, 0x56,
	0xd7, 0x9b, 0x31, 0x09, 0x83, 0xc6, 0x88, 0xf5, 0x0d, 0xb6, 0xd7, 0x8e, 0xea, 0xc7, 0xdb, 0x27,
	0x41, 0xb0, 0x68, 0x4d, 0xb0, 0x5c, 0x2f, 0xf8, 0x54, 0xdc, 0xea, 0x1c, 0x7c, 0xcf, 0xfd, 0xda,
	0x2a, 0x19, 0x57, 0xd9, 0xcb, 0xeb, 0xd0, 0xb0, 0x6c, 0x32, 0x80, 0xed, 0x62, 0x35, 0x69, 0x8c,
	0xe7, 0x4c, 0xf5, 0xca, 0x2e, 0xdf, 0x4d, 0x72, 0xbf, 0xfd, 0xe7, 0xf5, 0x27, 0x3d, 0xa6, 0x7a,
	0xf4, 0x67, 0xee, 0x3f, 0x4e, 0x52, 0xdd, 0x33, 0x97, 0x41, 0x2c, 0x06, 0x61, 0xd1, 0xdc, 0xd3,
	0x2b, 0x61, 0x78, 0x97, 0x15, 0xc6, 0x87, 0xc8, 0x93, 0x94, 0x63, 0x58, 0x50, 0x83, 0xa2, 0x64,
	0x34, 0x5f, 0x9f, 0x5c, 0xc1, 0x06, 0xf2, 0x91, 0x95, 0x5a, 0xb3, 0x52, 0xef, 0x27, 0xb9, 0xff,
	0x60, 0x26, 0x25, 0x06, 0xa9, 0xc6, 0x41, 0xa6, 0xc7, 0x7f, 0xa7, 0x55, 0x15, 0xa7, 0xc7, 0xd0,
	0x9a, 0x99, 0xa4, 0x32, 0xc1, 0x15, 0x92, 0xbb, 0xd0, 0xd0, 0xe2, 0x1a, 0xb9, 0x5b, 0x6d, 0xe4,
	0x0e, 0x94, 0x40, 0xeb, 0xac, 0x4a, 0x4d, 0xe9, 0x27, 0xfd, 0x51, 0x87, 0xdd, 0x08, 0x95, 0xe9,
	0xeb, 0x2a, 0x18, 0x43, 0xd8, 0x9d, 0x66, 0xeb, 0x5f, 0x19, 0xf5, 0xbb, 0x42, 0x91, 0x45, 0x61,
	0x74, 0x66, 0xb4, 0x2a, 0x83, 0x70, 0xcf, 0x05, 0xa1, 0x8a, 0x74, 0x70, 0x61, 0xc3, 0x7c, 0x5e,
	0x8b, 0x2a, 0x1e, 0xd9, 0x87, 0x06, 0x4a, 0x29, 0x64, 0x7b, 0xbd, 0x98, 0xf0, 0xbc, 0x16, 0xb9,
	0x63, 0x67, 0x13, 0x9a, 0xd2, 0x8e, 0x43, 0x5b, 0xf0, 0x7f, 0x35, 0x98, 0x73, 0x85, 0x0a, 0xd8,
	0x39, 0x1b, 0x21, 0x9f, 0x4e, 0xfa, 0x12, 0xd6, 0xaf, 0x71, 0x5c, 0xc6, 0x9f, 0x4e, 0x72, 0xff,
	0x70, 0xc9, 0x7c, 0x99, 0x4c, 0xb9, 0x66, 0x2a, 0x4e, 0x53, 0x1a, 0x15, 0x74, 0xf2, 0x08, 0xfe,
	0xeb, 0x32, 0xcd, 0x6e, 0xec, 0x34, 0xb2, 0x14, 0xba, 0x07, 0xbb, 0xa5, 0xa0, 0xeb, 0xe0, 0xe4,
	0xdb, 0x1a, 0x34, 0x23, 0x1b, 0x6c, 0x72, 0x01, 0x9b, 0xd5, 0xda, 0xc8, 0xc3, 0x5b, 0x73, 0xef,
	0xd1, 0x9b, 0x28, 0xe5, 0x7c, 0x35, 0xf2, 0x16, 0xb6, 0xa6, 0x1b, 0x26, 0x4b, 0xae, 0x2c, 0xae,
	0xdf, 0x2b, 0xdb, 0xb7, 0x4f, 0xc6, 0x0c, 0xa5, 0xb5, 0x67, 0x75, 0xf2, 0x01, 0x9a, 0xce, 0x3f,
	0xe2, 0x2f, 0xd3, 0x9e, 0x8b, 0x8c, 0x77, 0xb4, 0x9a, 0x30, 0xd7, 0x5a, 0xc3, 0x7a, 0x41, 0x0e,
	0x97, 0xb4, 0x35, 0xb7, 0x15, 0xcf, 0x5f, 0x89, 0x57, 0xb5, 0x3a, 0xaf, 0x3e, 0x9f, 0xde, 0x9e,
	0xb6, 0xf2, 0xbf, 0x38, 0xff, 0x0e, 0x5f, 0x36, 0xed, 0x9e, 0x5e, 0xfc, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x8b, 0xe0, 0xc0, 0xa4, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RunnerClient is the client API for Runner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RunnerClient interface {
	// Register registers a new runner to the Engine.
	// This endpoint should only be called when the runner is ready to receive execution and emit events.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Execution returns a stream of executions that contains the Execution the Runner must execute.
	Execution(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (Runner_ExecutionClient, error)
	// Result should be used to return the result of an Execution.
	Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error)
	// Event should be used to emits an event.
	Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type runnerClient struct {
	cc *grpc.ClientConn
}

func NewRunnerClient(cc *grpc.ClientConn) RunnerClient {
	return &runnerClient{cc}
}

func (c *runnerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/mesg.grpc.runner.Runner/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Execution(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (Runner_ExecutionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Runner_serviceDesc.Streams[0], "/mesg.grpc.runner.Runner/Execution", opts...)
	if err != nil {
		return nil, err
	}
	x := &runnerExecutionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Runner_ExecutionClient interface {
	Recv() (*execution.Execution, error)
	grpc.ClientStream
}

type runnerExecutionClient struct {
	grpc.ClientStream
}

func (x *runnerExecutionClient) Recv() (*execution.Execution, error) {
	m := new(execution.Execution)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runnerClient) Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/mesg.grpc.runner.Runner/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runnerClient) Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/mesg.grpc.runner.Runner/Event", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerServer is the server API for Runner service.
type RunnerServer interface {
	// Register registers a new runner to the Engine.
	// This endpoint should only be called when the runner is ready to receive execution and emit events.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Execution returns a stream of executions that contains the Execution the Runner must execute.
	Execution(*ExecutionRequest, Runner_ExecutionServer) error
	// Result should be used to return the result of an Execution.
	Result(context.Context, *ResultRequest) (*ResultResponse, error)
	// Event should be used to emits an event.
	Event(context.Context, *EventRequest) (*EventResponse, error)
}

// UnimplementedRunnerServer can be embedded to have forward compatible implementations.
type UnimplementedRunnerServer struct {
}

func (*UnimplementedRunnerServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedRunnerServer) Execution(req *ExecutionRequest, srv Runner_ExecutionServer) error {
	return status.Errorf(codes.Unimplemented, "method Execution not implemented")
}
func (*UnimplementedRunnerServer) Result(ctx context.Context, req *ResultRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (*UnimplementedRunnerServer) Event(ctx context.Context, req *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Event not implemented")
}

func RegisterRunnerServer(s *grpc.Server, srv RunnerServer) {
	s.RegisterService(&_Runner_serviceDesc, srv)
}

func _Runner_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.grpc.runner.Runner/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Execution_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecutionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RunnerServer).Execution(m, &runnerExecutionServer{stream})
}

type Runner_ExecutionServer interface {
	Send(*execution.Execution) error
	grpc.ServerStream
}

type runnerExecutionServer struct {
	grpc.ServerStream
}

func (x *runnerExecutionServer) Send(m *execution.Execution) error {
	return x.ServerStream.SendMsg(m)
}

func _Runner_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.grpc.runner.Runner/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Result(ctx, req.(*ResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Runner_Event_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServer).Event(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.grpc.runner.Runner/Event",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServer).Event(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Runner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mesg.grpc.runner.Runner",
	HandlerType: (*RunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Runner_Register_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _Runner_Result_Handler,
		},
		{
			MethodName: "Event",
			Handler:    _Runner_Event_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execution",
			Handler:       _Runner_Execution_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/grpc/runner/runner.proto",
}
