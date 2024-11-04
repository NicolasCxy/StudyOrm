package reflect

import (
	"JoeyOrm/orm/internal"
	"reflect"
	"unicode"
)

type Register struct {
	models map[reflect.Type]*Model
}

func NewRegister() *Register {
	return &Register{models: make(map[reflect.Type]*Model, 32)}
}

type Model struct {
	TableName string
	FieldMap  map[string]*Field
}

type Field struct {
	ColName string
}

func (r *Register) Get(val any) (*Model, error) {
	typ := reflect.TypeOf(val)
	m, ok := r.models[typ]
	if !ok {
		var err error
		m, err = r.ParseModel(val)
		if err != nil {
			return nil, err
		}
		r.models[typ] = m
	}
	return m, nil
}

func (r *Register) ParseModel(entity any) (*Model, error) {
	typ := reflect.TypeOf(entity)

	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return nil, internal.NewUnSupportedExpression(typ)
	}

	typ = typ.Elem()

	numFields := typ.NumField()
	fds := make(map[string]*Field, numFields)
	for i := 0; i < numFields; i++ {
		fdType := typ.Field(i)
		fds[fdType.Name] = &Field{ColName: underscoreName(fdType.Name)}
	}

	return &Model{TableName: underscoreName(typ.Name()), FieldMap: fds}, nil
}

func underscoreName(tableName string) string {
	var buf []byte
	for i, v := range tableName {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(unicode.ToLower(v)))
		}
	}
	return string(buf)
}
