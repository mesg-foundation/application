package execution

import "github.com/mesg-foundation/engine/hash/hashserializer"

// HashSerialize returns the hashserialized string of this type
func (data *Execution) HashSerialize() string {
	ser := hashserializer.New()
	ser.AddString("2", data.ParentHash.String())
	ser.AddString("3", data.EventHash.String())
	ser.AddString("5", data.InstanceHash.String())
	ser.AddString("6", data.TaskKey)
	ser.Add("7", data.Inputs)
	ser.AddStringSlice("10", data.Tags)
	ser.AddString("11", data.ProcessHash.String())
	ser.AddString("12", data.NodeKey)
	ser.AddString("13", data.ExecutorHash.String())
	ser.AddString("14", data.Price)
	return ser.HashSerialize()
}