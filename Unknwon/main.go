// test project main.go
package main

import (
	"fmt"
	//"sort"
	//"gotest"
	"reflect"
	//"time"
	//"runtime"
	"sync"
)

type humen struct {
	Sex int
}

type person struct {
	humen
	women
	Name string
	Age  int
	//Contact struct {
	//Phone, City string
	//}
}

type women struct {
	Sex int
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello World")
}

func Pingpong(s []int) []int {
	s = append(s, 3)
	return s
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

func main() {
	//var s [2]string
	//a := [...]int{19: 1}
	//aa := [2]int{1, 2}
	//bb := [2]int{11, 2}
	//p := new([10]int)
	//slice := make([]int, 3, 10)
	//slice = append(slice, 1, 2, 3)
	//fmt.Printf("%v %p \n", slice, slice)
	//slice = append(slice, 4, 5, 6)
	//fmt.Printf("%v %p \n", slice, slice)
	//slice = append(slice, 7, 8, 9)
	//fmt.Printf("%v %p \n", slice, slice)
	//slicetesta := []int{1, 2, 3, 4, 5, 6}
	//slicetetsb := []int{7, 8, 9}
	//copy(slicetesta, slicetetsb)
	//fmt.Println(slicetesta, slicetetsb)
	//slicea := slice[0:9]
	//sliceb := slice[11:]
	//fmt.Println(s)
	//fmt.Println(a)
	//fmt.Println(aa == bb)
	//fmt.Println(p)
	//fmt.Println(len(slicea), cap(slicea))
	//fmt.Println(len(sliceb), cap(sliceb))
	//m := make(map[int]string)
	//m[1] = "ok"
	//m[2] = "xxx"
	//delete(m, 1)
	//am := m[1]
	//bm := m[2]
	//fmt.Println(am, bm)
	//mm := make(map[int]map[int]string)
	//mm[1] = make(map[int]string)
	//mm[1][1] = "yyy"
	//fmt.Println(mm)
	//m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	//s := make([]int, len(m))
	//i := 0
	//for k, _ := range m {
	//s[i] = k
	//i++
	//}
	//sort.Ints(s)
	//fmt.Println(s)
	//A(1, 2, 3, 4, 5)
	//fun := closure(10)
	//fun(10)
	//fmt.Println(fun(10))
	//A()
	//B()
	//C()
	//a := person{Name: "name", Age: 19, humen: humen{Sex: 99}, women: women{Sex: 109}}
	//a.Contact.City = "shenzh"
	//a.Contact.Phone = "132333356"
	//fmt.Println(a.humen.Sex, a.women.Sex)
	//f := person{}
	//(*person).Print(&a)
	//f.Print()
	//u := User{1, "OK", 12}
	//Info(u)
	//x := 123
	//v := reflect.ValueOf(&x)
	//v.Elem().SetInt(999)
	//fmt.Println(x)
	//u := User{1, "OK", 12}
	//Set(&u)
	//fmt.Println(u)
	//
	//v := reflect.ValueOf(u)
	//mv := v.MethodByName("Hello")
	//args := []reflect.Value{reflect.ValueOf("joe")}
	//mv.Call(args)
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//wg.Add(10)
	//for i := 0; i < 10; i++ {
	//go Go(&wg)
	//}
	//wg.Wait()
	//gotest.Hello()
	//s := make([]int, 0)
	//fmt.Println(s)
	//s = Pingpong(s)
	//fmt.Println(s)
	//t := time.Now()
	//fmt.Println(t.Format(time.ANSIC))
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func( /*v string*/ ) {
			fmt.Println(v)
		}( /*v*/ )
	}

	select {}

}

func Go(wg *sync.WaitGroup) {
	fmt.Printf("Go Go Go!")
	wg.Done()
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

//Reciver
//func (a *person) Print(): {
//fmt.Println("person")
//}

//func closure(a int) func(int) int {
//return func(b int) int {
//return a * b
//}
//}

//func A(a ...int) {
//fmt.Println(a)
//}

//func A() {
//	fmt.Println("func A")
//}

//func B() {
//	defer func() {
//		if error := recover(); error != nil {
//			fmt.Println("Recoved in B")
//		}
//	}()
//	panic("func b")
//	fmt.Println("func B")
//}

//func C() {
//	fmt.Println("func C")
//}
