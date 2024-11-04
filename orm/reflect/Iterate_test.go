package reflect

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_Iterate_ArrOrSlice(t *testing.T) {
	testCase := []struct {
		name      string
		entity    any
		wantValue []any
		wantError error
	}{
		{
			name:      "arrTest",
			entity:    [3]any{1, 2, 3},
			wantValue: []any{1, 2, 3},
		},
		{
			name:      "sliceTest",
			entity:    []any{1, 2, 3},
			wantValue: []any{1, 2, 3},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			value, err := IterateArrOrSlice(tc.entity)
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantValue, value)
		})
	}
}

func Test_Iterate_Map(t *testing.T) {
	testCase := []struct {
		name      string
		entity    any
		wantVals  []any
		wantKeys  []any
		wantError error
	}{
		{
			name: "mapTest",
			entity: map[string]string{
				"A": "a",
				"B": "b",
				"C": "c",
			},
			wantKeys: []any{"A", "B", "C"},
			wantVals: []any{"a", "b", "c"},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			keys, values, err := IterateMap(tc.entity)
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}
			sortArr(keys)
			sortArr(values)
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantVals, values)
		})
	}
}

func sortArr(arr []any) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].(string) < arr[j].(string)
	})
}
