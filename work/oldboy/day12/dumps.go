package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var v = CreateStu()
	fmt.Println("转换前:", v)
	fmt.Println("----------------------------------------------------------")

	bytes, err := MarshalStruct(&v)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("我的Marshal")
		fmt.Println("byte:", bytes)
		fmt.Println("str:", string(bytes))
		fmt.Println("----------------------------------------------------------")
		bytes, _ = json.Marshal(v)
		fmt.Println("官方Marshal")
		fmt.Println("byte:", bytes)
		fmt.Println("str:", string(bytes))
	}
}

// MarshalStruct 序列化结构体为JSON
func MarshalStruct(v interface{}) ([]byte, error) {
	value := reflect.ValueOf(v)

	// 只接受 struct || &struct 两种类型
	if value.Kind() == reflect.Ptr {
		if value.Elem().Kind() != reflect.Struct {
			return nil, errors.New("只能转换结构体")
		}
		value = value.Elem()
	} else if value.Kind() != reflect.Struct {
		return nil, errors.New("只能转换结构体")
	}

	s := "{"
	for i := 0; i < value.Type().NumField(); i++ {
		field := value.Field(i)
		categoryKey := value.Type().Field(i).Tag.Get("category")

		switch field.Kind() {
		case reflect.String:
			s += fmt.Sprintf("\"%v\":\"%v\",", categoryKey, field.String())
		case reflect.Int:
			s += fmt.Sprintf("\"%v\":%v,", categoryKey, field.Int())
		}
	}
	s = s[:len(s)-1] + "}"

	return []byte(s), nil
}

type Student struct {
	ID      int    `category:"id" json:"id"`
	Name    string `category:"name" json:"name"`
	Chinese int    `category:"chinese" json:"chinese"`
	English int    `category:"english" json:"english"`
}

func CreateStu() Student {
	return Student{ID: 100, Name: "Faker", Chinese: 100, English: 100}
}
