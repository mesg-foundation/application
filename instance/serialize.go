package instance

import "github.com/mesg-foundation/engine/hash/hashserializer"

// HashSerialize returns the hashserialized string of this type
func (data *Instance) HashSerialize() string {
	if data == nil {
		return ""
	}
	ser := hashserializer.New()
	ser.AddString("2", data.ServiceHash.String())
	ser.AddString("3", data.EnvHash.String())
	return ser.HashSerialize()
}