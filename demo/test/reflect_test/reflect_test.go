package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	obj := Person{Id: "123", Name: "张三", Age: 50}
	if isStruct(reflect.TypeOf(obj)) {
		doStruct(obj)
	}
	if isStructPtr(reflect.TypeOf(&obj)) {
		doPointer(&obj)
	}
}

func doStruct(obj interface{}) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	fmt.Println(objT)
	fmt.Println(objV)
	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < objT.NumField(); i++ {
		field := objT.Field(i)
		value := objV.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	// 获取方法
	for i := 0; i < objT.NumMethod(); i++ {
		m := objT.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
	m := objV.MethodByName("Hello")
	args := []reflect.Value{
		reflect.ValueOf("你好"),
		reflect.ValueOf(20),
	}
	ret := m.Call(args)
	fmt.Println(ret)

}
func doPointer(obj interface{}) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	fmt.Println(objT)
	fmt.Println(objV)
	// 获取方法字段
	for i := 0; i < objT.Elem().NumField(); i++ {
		field := objT.Elem().Field(i)
		value := objV.Elem().Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	// 获取方法
	for i := 0; i < objT.NumMethod(); i++ {
		m := objT.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
	m1 := objV.MethodByName("HelloPtr")
	args1 := []reflect.Value{
		reflect.ValueOf("你好"),
		reflect.ValueOf(20),
	}
	ret1 := m1.Call(args1)
	fmt.Println(ret1)

	m2 := objV.MethodByName("Hello")
	args2 := []reflect.Value{
		reflect.ValueOf("你好"),
		reflect.ValueOf(20),
	}
	ret2 := m2.Call(args2)
	fmt.Println(ret2)
}
func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

type Person struct {
	Id   string
	Name string
	Age  int
}

func (s *Person) HelloPtr(param string, age int) string {
	fmt.Println(age + s.Age + 10)
	return "hello_ptr:" + s.Name + ":" + param
}
func (s Person) Hello(param string, age int) string {
	fmt.Println(age + s.Age)
	return "hello:" + s.Name + ":" + param
}
