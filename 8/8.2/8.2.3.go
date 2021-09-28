// 자료형 메타 데이터 ( 자료에 대한 자료 )

package main

import (
	"errors"
	"fmt"
	"reflect"
)

// reflect 패키지를 이용하면 자료형 메타 데이터를 이용할 수 있음

func NewMap(key, value interface{}) interface{} {
	keyType := reflect.TypeOf(key)
	valueType := reflect.TypeOf(value)
	mapType := reflect.MapOf(keyType, valueType)
	mapValue := reflect.MakeMap(mapType)
	return mapValue.Interface()
}

// 구조체가 어떤 필드를 가지고 있는지 확인하는 방법
func FieldNames(s interface{}) ([]string, error) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("FieldNames: s is not a struct")
	}
	names := []string{}
	n := t.NumField()
	for i := 0; i < n; i++ {
		names = append(names, t.Field(i).Name)
	}
	return names, nil
}

func AppendNilError(f interface{}, err error) (interface{}, error) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return nil, errors.New("AppendNilError : f is not a function")
	}
	in, out := []reflect.Type{}, []reflect.Type{}
	for i := 0; i < t.NumIn(); i++ {
		in = append(in, t.In(i))
	}
	for i := 0; i < t.NumOut(); i++ {
		in = append(in, t.Out(i))
	}
	out = append(out, reflect.TypeOf((*error)(nil)).Elem())
	funcType := reflect.FuncOf(in, out, t.IsVariadic())
	v := reflect.ValueOf(f)
	funcValue := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		results := v.Call(args)
		results = append(results, reflect.ValueOf(&err).Elem())
		return results
	})
	return funcValue.Interface(), nil
}

func main() {
	m := NewMap("", 0).(map[string]int)
	m["abc"] = 3
	fmt.Println("NewMap : ", m)

	s := struct {
		id   int
		Name string
		Age  int
	}{}
	fmt.Print("FieldNames : ")
	fmt.Println(FieldNames(s))

	f := func() {
		fmt.Println("called")
	}
	f2, err := AppendNilError(f, errors.New("test error"))
	fmt.Println("AppendNilError.err : ", err)
	fmt.Println(f2.(func() error)())
}
