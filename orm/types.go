package orm

import (
	"context"
	"database/sql"
)

// Querier 查询接口
type Querier[T any] interface {
	Get(cxt context.Context) (*T, error)
	GetMulti(cxt context.Context) (*[]T, error)
}

// Deleter 查询接口
type Deleter[T any] interface {
	Delete(cxt context.Context) (*T, error)
	DeleteMulti(cxt context.Context) (*[]T, error)
}

// Executor 增删改接口
type Executor interface {
	Exec(ctx context.Context) (sql.Result, error)
}

type Query struct {
	SQL  string
	Args []any
}

type QueryBuilder interface {
	Build() (*Query, error)
}
