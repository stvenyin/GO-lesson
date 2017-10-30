package main

import (
	"fmt"
	//"sort"
	"reflect"
)"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello World")
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}

func (u User) Hello() {
	fmt.Println("Hello World")
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	} else {
		if f.Kind() == reflect.String {
			f.SetString("BYEBYE")
		}
	}

}

func main() {

	fmt.Println("xxx")
	u := User{1, "OK", 12}
	Info(u)

	//modifidy
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	Set(&u)
	fmt.Println(u)

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("joe")}

	mv.Call(args)
}