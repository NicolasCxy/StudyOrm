package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReflect(t *testing.T) {

	type User struct {
		Name string
		age  int
	}

	testCase := []struct {
		name    string
		entity  any
		wantErr error
		wantRes map[string]any
	}{
		{
			name:   "struct",
			entity: User{"cxy", 18},
			wantRes: map[string]any{
				"Name": "cxy",
				"age":  0,
			},
		},
		{
			name:   "pointer",
			entity: User{"cxy", 18},
			wantRes: map[string]any{
				"Name": "cxy",
				"age":  0,
			},
		},
		{
			name: "multiple pointer",
			entity: func() **User {
				res := &User{
					Name: "cxy",
					age:  0,
				}
				return &res
			}(),
			wantRes: map[string]any{
				"Name": "cxy",
				"age":  0,
			},
		}, {
			name:    "basicType",
			entity:  "152215",
			wantErr: errors.New("不支持类型"),
		},
		{
			name:    "nil",
			entity:  nil,
			wantErr: errors.New("entity is nil"),
		},
		{
			name:    "user_nil",
			entity:  (*User)(nil),
			wantErr: errors.New("访问对象或字段空间为零"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFields(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}

}

func TestSetField(t *testing.T) {

	type User struct {
		Name string
		age  int
	}

	testCase := []struct {
		name       string
		entity     any
		field      string
		newValue   any
		wantErr    error
		wantEntity any
	}{
		{
			name:       "struct",
			entity:     &User{"cxy", 18},
			field:      "Name",
			newValue:   "LaoSi",
			wantEntity: &User{"LaoSi", 18},
		},
		{
			name:       "pointer exported",
			entity:     &User{"cxy", 18},
			field:      "age",
			newValue:   19,
			wantEntity: &User{"cxy", 19},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			err := SetField(tc.entity, tc.field, tc.newValue)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantEntity, tc.entity)
		})
	}

	//var i = "cxy"
	//ptr := &i
	//reflect.ValueOf(ptr).Elem().Set(reflect.ValueOf("joey"))
	//assert.Equal(t, "joey", i)
}
