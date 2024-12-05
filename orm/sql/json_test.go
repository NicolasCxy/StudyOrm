package sql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson_Value(t *testing.T) {
	js := JsonColum[User]{Valid: true, Val: User{
		Name: "cxy",
		Age:  22,
	}}

	value, err := js.Value()
	assert.Nil(t, err)
	assert.Equal(t, `{"Name":"cxy","Age":22}`, string(value.([]byte)))

	js = JsonColum[User]{}
	value, err = js.Value()
	if err != nil {
		return
	}
	assert.Nil(t, err)
	assert.Nil(t, value)

}

func TestJson_Scan(t *testing.T) {
	js := JsonColum[User]{}

	err := js.Scan(`{"Name":"cxy6666","Age":33}`)
	assert.Nil(t, err)

	value, err := js.Value()
	assert.Nil(t, err)
	fmt.Println(string(value.([]byte)))
	//assert.Equal(t, `{"Name":"cxy","Age":22}`, string(value.([]byte)))
}

type User struct {
	Name string
	Age  int
}
