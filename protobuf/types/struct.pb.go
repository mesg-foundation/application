// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/types/struct.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
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

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//  The JSON representation for `NullValue` is JSON `null`.
type NullValue int32

const (
	// Null value.
	NullValue_NULL_VALUE NullValue = 0
)

var NullValue_name = map[int32]string{
	0: "NULL_VALUE",
}

var NullValue_value = map[string]int32{
	"NULL_VALUE": 0,
}

func (x NullValue) String() string {
	return proto.EnumName(NullValue_name, int32(x))
}

func (NullValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84428fcddb6debf8, []int{0}
}

// `Struct` represents a structured data value, consisting of fields
// which map to dynamically typed values. In some languages, `Struct`
// might be supported by a native representation. For example, in
// scripting languages like JS a struct is represented as an
// object. The details of that representation are described together
// with the proto support for the language.
//
// The JSON representation for `Struct` is JSON object.
type Struct struct {
	// Unordered map of dynamically typed values.
	Fields               map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" hash:"name:1" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Struct) Reset()         { *m = Struct{} }
func (m *Struct) String() string { return proto.CompactTextString(m) }
func (*Struct) ProtoMessage()    {}
func (*Struct) Descriptor() ([]byte, []int) {
	return fileDescriptor_84428fcddb6debf8, []int{0}
}
func (m *Struct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Struct.Unmarshal(m, b)
}
func (m *Struct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Struct.Marshal(b, m, deterministic)
}
func (m *Struct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Struct.Merge(m, src)
}
func (m *Struct) XXX_Size() int {
	return xxx_messageInfo_Struct.Size(m)
}
func (m *Struct) XXX_DiscardUnknown() {
	xxx_messageInfo_Struct.DiscardUnknown(m)
}

var xxx_messageInfo_Struct proto.InternalMessageInfo

func (m *Struct) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

// `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of that
// variants, absence of any variant indicates an error.
//
// The JSON representation for `Value` is JSON value.
type Value struct {
	// The kind of value.
	//
	// Types that are valid to be assigned to Kind:
	//	*Value_NullValue
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_StructValue
	//	*Value_ListValue
	Kind                 isValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_84428fcddb6debf8, []int{1}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Kind interface {
	isValue_Kind()
	Equal(interface{}) bool
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=mesg.protobuf.NullValue,oneof" json:"null_value,omitempty" hash:"name:1"`
}
type Value_NumberValue struct {
	NumberValue float64 `protobuf:"fixed64,2,opt,name=number_value,json=numberValue,proto3,oneof" json:"number_value,omitempty" hash:"name:2"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof" json:"string_value,omitempty" hash:"name:3"`
}
type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof" json:"bool_value,omitempty" hash:"name:4"`
}
type Value_StructValue struct {
	StructValue *Struct `protobuf:"bytes,5,opt,name=struct_value,json=structValue,proto3,oneof" json:"struct_value,omitempty" hash:"name:5"`
}
type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,6,opt,name=list_value,json=listValue,proto3,oneof" json:"list_value,omitempty" hash:"name:6"`
}

func (*Value_NullValue) isValue_Kind()   {}
func (*Value_NumberValue) isValue_Kind() {}
func (*Value_StringValue) isValue_Kind() {}
func (*Value_BoolValue) isValue_Kind()   {}
func (*Value_StructValue) isValue_Kind() {}
func (*Value_ListValue) isValue_Kind()   {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetNullValue() NullValue {
	if x, ok := m.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return NullValue_NULL_VALUE
}

func (m *Value) GetNumberValue() float64 {
	if x, ok := m.GetKind().(*Value_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetStructValue() *Struct {
	if x, ok := m.GetKind().(*Value_StructValue); ok {
		return x.StructValue
	}
	return nil
}

func (m *Value) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_NullValue)(nil),
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_ListValue)(nil),
	}
}

