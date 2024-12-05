package sql

import (
	"JoeyOrm/orm/internal"
	"database/sql/driver"
	"encoding/json"
)

type JsonColum[T any] struct {
	Val   T
	Valid bool
}

// Value go类型结构体 - json
func (j *JsonColum[T]) Value() (driver.Value, error) {
	if !j.Valid {
		return nil, nil
	}
	return json.Marshal(j.Val)
}

// Scan json - go类型结构体
func (j *JsonColum[T]) Scan(src any) error {

	var bs []byte
	switch data := src.(type) {
	case []byte:

		bs = data
	case string:
		bs = []byte(data)
	case nil:
		return nil
	default:
		return internal.NewErrorParsUnSupported(data)
	}

	err := json.Unmarshal(bs, &j.Val)
	if err != nil {
		return err
	}
	//设置可用
	j.Valid = true
	return nil
}
