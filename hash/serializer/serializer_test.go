package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSerializable struct {
	a int
	b string
	c []string
}

func (data testSerializable) Serialize() string {
	ser := New()
	ser.AddInt("1", data.a)
	ser.AddString("2", data.b)
	ser.AddStringSlice("3", data.c)
	return ser.Serialize()
}

func TestSerializer(t *testing.T) {
	tests := []struct {
		f func(*Serializer)
		s string
	}{
		{func(ser *Serializer) { ser.AddBool("1", true) }, "1:true;"},
		{func(ser *Serializer) { ser.AddBool("2", false) }, ""},
		{func(ser *Serializer) { ser.AddFloat("3", 3.14159265359) }, "3:3.14159265359;"},
		{func(ser *Serializer) { ser.AddFloat("4", 0.0) }, ""},
		{func(ser *Serializer) { ser.AddFloat("5", -42.14159265359) }, "5:-42.14159265359;"},
		{func(ser *Serializer) { ser.AddInt("6", 42) }, "6:42;"},
		{func(ser *Serializer) { ser.AddInt("7", 0.0) }, ""},
		{func(ser *Serializer) { ser.AddInt("8", -42) }, "8:-42;"},
		{func(ser *Serializer) { ser.AddString("9", "") }, ""},
		{func(ser *Serializer) { ser.AddString("10", "hello") }, "10:hello;"},
		{func(ser *Serializer) { ser.AddStringSlice("11", []string{"c", "b", "a"}) }, "11:0:c;1:b;2:a;;"},
		{func(ser *Serializer) { ser.AddStringSlice("12", []string{}) }, ""},
		{func(ser *Serializer) { ser.AddStringSlice("13", []string{"c", "", "a"}) }, "13:0:c;2:a;;"},
		{func(ser *Serializer) { ser.Add("14", testSerializable{42, "hello", []string{"c", "b", "a"}}) }, "14:1:42;2:hello;3:0:c;1:b;2:a;;;"},
		{func(ser *Serializer) { ser.Add("15", testSerializable{}) }, ""},
	}
	// individual test
	for _, tt := range tests {
		ser := New()
		tt.f(ser)
		assert.Equal(t, tt.s, ser.Serialize())
	}
}