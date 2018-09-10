package casting

import (
	"testing"

	"github.com/mesg-foundation/core/interface/grpc/core"
	"github.com/stretchr/testify/require"
)

func TestServiceCast(t *testing.T) {
	var tests = []struct {
		service   *core.Service
		data      map[string]string
		expected  map[string]interface{}
		expectErr bool
	}{
		{
			createTestServcieWithInputs(nil),
			map[string]string{},
			map[string]interface{}{},
			false,
		},
		{
			createTestServcieWithInputs(map[string]string{
				"a": "String",
				"b": "Number",
				"c": "Number",
				"d": "Boolean",
			}),
			map[string]string{
				"a": "_",
				"b": "1",
				"c": "1.1",
				"d": "true",
			},
			map[string]interface{}{
				"a": "_",
				"b": int64(1),
				"c": 1.1,
				"d": true,
			},
			false,
		},
		{
			createTestServcieWithInputs(map[string]string{"a": "String"}),
			map[string]string{"b": "_"},
			map[string]interface{}{},
			true,
		},
		{
			createTestServcieWithInputs(map[string]string{"a": "Number"}),
			map[string]string{"a": "_"},
			map[string]interface{}{},
			true,
		},
		{
			createTestServcieWithInputs(map[string]string{"a": "Boolean"}),
			map[string]string{"a": "_"},
			map[string]interface{}{},
			true,
		},
		{
			createTestServcieWithInputs(map[string]string{"a": "Object"}),
			map[string]string{"a": `{"b":1}`},
			map[string]interface{}{"a": map[string]interface{}{"b": float64(1)}},
			false,
		},
	}

	for _, tt := range tests {
		got, err := TaskInputs(tt.service, "test", tt.data)
		if tt.expectErr {
			require.NotNil(t, err)
		} else {
			require.Equal(t, len(tt.expected), len(got), "maps len are not equal")
			require.Equal(t, tt.expected, got, "maps are not equal")
		}
	}

	// test if non-existing key returns error
	_, err := TaskInputs(tests[0].service, "_", nil)
	require.NotNil(t, err)
}

// creates test service with given inputs name and type under "test" task key.
func createTestServcieWithInputs(inputs map[string]string) *core.Service {
	s := &core.Service{
		Tasks: map[string]*core.Task{
			"test": {
				Inputs: make(map[string]*core.Parameter),
			},
		},
	}

	for name, itype := range inputs {
		s.Tasks["test"].Inputs[name] = &core.Parameter{Type: itype}
	}
	return s
}
