package reflect

import (
	"JoeyOrm/orm/entity"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMethod(t *testing.T) {

	testCase := []struct {
		name       string
		entity     any
		inputValue []any
		wantRes    map[string]FuncInfo
		wantErr    error
	}{
		{
			name:   "GetName",
			entity: entity.NewUser("cxy"),
			wantRes: map[string]FuncInfo{
				"GetName": {
					Name:      "GetName",
					InputType: []reflect.Type{reflect.TypeOf(entity.User{})},
					OutputType: []reflect.Type{
						reflect.TypeOf("string"),
					},
					Result: []any{"cxy"},
				},
			},
		},
		{
			name:       "GetCustomName",
			entity:     entity.NewUser("cxy"),
			inputValue: []any{"erPang"},
			wantRes: map[string]FuncInfo{
				"GetCustomName": {
					Name:       "GetCustomName",
					InputType:  []reflect.Type{reflect.TypeOf(entity.User{}), reflect.TypeOf("str")},
					OutputType: []reflect.Type{reflect.TypeOf("string")},
					Result:     []any{"erPang"},
				},
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result, err := IterateFunc(tc.entity, tc.name, tc.inputValue)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, result)
		})
	}
}
