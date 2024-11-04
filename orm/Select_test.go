package orm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelector_Build(t *testing.T) {
	testCases := []struct {
		name      string
		builder   QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name:    "no_from",
			builder: NewSelector[TestModel](),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model`;", Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "from_empty",
			builder: NewSelector[TestModel]().From(""),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model`;", Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "from_test_milt",
			builder: NewSelector[TestModel]().From("test_db.test_model"),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_db`.`test_model`;", Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "empty_where",
			builder: NewSelector[TestModel]().Where(),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model`;",
			},
		},
		{
			name:    "where",
			builder: NewSelector[TestModel]().Where(C("Age").Eq(18)),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE `age` = ?;",
				Args: []any{18},
			},
		},
		{
			name:    "not",
			builder: NewSelector[TestModel]().Where(Not(C("Age").Eq(18))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE  NOT (`age` = ?);",
				Args: []any{18},
			},
		},
		{
			name:    "And",
			builder: NewSelector[TestModel]().Where(C("Age").Eq(18).And(C("FirstName").Eq("erPang"))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE (`age` = ?) AND (`first_name` = ?);",
				Args: []any{18, "erPang"},
			},
		},
		{
			name:    "or",
			builder: NewSelector[TestModel]().Where(C("Age").Eq(18).Or(C("FirstName").Eq("erPang"))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE (`age` = ?) OR (`first_name` = ?);",
				Args: []any{18, "erPang"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query, err := tc.builder.Build()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, query)
		})
	}
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       uint8
	LastName  string
}
