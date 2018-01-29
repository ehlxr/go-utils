package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age" default:"18"`
	addr string `json:"addr"`
}

func (u User) Do(in string) (string, int) {
	fmt.Printf("%s Name is %s, Age is %d \n", in, u.Name, u.Age)
	return u.Name, u.Age
}

func main() {
	u := User{"tom", 27, "beijing"}

	// 获取对象的 Value
	v := reflect.ValueOf(u)
	fmt.Println("Value:", v)
	// fmt.Printf("%v\n", u)

	// 获取对象的 Type
	t := reflect.TypeOf(u)
	fmt.Println("Type:", t)
	// fmt.Printf("%T\n", u)

	t1 := v.Type()
	fmt.Println(t == t1)

	v1 := reflect.New(t)
	fmt.Println(v1)
	fmt.Println()

	// 获取 Kind 类型
	k := t.Kind()
	fmt.Println("Kind:", k)
	k1 := v.Kind()
	fmt.Println(k == k1)
	fmt.Println()

	// 修改反射对象的值
	i := 20
	fmt.Println("before i =", i)
	e := reflect.Indirect(reflect.ValueOf(&i))
	// e := reflect.ValueOf(&i).Elem()
	if e.CanSet() {
		e.SetInt(22)
	}
	fmt.Println("after i =", i)
	fmt.Println()

	// 反射字段操作
	// elem := reflect.Indirect(reflect.ValueOf(&u))
	elem := reflect.ValueOf(&u).Elem()
	for i := 0; i < t.NumField(); i++ {
		// 反射获取字段的元信息，例如：名称、Tag 等
		ft := t.Field(i)
		fmt.Println("field name:", ft.Name)
		tag := ft.Tag
		fmt.Println("Tag:", tag)
		fmt.Println("Tag json:", tag.Get("json"))

		// 反射修改字段的值
		fv := elem.Field(i)
		if fv.CanSet() {
			if fv.Kind() == reflect.Int {
				fmt.Println("change age to 30")
				fv.SetInt(30)
			}
			if fv.Kind() == reflect.String {
				fmt.Println("change name to jerry")
				fv.SetString("jerry")
			}
		}
		fmt.Println()
	}
	fmt.Println("after user:", u)
	fmt.Println()

	// 反射方法操作
	for i := 0; i < v.NumMethod(); i++ {
		method := t.Method(i) // 获取方法信息对象，方法 1
		mt := method.Type     // 获取方法信息 Type 对象，方法 1

		// m := v.Method(i) // 获取方法信息对象，方法 2
		// mt := m.Type()   // 获取方法信息 Type 对象，方法 2

		fmt.Println("method name:", method.Name)

		in := []reflect.Value{}

		// 获取方法入参类型
		for j := 0; j < mt.NumIn(); j++ {
			fmt.Println("method in type:", mt.In(j))
			if mt.In(j).Kind() == reflect.String {
				in = append(in, reflect.ValueOf("welcome"))
			}
			// 方法 1 获取的方法信息对象会把方法的接受者也当着入参之一
			if mt.In(j).Name() == t.Name() {
				in = append(in, v)
			}
		}

		// 获取方法返回类型
		for j := 0; j < mt.NumOut(); j++ {
			fmt.Println("method out type:", mt.Out(j))
		}

		// 反射方法调用
		// out := m.Call(in) // 方法 1 获取的 Method 对象反射调用方式
		out := method.Func.Call(in) // 方法 1 获取的 Method 对象反射调用方式
		for _, o := range out {
			fmt.Println("out:", o)
		}
	}
	fmt.Println()

	// Value 转原始类型
	if u1, ok := v.Interface().(User); ok {
		fmt.Println("after:", u1.Name, u1.Age)
	}
}
