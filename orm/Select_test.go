package orm

import (
	"JoeyOrm/orm/internal"
	"JoeyOrm/orm/reflect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSelector_Build(t *testing.T) {

	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	testCases := []struct {
		name      string
		builder   QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name:    "no_from",
			builder: NewSelector[TestModel](db),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model`;", Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "from_empty",
			builder: NewSelector[TestModel](db).From(""),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model`;", Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "from_test_milt",
			builder: NewSelector[TestModel](db).From("test_db.test_model"),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_db`.`test_model`;", Args: nil,
			},
			wantErr: nil,
		},
		{
			name:    "empty_where",
			builder: NewSelector[TestModel](db).Where(),
			wantQuery: &Query{
				SQL: "SELECT * FROM `test_model`;",
			},
		},
		{
			name:    "where",
			builder: NewSelector[TestModel](db).Where(C("Age").Eq(18)),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE `age` = ?;",
				Args: []any{18},
			},
		},
		{
			name:    "not",
			builder: NewSelector[TestModel](db).Where(Not(C("Age").Eq(18))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE  NOT (`age` = ?);",
				Args: []any{18},
			},
		},
		{
			name:    "And",
			builder: NewSelector[TestModel](db).Where(C("Age").Eq(18).And(C("FirstName").Eq("erPang"))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE (`age` = ?) AND (`first_name` = ?);",
				Args: []any{18, "erPang"},
			},
		},
		{
			name:    "or",
			builder: NewSelector[TestModel](db).Where(C("Age").Eq(18).Or(C("FirstName").Eq("erPang"))),
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

func TestTags_Build(t *testing.T) {
	testCases := []struct {
		name      string
		val       func() any
		wantModel *reflect.Model
		wantError error
	}{
		{
			name: "empty_column",
			val: func() any {
				type EmptyColumn struct {
					FirstName string `orm:"colum="`
				}
				return &EmptyColumn{}
			},
			wantModel: &reflect.Model{
				TableName: "empty_column",
				FieldMap: map[string]*reflect.Field{
					"FirstName": {
						ColName: "first_name",
					},
				},
			},
		},
		{
			name: "no_set_column",
			val: func() any {
				type EmptyColumn struct {
					FirstName string
				}
				return &EmptyColumn{}
			},
			wantModel: &reflect.Model{
				TableName: "empty_column",
				FieldMap: map[string]*reflect.Field{
					"FirstName": {
						ColName: "first_name",
					},
				},
			},
		}, {
			name: "invalid_column",
			val: func() any {
				type EmptyColumn struct {
					FirstName string `orm:"column"`
				}
				return &EmptyColumn{}
			},

			wantError: internal.NewErrInvalidTagContent("column"),
		},
		{
			name: "invalid_column",
			val: func() any {
				type EmptyColumn struct {
					FirstName string `orm:"column"`
				}
				return &EmptyColumn{}
			},

			wantError: internal.NewErrInvalidTagContent("column"),
		},
		{
			name: "ignore_column",
			val: func() any {
				type EmptyColumn struct {
					FirstName string `orm:"adb=ccc"`
				}
				return &EmptyColumn{}
			},
			wantModel: &reflect.Model{
				TableName: "empty_column",
				FieldMap: map[string]*reflect.Field{
					"FirstName": {
						ColName: "first_name",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			register := reflect.NewRegister()
			m, err := register.Get(tc.val())
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantModel, m)
		})
	}
}

func TestTable_Build(t *testing.T) {
	testCases := []struct {
		name      string
		val       any
		wantModel *reflect.Model
		wantError error
	}{
		{
			name: "table_name",
			val:  &CustomTableName{},
			wantModel: &reflect.Model{
				TableName: "custom_table_name",
				FieldMap: map[string]*reflect.Field{
					"FirstName": {
						ColName: "first_name",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			register := reflect.NewRegister()
			m, err := register.Get(tc.val)
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantModel, m)
		})
	}
}

func TestRegistry_Build(t *testing.T) {
	testCases := []struct {
		name      string
		val       any
		option    reflect.ModelOpt
		wantModel *reflect.Model
		wantError error
	}{ //自定义表名和列名
		{
			name:   "table_name_code",
			val:    &CustomTableName{},
			option: reflect.ModelWithTableName("test_table_name_cc"),
			wantModel: &reflect.Model{
				TableName: "test_table_name_cc",
				FieldMap: map[string]*reflect.Field{
					"FirstName": {
						ColName: "fi_name",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			register := reflect.NewRegister()
			m, err := register.Register(tc.val, tc.option)
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantModel, m)
		})
	}
}

func TestCustomField_Build(t *testing.T) {
	testCases := []struct {
		name        string
		val         any
		option      reflect.ModelOpt
		field       string
		wantColName string
		wantError   error
	}{ //自定义表名和列名
		{
			name:        "table_name_code",
			val:         &CustomTableName{},
			option:      reflect.ModelWithColumnName("FirstName", ""),
			field:       "FirstName",
			wantColName: "",
		},
		{
			name:      "invalid_field",
			val:       &CustomTableName{},
			option:    reflect.ModelWithColumnName("FirstNameXXX", ""),
			wantError: internal.NewNotFoundField("FirstNameXXX"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			register := reflect.NewRegister()
			m, err := register.Register(tc.val, tc.option)
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}

			fd, ok := m.FieldMap[tc.field]
			require.True(t, ok)
			assert.Equal(t, tc.wantColName, fd.ColName)
		})
	}
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       uint8
	LastName  string
}

type CustomTableName struct {
	FirstName string `orm:"column=fi_name"`
}

var _ reflect.TableName = &CustomTableName{}

func (c CustomTableName) CustomName() string {
	return ""
}
