package main

import (
	"fmt"
	"reflect"
	"bytes"
	"strconv"
)

type Student struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (s *Student) String() string {
	return fmt.Sprintf("name:%s,id:%d", s.Name, s.Id)
}

func (s Student) GetId() int {
	return s.Id
}

func (s *Student) GetName() string {
	return s.Name
}

func print(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.Kind()) // ptr
	t = t.Elem()
	fmt.Println(t.Name())    //Student
	fmt.Println(t.PkgPath()) //main
	for i := 0; i < t.NumField(); i ++ {
		field := t.Field(i)
		fmt.Println(field)
		fmt.Println("jsonkey:", field.Tag.Get("json"))
	}

	field, ok := t.FieldByName("Name")
	if ok {
		fmt.Printf("%#v\n", field)
	}

	for i := 0; i < t.NumMethod(); i ++ {
		method := t.Method(i)
		fmt.Println("-----", method.Name) // GetId 可以拿到 GetName String 拿不到  (s *Student)String() 指针类型拿不到
	}

	v := reflect.ValueOf(x).Elem() // 通过Elem可以取到真实的数据 使用之前必须先调用
	vfield := v.FieldByName("Name")
	fmt.Println(vfield.String())

	method := v.MethodByName("GetId")

	ret := method.Call(nil)
	fmt.Println("======",ret[0].String())
}

func marshal(x interface{}) string {
	t := reflect.TypeOf(x).Elem()
	v := reflect.ValueOf(x).Elem()
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "{\n")
	for i := 0; i < t.NumField(); i ++ {
		fieldt := t.Field(i)
		fieldv := v.Field(i)
		jsonKey := fieldt.Tag.Get("json")
		if jsonKey == "" {
			jsonKey = fieldt.Name
		}

		var jsonValue string
		switch fieldt.Type.Kind() {
		case reflect.Int:
			jsonValue = strconv.Itoa(int(fieldv.Int()))
		case reflect.String:
			jsonValue = `"` + fieldv.String() + `""`
		}
		fmt.Fprintf(buf, "    \"%s\":%s,\n", jsonKey, jsonValue)
	}
	fmt.Fprintf(buf, "}\n")
	return buf.String()

}

func main() {
	s := &Student{
		Name: "dcl",
		Id:   1,
	}
	print(s)
	fmt.Println(marshal(s))


	s1 := []string{"hello", "world"}
	s2 := s1
	//fmt.Println( s1 == s2 )//不可以
	fmt.Println( reflect.DeepEqual(s1, s2 ))// 可以 true
	// 数组 一个一个元素进行比较
	// struct 一个一个字段进行比较
}