// `ListValue` is a wrapper around a repeated field of values.
//
// The JSON representation for `ListValue` is JSON array.
type ListValue struct {
	// Repeated field of dynamically typed values.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" hash:"name:1"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListValue) Reset()         { *m = ListValue{} }
func (m *ListValue) String() string { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()    {}
func (*ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_84428fcddb6debf8, []int{2}
}
func (m *ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListValue.Unmarshal(m, b)
}
func (m *ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListValue.Marshal(b, m, deterministic)
}
func (m *ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListValue.Merge(m, src)
}
func (m *ListValue) XXX_Size() int {
	return xxx_messageInfo_ListValue.Size(m)
}
func (m *ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListValue proto.InternalMessageInfo

func (m *ListValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterEnum("mesg.protobuf.NullValue", NullValue_name, NullValue_value)
	proto.RegisterType((*Struct)(nil), "mesg.protobuf.Struct")
	proto.RegisterMapType((map[string]*Value)(nil), "mesg.protobuf.Struct.FieldsEntry")
	proto.RegisterType((*Value)(nil), "mesg.protobuf.Value")
	proto.RegisterType((*ListValue)(nil), "mesg.protobuf.ListValue")
}

func init() { proto.RegisterFile("protobuf/types/struct.proto", fileDescriptor_84428fcddb6debf8) }

var fileDescriptor_84428fcddb6debf8 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe7, 0x66, 0x8d, 0xe8, 0x0b, 0x9b, 0xb6, 0x08, 0xa4, 0x6a, 0x43, 0x50, 0x72, 0xaa,
	0x26, 0x91, 0x40, 0x06, 0x13, 0xea, 0x05, 0x51, 0x69, 0x80, 0x44, 0x36, 0xaa, 0xc2, 0x86, 0xc4,
	0xa5, 0x6a, 0x52, 0x37, 0x8d, 0xe6, 0xd8, 0x53, 0x6c, 0x83, 0x7a, 0xe3, 0x93, 0x70, 0xe0, 0x02,
	0xe2, 0x43, 0x71, 0xe0, 0x23, 0x70, 0xe2, 0x88, 0x6c, 0x37, 0x81, 0x36, 0xb9, 0xd9, 0x7e, 0xff,
	0xdf, 0xdf, 0xff, 0xbc, 0xe7, 0xc0, 0xe1, 0x75, 0xc1, 0x04, 0x8b, 0xe5, 0x3c, 0x10, 0xcb, 0x6b,
	0xcc, 0x03, 0x2e, 0x0a, 0x99, 0x08, 0x5f, 0x9f, 0xba, 0x3b, 0x39, 0xe6, 0xa9, 0x5f, 0x2a, 0x0e,
	0xbc, 0x94, 0xa5, 0x2c, 0xa8, 0x00, 0xb5, 0xd3, 0x1b, 0xbd, 0x32, 0x32, 0xef, 0x1b, 0x02, 0xfb,
	0xad, 0xf6, 0x70, 0x23, 0xb0, 0xe7, 0x19, 0x26, 0x33, 0xde, 0x45, 0x3d, 0xab, 0xef, 0x84, 0xf7,
	0xfd, 0x35, 0x3b, 0xdf, 0xc8, 0xfc, 0x17, 0x5a, 0x73, 0x4a, 0x45, 0xb1, 0x1c, 0xee, 0xff, 0xfe,
	0x79, 0x6f, 0x67, 0x31, 0xe5, 0x8b, 0x81, 0x47, 0xa7, 0x39, 0x1e, 0x3c, 0xf2, 0xc6, 0x2b, 0x8f,
	0x83, 0x37, 0xe0, 0xfc, 0xa7, 0x74, 0xf7, 0xc0, 0xba, 0xc2, 0xcb, 0x2e, 0xea, 0xa1, 0x7e, 0x67,
	0xac, 0x96, 0xee, 0x11, 0xb4, 0x3f, 0x4e, 0x89, 0xc4, 0xdd, 0x56, 0x0f, 0xf5, 0x9d, 0xf0, 0xd6,
	0xc6, 0x6d, 0x97, 0xaa, 0x36, 0x36, 0x92, 0x41, 0xeb, 0x29, 0xf2, 0xbe, 0x58, 0xd0, 0xd6, 0x87,
	0x6e, 0x04, 0x40, 0x25, 0x21, 0x13, 0x83, 0x2b, 0xcb, 0xdd, 0xb0, 0xbb, 0x81, 0x9f, 0x4b, 0x42,
	0xb4, 0xba, 0x21, 0xe3, 0xab, 0xad, 0x71, 0x87, 0x96, 0x75, 0xf7, 0x04, 0x6e, 0x52, 0x99, 0xc7,
	0xb8, 0x98, 0xfc, 0x8b, 0x83, 0x6a, 0x54, 0xa8, 0x28, 0xc7, 0x08, 0x2b, 0x8e, 0x8b, 0x22, 0xa3,
	0xe9, 0x8a, 0xb3, 0xd4, 0xa7, 0xd5, 0xb8, 0x63, 0xcd, 0x19, 0xa1, 0xe1, 0x42, 0x80, 0x98, 0xb1,
	0x32, 0xfd, 0x76, 0x0f, 0xf5, 0x6f, 0xd4, 0xa8, 0xc7, 0x3a, 0xa3, 0x92, 0x19, 0xe6, 0x4c, 0xdf,
	0x25, 0x13, 0xb1, 0xa2, 0xda, 0xba, 0x65, 0xb7, 0x1b, 0x07, 0x54, 0x33, 0x7b, 0x52, 0x46, 0x90,
	0x89, 0xa8, 0x1a, 0x48, 0x32, 0x5e, 0x9a, 0xd9, 0xda, 0x6c, 0xb3, 0x81, 0x51, 0xc6, 0x45, 0x73,
	0x03, 0x4f, 0x74, 0x38, 0x52, 0xd5, 0x6d, 0xd8, 0xbe, 0xca, 0xe8, 0xcc, 0x8b, 0xa0, 0x53, 0x41,
	0xee, 0x33, 0xb0, 0xb5, 0x7b, 0xf9, 0x98, 0x1a, 0xc7, 0xdb, 0xf8, 0x7e, 0x0c, 0x76, 0x74, 0x08,
	0x9d, 0x6a, 0x86, 0xee, 0x2e, 0xc0, 0xf9, 0x45, 0x14, 0x4d, 0x2e, 0x9f, 0x47, 0x17, 0xa7, 0x7b,
	0x5b, 0xc3, 0xcf, 0xe8, 0xfb, 0xaf, 0xbb, 0x08, 0xf6, 0x13, 0x96, 0xaf, 0xfb, 0x0e, 0x1d, 0xd3,
	0x84, 0x91, 0xda, 0x8f, 0xd0, 0x87, 0x87, 0x69, 0x26, 0x16, 0x32, 0xf6, 0x13, 0x96, 0x07, 0x4a,
	0xf8, 0x60, 0xce, 0x24, 0x9d, 0x4d, 0x45, 0xc6, 0x68, 0x80, 0x69, 0x9a, 0x51, 0x1c, 0xac, 0xff,
	0x50, 0x7f, 0x10, 0xfa, 0xda, 0xb2, 0x5e, 0x8e, 0x86, 0x3f, 0x5a, 0x77, 0xce, 0x94, 0xf3, 0xa8,
	0x4c, 0xfc, 0x1e, 0x13, 0xf2, 0x9a, 0xb2, 0x4f, 0xf4, 0x9d, 0xd2, 0xc5, 0xb6, 0xe6, 0x8e, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xab, 0xa8, 0xd2, 0xf4, 0x91, 0x03, 0x00, 0x00,
}

func (this *Struct) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Struct)
	if !ok {
		that2, ok := that.(Struct)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Fields) != len(that1.Fields) {
		return false
	}
	for i := range this.Fields {
		if !this.Fields[i].Equal(that1.Fields[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Value) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Value)
	if !ok {
		that2, ok := that.(Value)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Kind == nil {
		if this.Kind != nil {
			return false
		}
	} else if this.Kind == nil {
		return false
	} else if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Value_NullValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Value_NullValue)
	if !ok {
		that2, ok := that.(Value_NullValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.NullValue != that1.NullValue {
		return false
	}
	return true
}
func (this *Value_NumberValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Value_NumberValue)
	if !ok {
		that2, ok := that.(Value_NumberValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.NumberValue != that1.NumberValue {
		return false
	}
	return true
}
func (this *Value_StringValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Value_StringValue)
	if !ok {
		that2, ok := that.(Value_StringValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.StringValue != that1.StringValue {
		return false
	}
	return true
}
func (this *Value_BoolValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Value_BoolValue)
	if !ok {
		that2, ok := that.(Value_BoolValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BoolValue != that1.BoolValue {
		return false
	}
	return true
}
func (this *Value_StructValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Value_StructValue)
	if !ok {
		that2, ok := that.(Value_StructValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.StructValue.Equal(that1.StructValue) {
		return false
	}
	return true
}
func (this *Value_ListValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Value_ListValue)
	if !ok {
		that2, ok := that.(Value_ListValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ListValue.Equal(that1.ListValue) {
		return false
	}
	return true
}
func (this *ListValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListValue)
	if !ok {
		that2, ok := that.(ListValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if !this.Values[i].Equal(that1.Values[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}