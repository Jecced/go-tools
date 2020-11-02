package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name string `user:"haha"`
}

type Face interface {
}

func TestName(t *testing.T) {
	u := User{}
	ut := reflect.TypeOf(u)
	uv := reflect.ValueOf(u)
	fmt.Printf("reflect.TypeOf\t %T\n", ut)
	fmt.Printf("reflect.ValueOf\t %T\n", uv)

	field := ut.Field(0)

	fmt.Println("\ntag:")
	fmt.Println(field.Tag.Get("user"))

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))

	//o.reflectType.Elem().Field(i).Tag
}
