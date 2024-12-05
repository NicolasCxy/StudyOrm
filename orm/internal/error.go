package internal

import (
	"errors"
	"fmt"
)

var (
	ErrorPointerOnly = errors.New("只支持指向结构体的一级指针，如*User{}")
)

func NewUnSupportedExpression(expr any) error {
	return fmt.Errorf("不支持的表达式，类型： %v", expr)
}

func NewNotFoundField(field string) error {
	return fmt.Errorf("sql生成失败，不包含字段 %v", field)
}

func NewErrInvalidTagContent(pair string) error {
	return fmt.Errorf("无效的tag： %s", pair)
}

func NewErrorParsUnSupported(expr any) error {
	return fmt.Errorf("json解析只支持string，[]byte类型，不支持： %v", expr)
}
