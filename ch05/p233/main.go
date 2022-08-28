package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `Id:"ユーザーID"`
	Name string `Name:"名前"`
	Age  uint   `Age:"年齢"`
}

func main() {
	u := User{Id: 1, Name: "Taro", Age: 32}

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
}
