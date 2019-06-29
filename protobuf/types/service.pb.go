// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/types/service.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Service represents the service's type.
type Service struct {
	Hash                 string                 `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	Sid                  string                 `protobuf:"bytes,12,opt,name=sid,proto3" json:"sid,omitempty"`
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Configuration        *Service_Configuration `protobuf:"bytes,8,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Tasks                []*Service_Task        `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Events               []*Service_Event       `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty"`
	Dependencies         []*Service_Dependency  `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Repository           string                 `protobuf:"bytes,9,opt,name=repository,proto3" json:"repository,omitempty"`
	Source               string                 `protobuf:"bytes,13,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Service) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service) GetConfiguration() *Service_Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *Service) GetTasks() []*Service_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Service) GetEvents() []*Service_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Service) GetDependencies() []*Service_Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Service) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *Service) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// Events are emitted by the service whenever the service wants.
type Service_Event struct {
	Key                  string               `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Data                 []*Service_Parameter `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Event) Reset()         { *m = Service_Event{} }
func (m *Service_Event) String() string { return proto.CompactTextString(m) }
func (*Service_Event) ProtoMessage()    {}
func (*Service_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 0}
}

func (m *Service_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Event.Unmarshal(m, b)
}
func (m *Service_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Event.Marshal(b, m, deterministic)
}
func (m *Service_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Event.Merge(m, src)
}
func (m *Service_Event) XXX_Size() int {
	return xxx_messageInfo_Service_Event.Size(m)
}
func (m *Service_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Event proto.InternalMessageInfo

func (m *Service_Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service_Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service_Event) GetData() []*Service_Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

// Task is a function that requires inputs and returns output.
type Service_Task struct {
	Key                  string               `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Inputs               []*Service_Parameter `protobuf:"bytes,6,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*Service_Parameter `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Task) Reset()         { *m = Service_Task{} }
func (m *Service_Task) String() string { return proto.CompactTextString(m) }
func (*Service_Task) ProtoMessage()    {}
func (*Service_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 1}
}

func (m *Service_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Task.Unmarshal(m, b)
}
func (m *Service_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Task.Marshal(b, m, deterministic)
}
func (m *Service_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Task.Merge(m, src)
}
func (m *Service_Task) XXX_Size() int {
	return xxx_messageInfo_Service_Task.Size(m)
}
func (m *Service_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Task proto.InternalMessageInfo

func (m *Service_Task) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service_Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service_Task) GetInputs() []*Service_Parameter {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Service_Task) GetOutputs() []*Service_Parameter {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// Parameter describes the task's inputs, the task's outputs, and the event's data.
type Service_Parameter struct {
	Key                  string               `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Optional             bool                 `protobuf:"varint,4,opt,name=optional,proto3" json:"optional,omitempty"`
	Repeated             bool                 `protobuf:"varint,9,opt,name=repeated,proto3" json:"repeated,omitempty"`
	Object               []*Service_Parameter `protobuf:"bytes,10,rep,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Parameter) Reset()         { *m = Service_Parameter{} }
func (m *Service_Parameter) String() string { return proto.CompactTextString(m) }
func (*Service_Parameter) ProtoMessage()    {}
func (*Service_Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 2}
}

func (m *Service_Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Parameter.Unmarshal(m, b)
}
func (m *Service_Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Parameter.Marshal(b, m, deterministic)
}
func (m *Service_Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Parameter.Merge(m, src)
}
func (m *Service_Parameter) XXX_Size() int {
	return xxx_messageInfo_Service_Parameter.Size(m)
}
func (m *Service_Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Parameter proto.InternalMessageInfo

func (m *Service_Parameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service_Parameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service_Parameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Service_Parameter) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *Service_Parameter) GetRepeated() bool {
	if m != nil {
		return m.Repeated
	}
	return false
}

func (m *Service_Parameter) GetObject() []*Service_Parameter {
	if m != nil {
		return m.Object
	}
	return nil
}

// A configuration is the configuration of the main container of the service's instance.
type Service_Configuration struct {
	Volumes              []string `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty"`
	VolumesFrom          []string `protobuf:"bytes,2,rep,name=volumesFrom,proto3" json:"volumesFrom,omitempty"`
	Ports                []string `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	Args                 []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	Command              string   `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	Env                  []string `protobuf:"bytes,6,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Configuration) Reset()         { *m = Service_Configuration{} }
func (m *Service_Configuration) String() string { return proto.CompactTextString(m) }
func (*Service_Configuration) ProtoMessage()    {}
func (*Service_Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 3}
}

func (m *Service_Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Configuration.Unmarshal(m, b)
}
func (m *Service_Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Configuration.Marshal(b, m, deterministic)
}
func (m *Service_Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Configuration.Merge(m, src)
}
func (m *Service_Configuration) XXX_Size() int {
	return xxx_messageInfo_Service_Configuration.Size(m)
}
func (m *Service_Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Configuration proto.InternalMessageInfo

func (m *Service_Configuration) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Service_Configuration) GetVolumesFrom() []string {
	if m != nil {
		return m.VolumesFrom
	}
	return nil
}

func (m *Service_Configuration) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Service_Configuration) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Service_Configuration) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Service_Configuration) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

// A dependency is a configuration of an other container that runs separately from the service.
type Service_Dependency struct {
	Key                  string   `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Volumes              []string `protobuf:"bytes,2,rep,name=volumes,proto3" json:"volumes,omitempty"`
	VolumesFrom          []string `protobuf:"bytes,3,rep,name=volumesFrom,proto3" json:"volumesFrom,omitempty"`
	Ports                []string `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	Args                 []string `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
	Command              string   `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	Env                  []string `protobuf:"bytes,9,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Dependency) Reset()         { *m = Service_Dependency{} }
func (m *Service_Dependency) String() string { return proto.CompactTextString(m) }
func (*Service_Dependency) ProtoMessage()    {}
func (*Service_Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_578ec16019661f8f, []int{0, 4}
}

func (m *Service_Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Dependency.Unmarshal(m, b)
}
func (m *Service_Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Dependency.Marshal(b, m, deterministic)
}
func (m *Service_Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Dependency.Merge(m, src)
}
func (m *Service_Dependency) XXX_Size() int {
	return xxx_messageInfo_Service_Dependency.Size(m)
}
func (m *Service_Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Dependency proto.InternalMessageInfo

func (m *Service_Dependency) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Service_Dependency) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Service_Dependency) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Service_Dependency) GetVolumesFrom() []string {
	if m != nil {
		return m.VolumesFrom
	}
	return nil
}

