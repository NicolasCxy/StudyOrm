package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func IterateFields(entity any) (map[string]any, error) {
	if entity == nil {
		return nil, errors.New("entity is nil")
	}

	typ := reflect.TypeOf(entity)
	val := reflect.ValueOf(entity)

	if val.IsZero() {
		return nil, errors.New("访问对象或字段空间为零")
	}

	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("不支持类型")
	}

	numFields := typ.NumField()
	typ.NumMethod()

	m := make(map[string]any, numFields)
	for i := 0; i < numFields; i++ {
		fd := typ.Field(i)
		faVal := val.Field(i)
		if fd.IsExported() {
			m[fd.Name] = faVal.Interface()
		} else {
			m[fd.Name] = reflect.Zero(fd.Type).Interface()
		}
	}

	return m, nil
}

func SetField(entity any, field string, newValue any) error {
	if entity == nil {
		return errors.New("entity is nil")
	}

	val := reflect.ValueOf(entity)
	for val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}

	fieldValue := val.FieldByName(field)
	if !fieldValue.CanSet() {
		return fmt.Errorf("不可修改字段")
	}

	fieldValue.Set(reflect.ValueOf(newValue))

	return nil
}
