package service

import (
	"testing"

	"github.com/mesg-foundation/engine/protobuf/types"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		name      string
		params    []*Service_Parameter
		data      []*types.Value
		expecterr bool
	}{
		{
			name: "no parameters and no data",
		},
		{
			name: "one optional parameter without data",
			params: []*Service_Parameter{
				{
					Optional: true,
				},
			},
		},
		{
			name: "simple types (string,number,boolean)",
			params: []*Service_Parameter{
				{
					Key:  "string",
					Type: "String",
				},
				{
					Key:  "number",
					Type: "Number",
				},
				{
					Key:  "boolean",
					Type: "Boolean",
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_StringValue{},
				},
				{
					Kind: &types.Value_NumberValue{},
				},
				{
					Kind: &types.Value_BoolValue{},
				},
			},
		},
		{
			name: "any type",
			params: []*Service_Parameter{
				{
					Key:  "key",
					Type: "Any",
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_ListValue{},
				},
			},
		},
		{
			name: "array type",
			params: []*Service_Parameter{
				{
					Key:      "key",
					Type:     "Number",
					Repeated: true,
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_ListValue{
						ListValue: &types.Value_List{
							Values: []*types.Value{
								{
									Kind: &types.Value_NumberValue{},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "required parameter without data",
			params: []*Service_Parameter{
				{
					Optional: false,
				},
			},
			expecterr: true,
		},
		{
			name: "invalid parameter type",
			params: []*Service_Parameter{
				{
					Key:  "key",
					Type: "-",
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_NullValue{},
				},
			},
			expecterr: true,
		},
		{
			name: "invalid string type",
			params: []*Service_Parameter{
				{
					Key:  "key",
					Type: "String",
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_NullValue{},
				},
			},
			expecterr: true,
		},
		{
			name: "invalid number type",
			params: []*Service_Parameter{
				{
					Key:  "key",
					Type: "Number",
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_NullValue{},
				},
			},
			expecterr: true,
		},
		{
			name: "invalid boolean type",
			params: []*Service_Parameter{
				{
					Key:  "key",
					Type: "Boolean",
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_NullValue{},
				},
			},
			expecterr: true,
		},
		{
			name: "invalid list type",
			params: []*Service_Parameter{
				{
					Key:      "key",
					Repeated: true,
					Type:     "String",
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_NullValue{},
				},
			},
			expecterr: true,
		},
		{
			name: "invalid object type",
			params: []*Service_Parameter{
				{
					Key:  "key",
					Type: "Object",
				},
			},
			data: []*types.Value{
				{
					Kind: &types.Value_NullValue{},
				},
			},
			expecterr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.expecterr {
				assert.Error(t, validateServiceParameters(tt.params, tt.data))
			} else {
				assert.NoError(t, validateServiceParameters(tt.params, tt.data))
			}
		})
	}
}
