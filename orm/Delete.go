package orm

import (
	"JoeyOrm/orm/reflect"
	"context"
	"strings"
)

type Delete[T any] struct {
	table string
	builder
	model *reflect.Model
	db    *DB
}

func NewDeleter[T any]() *Delete[T] {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	return &Delete[T]{db: db}
}

func (d *Delete[T]) Delete(cxt context.Context) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Delete[T]) DeleteMulti(cxt context.Context) (*[]T, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Delete[T]) Where(ps ...Predicate) *Delete[T] {
	d.where = ps
	return d
}

func (d *Delete[T]) From(table string) *Delete[T] {
	d.table = table
	return d
}

func (d Delete[T]) Build() (*Query, error) {
	d.sb = &strings.Builder{}
	var err error
	d.model, err = d.db.r.Get(new(T))
	if err != nil {
		return nil, err
	}
	d.sb.WriteString("DELETE FROM ")
	//parse
	if d.table == "" {
		d.sb.WriteByte('`')
		d.sb.WriteString(d.model.TableName)
		d.sb.WriteByte('`')
	} else {
		d.sb.WriteByte('`')
		d.sb.WriteString(d.table)
		d.sb.WriteByte('`')
	}

	err = d.BuilderPredicates("WHERE", d.model)
	if err != nil {
		return nil, err
	}

	d.sb.WriteByte(';')

	return &Query{SQL: d.sb.String(), Args: d.args}, nil
}
