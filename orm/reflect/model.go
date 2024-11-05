package reflect

import (
	"JoeyOrm/orm/internal"
	"reflect"
	"strings"
	"sync"
	"unicode"
)

var (
	tagKeyColumn = "column"
)

type ModelOpt func(model *Model) error

func ModelWithTableName(tableName string) ModelOpt {
	return func(model *Model) error {
		model.TableName = tableName
		return nil
	}
}

type Registry interface {
	Get(val any) (*Model, error)
	Register(val any, opts ...ModelOpt) (*Model, error)
}

type Register struct {
	models sync.Map
}

func NewRegister() *Register {
	return &Register{}
}

type Model struct {
	TableName string
	FieldMap  map[string]*Field
}

type Field struct {
	ColName string
}

// Get 从缓存读取，读不到就创建
func (r *Register) Get(val any) (*Model, error) {
	typ := reflect.TypeOf(val)
	m, ok := r.models.Load(typ)
	if ok {
		return m.(*Model), nil
	}
	//不包含就创建
	return r.Register(val)
}

// Register 创建model
func (r *Register) Register(val any, opts ...ModelOpt) (*Model, error) {
	m, err := r.ParseModel(val)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err = opt(m)
		if err != nil {
			return nil, err
		}
	}

	typ := reflect.TypeOf(val)
	r.models.Store(typ, m)
	return m, nil
}

//func (r *Register) Get1(val any) (*Model, error) {
//	typ := reflect.TypeOf(val)
//	r.lock.RLock()
//	m, ok := r.models[typ]
//	r.lock.RUnlock()
//	if ok {
//		return m, nil
//	}
//
//	r.lock.Lock()
//	defer r.lock.Unlock()
//
//	m, ok = r.models[typ]
//	if !ok {
//		var err error
//		m, err = r.ParseModel(val)
//		if err != nil {
//			return nil, err
//		}
//		r.models[typ] = m
//	}
//
//	return m, nil
//}

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
		tags, err := r.ParseTag(fdType.Tag)
		if err != nil {
			return nil, err
		}

		//支持自定义列名
		columnName := tags[tagKeyColumn]
		if columnName == "" {
			columnName = underscoreName(fdType.Name)
		}

		fds[fdType.Name] = &Field{ColName: columnName}
	}

	//支持自定义表名
	var tableName string
	if tn, ok := entity.(TableName); ok {
		tableName = tn.CustomName()
	}

	//如果没有设置自定义就用默认
	if tableName == "" {
		tableName = underscoreName(typ.Name())
	}

	return &Model{TableName: underscoreName(tableName), FieldMap: fds}, nil
}

func (r *Register) ParseTag(tag reflect.StructTag) (map[string]string, error) {
	ormTag := tag.Get("orm")
	if ormTag == "" {
		return map[string]string{}, nil
	}
	//orm:"column:cxy"

	//拆解字符串
	pairs := strings.Split(ormTag, ",")
	res := make(map[string]string, len(pairs))
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, internal.NewErrInvalidTagContent(pair)
		}
		res[kv[0]] = kv[1]
	}
	return res, nil
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

type TableName interface {
	CustomName() string
}
