package orm

import (
	reflect2 "JoeyOrm/orm/reflect"
	"context"
	"strings"
)

type Selector[T any] struct {
	table string
	builder
	model *reflect2.Model
	db    *DB
}

func NewSelector[T any](db *DB) *Selector[T] {
	return &Selector[T]{db: db}
}

func (s *Selector[T]) Build() (*Query, error) {
	s.sb = &strings.Builder{}
	m, err := s.db.r.Get(new(T))
	if err != nil {
		return nil, err
	}
	s.model = m

	s.sb.WriteString("SELECT * FROM ")
	if s.table == "" {
		s.sb.WriteByte('`')
		s.sb.WriteString(m.TableName)
		s.sb.WriteByte('`')
	} else {
		spValue := strings.Split(s.table, ".")
		s.sb.WriteByte('`')
		s.sb.WriteString(spValue[0])
		s.sb.WriteByte('`')
		s.sb.WriteByte('.')
		s.sb.WriteByte('`')
		s.sb.WriteString(spValue[1])
		s.sb.WriteByte('`')

	}

	err = s.BuilderPredicates("WHERE", s.model)
	if err != nil {
		return nil, err
	}

	s.sb.WriteByte(';')
	return &Query{SQL: s.sb.String(), Args: s.args}, nil
}

func (s *Selector[T]) From(tal string) *Selector[T] {
	s.table = tal
	return s
}

func (s *Selector[T]) Get(cxt context.Context) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Selector[T]) GetMulti(cxt context.Context) (*[]T, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Selector[T]) Where(ps ...Predicate) *Selector[T] {
	s.where = ps
	return s
}
