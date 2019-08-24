package workflow

import (
	"encoding/json"
	"fmt"

	"github.com/mesg-foundation/engine/hash"
)

// MarshalJSON for the task
func (t Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":         "task",
		"key":          t.Key,
		"instanceHash": t.InstanceHash.String(),
		"taskKey":      t.TaskKey,
	})
}

// MarshalJSON for the task
func (r Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":         "result",
		"instanceHash": r.InstanceHash.String(),
		"taskKey":      r.TaskKey,
	})
}

// MarshalJSON for the task
func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":         "event",
		"instanceHash": e.InstanceHash.String(),
		"eventKey":     e.EventKey,
	})
}

// MarshalJSON for the task
func (m Mapping) MarshalJSON() ([]byte, error) {
	// outputs, err := json.Marshal(m.Outputs)
	// if err != nil {
	// 	return nil, err
	// }
	return json.Marshal(map[string]interface{}{
		"type":    "mapping",
		"key":     m.Key,
		"outputs": m.Outputs,
	})
}

// UnmarshalJSON unmashals a workflow
func (w *Workflow) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*objMap["Hash"], &w.Hash)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*objMap["Key"], &w.Key)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*objMap["Edges"], &w.Edges)
	if err != nil {
		return err
	}

	var rawNodes []*json.RawMessage
	err = json.Unmarshal(*objMap["Nodes"], &rawNodes)
	if err != nil {
		return err
	}

	w.Graph.Nodes = make([]Node, len(rawNodes))
	for i, rawNode := range rawNodes {
		var nodeInfo map[string]interface{}
		err = json.Unmarshal(*rawNode, &nodeInfo)
		if err != nil {
			return err
		}
		nodeType := nodeInfo["type"]
		data := make(map[string]interface{})
		for key, value := range nodeInfo {
			if key == "type" {
				continue
			}
			if key == "instanceHash" {
				instanceHash, err := hash.Decode(value.(string))
				if err != nil {
					return err
				}
				data[key] = instanceHash
			} else {
				data[key] = value
			}
		}
		marshalData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		switch nodeType {
		case "task":
			var node Task
			if err := json.Unmarshal(marshalData, &node); err != nil {
				return err
			}
			w.Graph.Nodes[i] = &node
		case "event":
			var node Event
			if err := json.Unmarshal(marshalData, &node); err != nil {
				return err
			}
			w.Graph.Nodes[i] = &node
		case "result":
			var node Result
			if err := json.Unmarshal(marshalData, &node); err != nil {
				return err
			}
			w.Graph.Nodes[i] = &node
		case "mapping":
			var node Mapping
			if err := json.Unmarshal(marshalData, &node); err != nil {
				return err
			}
			w.Graph.Nodes[i] = &node
		default:
			return fmt.Errorf("type %q not supported", nodeType)
		}
	}
	return nil
}
