package orm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleter(t *testing.T) {
	testCase := []struct {
		name      string
		builder   QueryBuilder
		WantQuery *Query
		WantErr   error
	}{
		{
			name:      "delete",
			builder:   NewDeleter[TestModel](),
			WantQuery: &Query{SQL: "DELETE FROM `test_model`;", Args: nil},
		},
		{
			name:      "delete_where",
			builder:   NewDeleter[TestModel]().Where(C("Age").Eq(18)),
			WantQuery: &Query{SQL: "DELETE FROM `test_model` WHERE `age` = ?;", Args: []any{18}},
		},
		{
			name:      "delete_table",
			builder:   NewDeleter[TestModel]().From("test_model").Where(C("FirstName").Eq("cxy")),
			WantQuery: &Query{SQL: "DELETE FROM `test_model` WHERE `first_name` = ?;", Args: []any{"cxy"}},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			query, err := tc.builder.Build()
			assert.Equal(t, tc.WantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.WantQuery, query)
		})
	}
}