func (m *Service_Dependency) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Service_Dependency) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Service_Dependency) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Service_Dependency) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "types.Service")
	proto.RegisterType((*Service_Event)(nil), "types.Service.Event")
	proto.RegisterType((*Service_Task)(nil), "types.Service.Task")
	proto.RegisterType((*Service_Parameter)(nil), "types.Service.Parameter")
	proto.RegisterType((*Service_Configuration)(nil), "types.Service.Configuration")
	proto.RegisterType((*Service_Dependency)(nil), "types.Service.Dependency")
}

func init() { proto.RegisterFile("protobuf/types/service.proto", fileDescriptor_578ec16019661f8f) }

var fileDescriptor_578ec16019661f8f = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4b, 0x8e, 0xd3, 0x40,
	0x10, 0x95, 0xc7, 0x9f, 0xc4, 0x95, 0x89, 0x84, 0x9a, 0x08, 0x35, 0xd6, 0x08, 0x45, 0xac, 0x82,
	0x34, 0x38, 0xa3, 0xb0, 0x66, 0x33, 0x7c, 0xd6, 0xa8, 0x61, 0xc5, 0xae, 0x63, 0x57, 0x12, 0x93,
	0xb1, 0xdb, 0xea, 0x6e, 0x47, 0xca, 0x82, 0xc3, 0x70, 0x07, 0xce, 0xc0, 0x9e, 0x9b, 0x70, 0x04,
	0xd4, 0xe5, 0x38, 0x24, 0xd1, 0x30, 0x23, 0x3e, 0xbb, 0x7a, 0xaf, 0x5e, 0xa9, 0xea, 0xbd, 0x76,
	0x02, 0x17, 0xb5, 0x56, 0x56, 0xcd, 0x9b, 0xc5, 0xd4, 0x6e, 0x6b, 0x34, 0x53, 0x83, 0x7a, 0x53,
	0x64, 0x98, 0x12, 0xcd, 0x42, 0x22, 0x9f, 0xfe, 0x88, 0xa1, 0xf7, 0xbe, 0x6d, 0x30, 0x06, 0xc1,
	0x4a, 0x9a, 0x15, 0x87, 0xb1, 0x37, 0x89, 0x05, 0xd5, 0xec, 0x01, 0xf8, 0xa6, 0xc8, 0xf9, 0x39,
	0x51, 0xae, 0x74, 0xaa, 0x4a, 0x96, 0xc8, 0xbd, 0x56, 0xe5, 0x6a, 0x36, 0x86, 0x41, 0x8e, 0x26,
	0xd3, 0x45, 0x6d, 0x0b, 0x55, 0xf1, 0x33, 0x6a, 0x1d, 0x52, 0xec, 0x1a, 0x86, 0x99, 0xaa, 0x16,
	0xc5, 0xb2, 0xd1, 0x92, 0x34, 0xfd, 0xb1, 0x37, 0x19, 0xcc, 0x2e, 0x52, 0x3a, 0x23, 0xdd, 0x9d,
	0x90, 0xbe, 0x3a, 0xd4, 0x88, 0xe3, 0x11, 0xf6, 0x0c, 0x42, 0x2b, 0xcd, 0xda, 0xf0, 0x70, 0xec,
	0x4f, 0x06, 0xb3, 0x87, 0x27, 0xb3, 0x1f, 0xa4, 0x59, 0x8b, 0x56, 0xc1, 0x2e, 0x21, 0xc2, 0x0d,
	0x56, 0xd6, 0xf0, 0x88, 0xb4, 0xa3, 0x13, 0xed, 0x1b, 0xd7, 0x14, 0x3b, 0x0d, 0x7b, 0x09, 0xe7,
	0x39, 0xd6, 0x58, 0xe5, 0x58, 0x65, 0x05, 0x1a, 0xde, 0xa3, 0x99, 0xc7, 0x27, 0x33, 0xaf, 0x3b,
	0xc9, 0x56, 0x1c, 0xc9, 0xd9, 0x13, 0x00, 0x8d, 0xb5, 0x32, 0x85, 0x55, 0x7a, 0xcb, 0x63, 0x32,
	0x7f, 0xc0, 0xb0, 0x47, 0x10, 0x19, 0xd5, 0xe8, 0x0c, 0xf9, 0x90, 0x7a, 0x3b, 0x94, 0x7c, 0x86,
	0x90, 0xee, 0x70, 0x21, 0xaf, 0x71, 0xcb, 0x83, 0x36, 0xe4, 0x35, 0x6e, 0xff, 0x32, 0xe4, 0x4b,
	0x08, 0x72, 0x69, 0x25, 0xf7, 0xe9, 0x7e, 0x7e, 0x72, 0xff, 0x3b, 0xa9, 0x65, 0x89, 0x16, 0xb5,
	0x20, 0x55, 0xf2, 0xd5, 0x83, 0xc0, 0x65, 0xd6, 0xad, 0xef, 0xff, 0xeb, 0xfa, 0x2b, 0x88, 0x8a,
	0xaa, 0x6e, 0xf6, 0xa1, 0xff, 0xfe, 0x80, 0x9d, 0x8e, 0xcd, 0xa0, 0xa7, 0x1a, 0x4b, 0x23, 0xbd,
	0x7b, 0x46, 0x3a, 0x61, 0xf2, 0xdd, 0x83, 0x78, 0x4f, 0xff, 0xb7, 0xdb, 0x19, 0x04, 0x6e, 0x33,
	0xf7, 0xdb, 0x29, 0x57, 0xb3, 0x04, 0xfa, 0x8a, 0xba, 0xf2, 0x86, 0xde, 0xa6, 0x2f, 0xf6, 0xd8,
	0xf5, 0x34, 0xd6, 0x28, 0x2d, 0xe6, 0xf4, 0xe2, 0x7d, 0xb1, 0xc7, 0x2e, 0x07, 0x35, 0xff, 0x84,
	0x99, 0xe5, 0x70, 0x5f, 0x0e, 0xad, 0x2e, 0xf9, 0xe2, 0xc1, 0xf0, 0xe8, 0xd3, 0x67, 0x1c, 0x7a,
	0x1b, 0x75, 0xd3, 0x94, 0x68, 0xb8, 0x37, 0xf6, 0x27, 0xb1, 0xe8, 0xa0, 0xf3, 0xb2, 0x2b, 0xdf,
	0x6a, 0x55, 0xf2, 0x33, 0xea, 0x1e, 0x52, 0x6c, 0x04, 0x61, 0xad, 0xb4, 0x35, 0xf4, 0x1d, 0xc4,
	0xa2, 0x05, 0xce, 0xa1, 0xd4, 0x4b, 0xc3, 0x03, 0x22, 0xa9, 0x76, 0x5b, 0x32, 0x55, 0x96, 0xb2,
	0xca, 0x79, 0x48, 0xc6, 0x3b, 0xe8, 0x72, 0xc5, 0x6a, 0x43, 0x0f, 0x19, 0x0b, 0x57, 0x26, 0xdf,
	0x3c, 0x80, 0x5f, 0x3f, 0x81, 0x5b, 0x82, 0x1f, 0x41, 0x58, 0x94, 0x72, 0xd9, 0x25, 0xdf, 0x82,
	0x43, 0x23, 0x67, 0x77, 0x1a, 0xf1, 0xef, 0x30, 0x12, 0xdc, 0x66, 0x24, 0xfa, 0x13, 0x23, 0xf1,
	0xde, 0xc8, 0xf5, 0xd5, 0xc7, 0x74, 0x59, 0xd8, 0x55, 0x33, 0x4f, 0x33, 0x55, 0x4e, 0x4b, 0x34,
	0xcb, 0xe7, 0x0b, 0xd5, 0x54, 0x39, 0x05, 0x3f, 0xcd, 0x94, 0xc6, 0xe9, 0xf1, 0x3f, 0xe7, 0x3c,
	0x22, 0xfc, 0xe2, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x54, 0x02, 0x4b, 0x52, 0x05, 0x00,
	0x00,
}
