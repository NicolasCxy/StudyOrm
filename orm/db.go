package orm

import "JoeyOrm/orm/reflect"

type DBOption func(*DB)

type DB struct {
	r *reflect.Register
}

func NewDB(opts ...DBOption) (*DB, error) {
	db := &DB{r: reflect.NewRegister()}

	for _, opt := range opts {
		opt(db)
	}

	return db, nil
}
