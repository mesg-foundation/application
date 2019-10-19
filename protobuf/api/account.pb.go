// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/api/account.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	account "github.com/mesg-foundation/engine/account"
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

// The request's data for the `Get` API.
type AccountServiceGetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountServiceGetRequest) Reset()         { *m = AccountServiceGetRequest{} }
func (m *AccountServiceGetRequest) String() string { return proto.CompactTextString(m) }
func (*AccountServiceGetRequest) ProtoMessage()    {}
func (*AccountServiceGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3259f8b8a943c47e, []int{0}
}
func (m *AccountServiceGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountServiceGetRequest.Unmarshal(m, b)
}
func (m *AccountServiceGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountServiceGetRequest.Marshal(b, m, deterministic)
}
func (m *AccountServiceGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountServiceGetRequest.Merge(m, src)
}
func (m *AccountServiceGetRequest) XXX_Size() int {
	return xxx_messageInfo_AccountServiceGetRequest.Size(m)
}
func (m *AccountServiceGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountServiceGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountServiceGetRequest proto.InternalMessageInfo

func (m *AccountServiceGetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AccountServiceGetResponse struct {
	Account              *account.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AccountServiceGetResponse) Reset()         { *m = AccountServiceGetResponse{} }
func (m *AccountServiceGetResponse) String() string { return proto.CompactTextString(m) }
func (*AccountServiceGetResponse) ProtoMessage()    {}
func (*AccountServiceGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3259f8b8a943c47e, []int{1}
}
func (m *AccountServiceGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountServiceGetResponse.Unmarshal(m, b)
}
func (m *AccountServiceGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountServiceGetResponse.Marshal(b, m, deterministic)
}
func (m *AccountServiceGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountServiceGetResponse.Merge(m, src)
}
func (m *AccountServiceGetResponse) XXX_Size() int {
	return xxx_messageInfo_AccountServiceGetResponse.Size(m)
}
func (m *AccountServiceGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountServiceGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountServiceGetResponse proto.InternalMessageInfo

func (m *AccountServiceGetResponse) GetAccount() *account.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

// The request's data for the `List` API.
type AccountServiceListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountServiceListRequest) Reset()         { *m = AccountServiceListRequest{} }
func (m *AccountServiceListRequest) String() string { return proto.CompactTextString(m) }
func (*AccountServiceListRequest) ProtoMessage()    {}
func (*AccountServiceListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3259f8b8a943c47e, []int{2}
}
func (m *AccountServiceListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountServiceListRequest.Unmarshal(m, b)
}
func (m *AccountServiceListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountServiceListRequest.Marshal(b, m, deterministic)
}
func (m *AccountServiceListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountServiceListRequest.Merge(m, src)
}
func (m *AccountServiceListRequest) XXX_Size() int {
	return xxx_messageInfo_AccountServiceListRequest.Size(m)
}
func (m *AccountServiceListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountServiceListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountServiceListRequest proto.InternalMessageInfo

// The response's data for the `List` API.
type AccountServiceListResponse struct {
	Accounts             []*account.Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountServiceListResponse) Reset()         { *m = AccountServiceListResponse{} }
func (m *AccountServiceListResponse) String() string { return proto.CompactTextString(m) }
func (*AccountServiceListResponse) ProtoMessage()    {}
func (*AccountServiceListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3259f8b8a943c47e, []int{3}
}
func (m *AccountServiceListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountServiceListResponse.Unmarshal(m, b)
}
func (m *AccountServiceListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountServiceListResponse.Marshal(b, m, deterministic)
}
func (m *AccountServiceListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountServiceListResponse.Merge(m, src)
}
func (m *AccountServiceListResponse) XXX_Size() int {
	return xxx_messageInfo_AccountServiceListResponse.Size(m)
}
func (m *AccountServiceListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountServiceListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountServiceListResponse proto.InternalMessageInfo

func (m *AccountServiceListResponse) GetAccounts() []*account.Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// The request's data for the `Create` API.
type AccountServiceCreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountServiceCreateRequest) Reset()         { *m = AccountServiceCreateRequest{} }
func (m *AccountServiceCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AccountServiceCreateRequest) ProtoMessage()    {}
func (*AccountServiceCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3259f8b8a943c47e, []int{4}
}
func (m *AccountServiceCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountServiceCreateRequest.Unmarshal(m, b)
}
func (m *AccountServiceCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountServiceCreateRequest.Marshal(b, m, deterministic)
}
func (m *AccountServiceCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountServiceCreateRequest.Merge(m, src)
}
func (m *AccountServiceCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AccountServiceCreateRequest.Size(m)
}
func (m *AccountServiceCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountServiceCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountServiceCreateRequest proto.InternalMessageInfo

func (m *AccountServiceCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountServiceCreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The response's data for the `Create` API.
type AccountServiceCreateResponse struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Mnemonic             string   `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountServiceCreateResponse) Reset()         { *m = AccountServiceCreateResponse{} }
func (m *AccountServiceCreateResponse) String() string { return proto.CompactTextString(m) }
func (*AccountServiceCreateResponse) ProtoMessage()    {}
func (*AccountServiceCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3259f8b8a943c47e, []int{5}
}
func (m *AccountServiceCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountServiceCreateResponse.Unmarshal(m, b)
}
func (m *AccountServiceCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountServiceCreateResponse.Marshal(b, m, deterministic)
}
func (m *AccountServiceCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountServiceCreateResponse.Merge(m, src)
}
func (m *AccountServiceCreateResponse) XXX_Size() int {
	return xxx_messageInfo_AccountServiceCreateResponse.Size(m)
}
func (m *AccountServiceCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountServiceCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountServiceCreateResponse proto.InternalMessageInfo

func (m *AccountServiceCreateResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountServiceCreateResponse) GetMnemonic() string {
	if m != nil {
		return m.Mnemonic
	}
	return ""
}

// The request's data for the `Delete` API.
type AccountServiceDeleteRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountServiceDeleteRequest) Reset()         { *m = AccountServiceDeleteRequest{} }
func (m *AccountServiceDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AccountServiceDeleteRequest) ProtoMessage()    {}
func (*AccountServiceDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3259f8b8a943c47e, []int{6}
}
func (m *AccountServiceDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountServiceDeleteRequest.Unmarshal(m, b)
}
func (m *AccountServiceDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountServiceDeleteRequest.Marshal(b, m, deterministic)
}
func (m *AccountServiceDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountServiceDeleteRequest.Merge(m, src)
}
func (m *AccountServiceDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_AccountServiceDeleteRequest.Size(m)
}
func (m *AccountServiceDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountServiceDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountServiceDeleteRequest proto.InternalMessageInfo

// The response's data for the `Delete` API.
type AccountServiceDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountServiceDeleteResponse) Reset()         { *m = AccountServiceDeleteResponse{} }
func (m *AccountServiceDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*AccountServiceDeleteResponse) ProtoMessage()    {}
func (*AccountServiceDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3259f8b8a943c47e, []int{7}
}
func (m *AccountServiceDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountServiceDeleteResponse.Unmarshal(m, b)
}
func (m *AccountServiceDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountServiceDeleteResponse.Marshal(b, m, deterministic)
}
func (m *AccountServiceDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountServiceDeleteResponse.Merge(m, src)
}
func (m *AccountServiceDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_AccountServiceDeleteResponse.Size(m)
}
func (m *AccountServiceDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountServiceDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountServiceDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AccountServiceGetRequest)(nil), "mesg.api.v1.AccountServiceGetRequest")
	proto.RegisterType((*AccountServiceGetResponse)(nil), "mesg.api.v1.AccountServiceGetResponse")
	proto.RegisterType((*AccountServiceListRequest)(nil), "mesg.api.v1.AccountServiceListRequest")
	proto.RegisterType((*AccountServiceListResponse)(nil), "mesg.api.v1.AccountServiceListResponse")
	proto.RegisterType((*AccountServiceCreateRequest)(nil), "mesg.api.v1.AccountServiceCreateRequest")
	proto.RegisterType((*AccountServiceCreateResponse)(nil), "mesg.api.v1.AccountServiceCreateResponse")
	proto.RegisterType((*AccountServiceDeleteRequest)(nil), "mesg.api.v1.AccountServiceDeleteRequest")
	proto.RegisterType((*AccountServiceDeleteResponse)(nil), "mesg.api.v1.AccountServiceDeleteResponse")
}

func init() { proto.RegisterFile("protobuf/api/account.proto", fileDescriptor_3259f8b8a943c47e) }

var fileDescriptor_3259f8b8a943c47e = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4d, 0x4f, 0x3a, 0x31,
	0x10, 0xc6, 0xe1, 0xbf, 0xfc, 0x01, 0x87, 0xc4, 0x43, 0x0f, 0x66, 0x2d, 0x68, 0x48, 0x13, 0x15,
	0x2f, 0x45, 0xf0, 0x13, 0xf8, 0x92, 0x70, 0x91, 0xc4, 0xa0, 0x31, 0xd1, 0x8b, 0x29, 0xcb, 0x48,
	0x36, 0x71, 0xb7, 0x75, 0x5b, 0x30, 0x7e, 0x00, 0xbf, 0xb7, 0xa1, 0xfb, 0x12, 0x96, 0x60, 0xf1,
	0xd6, 0x69, 0x7f, 0xf3, 0x3c, 0x0f, 0x33, 0x2c, 0x50, 0x95, 0x48, 0x23, 0xa7, 0x8b, 0xb7, 0xbe,
	0x50, 0x61, 0x5f, 0x04, 0x81, 0x5c, 0xc4, 0x86, 0xdb, 0x4b, 0xd2, 0x8a, 0x50, 0xcf, 0xb9, 0x50,
	0x21, 0x5f, 0x0e, 0x68, 0xa7, 0x00, 0xcd, 0x97, 0x42, 0x5d, 0x46, 0x29, 0x9b, 0xcb, 0xb9, 0xec,
	0x17, 0xc8, 0xaa, 0xb2, 0x85, 0x3d, 0xa5, 0x0c, 0xe3, 0xe0, 0x5f, 0xa5, 0x4d, 0x0f, 0x98, 0x2c,
	0xc3, 0x00, 0x47, 0x68, 0x26, 0xf8, 0xb1, 0x40, 0x6d, 0x08, 0x81, 0x5a, 0x2c, 0x22, 0xf4, 0xab,
	0xdd, 0x6a, 0x6f, 0x6f, 0x62, 0xcf, 0x6c, 0x0c, 0x87, 0x5b, 0x78, 0xad, 0x64, 0xac, 0x91, 0x5c,
	0x40, 0x23, 0x4b, 0x60, 0x7b, 0x5a, 0xc3, 0x03, 0x6e, 0xd3, 0xda, 0x70, 0x7c, 0x39, 0xe0, 0x59,
	0xeb, 0x24, 0xc7, 0x58, 0x7b, 0x53, 0xee, 0x2e, 0xd4, 0xb9, 0x3f, 0xbb, 0x07, 0xba, 0xed, 0x31,
	0x33, 0x1b, 0x42, 0x33, 0x53, 0xd1, 0x7e, 0xb5, 0xeb, 0x39, 0xdc, 0x0a, 0x8e, 0x8d, 0xa1, 0x5d,
	0x56, 0xbc, 0x49, 0x50, 0x18, 0x74, 0xfc, 0x60, 0x42, 0xa1, 0xa9, 0x84, 0xd6, 0x9f, 0x32, 0x99,
	0xf9, 0xff, 0xec, 0x7d, 0x51, 0xb3, 0x47, 0xe8, 0x6c, 0x97, 0xcb, 0x22, 0xfa, 0xd0, 0x10, 0xb3,
	0x59, 0x82, 0x5a, 0x67, 0x92, 0x79, 0xb9, 0x52, 0x8d, 0x62, 0x8c, 0x64, 0x1c, 0x06, 0xb9, 0x6a,
	0x5e, 0xb3, 0xa3, 0xcd, 0x90, 0xb7, 0xf8, 0x8e, 0x45, 0x48, 0x76, 0xbc, 0x69, 0x9a, 0x3f, 0xa7,
	0xa6, 0xc3, 0x6f, 0x0f, 0xf6, 0xcb, 0x00, 0x79, 0x02, 0x6f, 0x84, 0x86, 0x9c, 0xf0, 0xb5, 0xff,
	0x0e, 0xff, 0x6d, 0xed, 0xf4, 0x74, 0x17, 0x96, 0x1a, 0xb1, 0x0a, 0x79, 0x86, 0xda, 0x6a, 0x25,
	0xc4, 0xd5, 0xb1, 0xb6, 0x50, 0x7a, 0xb6, 0x93, 0x2b, 0xa4, 0x5f, 0xa1, 0x9e, 0x0e, 0x93, 0xf4,
	0x1c, 0x4d, 0xa5, 0xf5, 0xd1, 0xf3, 0x3f, 0x90, 0xeb, 0x06, 0xe9, 0xe0, 0x9c, 0x06, 0xa5, 0xd1,
	0x3b, 0x0d, 0xca, 0x5b, 0x60, 0x95, 0xeb, 0xff, 0x2f, 0x9e, 0x50, 0xe1, 0xb4, 0x6e, 0xbf, 0xb3,
	0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0xc2, 0x42, 0xc0, 0xd4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	// Get returns an Account matching the criteria of the request.
	Get(ctx context.Context, in *AccountServiceGetRequest, opts ...grpc.CallOption) (*AccountServiceGetResponse, error)
	// List returns all Accounts matching the criteria of the request.
	List(ctx context.Context, in *AccountServiceListRequest, opts ...grpc.CallOption) (*AccountServiceListResponse, error)
	// Create an Account with a name and password.
	Create(ctx context.Context, in *AccountServiceCreateRequest, opts ...grpc.CallOption) (*AccountServiceCreateResponse, error)
	// Delete an Account.
	Delete(ctx context.Context, in *AccountServiceDeleteRequest, opts ...grpc.CallOption) (*AccountServiceDeleteResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Get(ctx context.Context, in *AccountServiceGetRequest, opts ...grpc.CallOption) (*AccountServiceGetResponse, error) {
	out := new(AccountServiceGetResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.AccountService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) List(ctx context.Context, in *AccountServiceListRequest, opts ...grpc.CallOption) (*AccountServiceListResponse, error) {
	out := new(AccountServiceListResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.AccountService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Create(ctx context.Context, in *AccountServiceCreateRequest, opts ...grpc.CallOption) (*AccountServiceCreateResponse, error) {
	out := new(AccountServiceCreateResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.AccountService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Delete(ctx context.Context, in *AccountServiceDeleteRequest, opts ...grpc.CallOption) (*AccountServiceDeleteResponse, error) {
	out := new(AccountServiceDeleteResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.v1.AccountService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	// Get returns an Account matching the criteria of the request.
	Get(context.Context, *AccountServiceGetRequest) (*AccountServiceGetResponse, error)
	// List returns all Accounts matching the criteria of the request.
	List(context.Context, *AccountServiceListRequest) (*AccountServiceListResponse, error)
	// Create an Account with a name and password.
	Create(context.Context, *AccountServiceCreateRequest) (*AccountServiceCreateResponse, error)
	// Delete an Account.
	Delete(context.Context, *AccountServiceDeleteRequest) (*AccountServiceDeleteResponse, error)
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) Get(ctx context.Context, req *AccountServiceGetRequest) (*AccountServiceGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccountServiceServer) List(ctx context.Context, req *AccountServiceListRequest) (*AccountServiceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAccountServiceServer) Create(ctx context.Context, req *AccountServiceCreateRequest) (*AccountServiceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccountServiceServer) Delete(ctx context.Context, req *AccountServiceDeleteRequest) (*AccountServiceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountServiceGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.AccountService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Get(ctx, req.(*AccountServiceGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.AccountService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).List(ctx, req.(*AccountServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.AccountService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Create(ctx, req.(*AccountServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountServiceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.v1.AccountService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Delete(ctx, req.(*AccountServiceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mesg.api.v1.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AccountService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AccountService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccountService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccountService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/account.proto",
}
