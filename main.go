package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" from:"name" msg:"名字必填"`
	Age  int    `json:"age" from:"age" msg:"年龄必填"`
}

func (this Person) GetName() string {
	return this.Name
}

func (this *Person) SetName(name string) {
	this.Name = name
}

func main() {
	p := Person{
		Name: "小明",
		Age:  20,
	}
	v := reflect.ValueOf(&p)
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	b := v.MethodByName("SetName").Call(params)
	fmt.Println(b)
	a := v.MethodByName("GetName").Call(nil)
	fmt.Println(a)
	fmt.Printf("%#v", p)
}
