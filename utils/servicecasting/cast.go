package casting

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/mesg-foundation/core/protobuf/coreapi"
)

type caster func(value string) (interface{}, error)

func castString(value string) (interface{}, error) {
	return value, nil
}

func castNumber(value string) (interface{}, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		return i, nil
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("input %q is not a Number type", value)
	}
	return f, nil
}

func castBoolean(value string) (interface{}, error) {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return nil, fmt.Errorf("input %q is not a Boolean type", value)
	}
	return b, nil
}

func castObject(value string) (interface{}, error) {
	var v interface{}
	if err := json.Unmarshal([]byte(value), &v); err != nil {
		return nil, fmt.Errorf("input %q is not a Object type", value)
	}
	return v, nil
}

func castAny(value string) (interface{}, error) {
	return value, nil
}

var casters = map[string]caster{
	"String":  castString,
	"Number":  castNumber,
	"Boolean": castBoolean,
	"Object":  castObject,
	"Any":     castAny,
}

// TaskInputs converts map[string]string to map[string]interface{} based on defined types in the service tasks map.
func TaskInputs(s *coreapi.Service, taskKey string, taskData map[string]string) (map[string]interface{}, error) {
	var task *coreapi.Task
	for i := range s.Tasks {
		if s.Tasks[i].Key == taskKey {
			task = s.Tasks[i]
			break
		}
	}
	if task == nil {
		return nil, fmt.Errorf("service %q: task %q not found", s.Name, taskKey)
	}

	m := make(map[string]interface{}, len(taskData))
	for key, value := range taskData {
		var param *coreapi.Parameter
		for i := range task.Inputs {
			if task.Inputs[i].Key == key {
				param = task.Inputs[i]
				break
			}
		}
		if param == nil {
			return nil, fmt.Errorf("service %q: param %q not found", s.Name, key)
		}

		newValue, err := taskInputs(value, param.Type, param.Repeated)
		if err != nil {
			return nil, fmt.Errorf("task %q - %s", taskKey, err)
		}
		if newValue != nil {
			m[key] = newValue
		}
	}
	return m, nil
}

// taskInputs converts single input based on its type.
func taskInputs(input, inputType string, repeated bool) (interface{}, error) {
	c, ok := casters[inputType]
	if !ok {
		return nil, fmt.Errorf("input %q type %s - invalid type", input, inputType)
	}

	if repeated {
		var ret []interface{}
		for _, value := range strings.Split(input, ",") {
			v, err := c(value)
			if err != nil {
				return nil, fmt.Errorf("input %q - %s is not valid type of %s", input, value, inputType)
			}
			ret = append(ret, v)
		}
		return ret, nil
	}
	return c(input)
}
