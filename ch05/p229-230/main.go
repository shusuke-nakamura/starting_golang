package main

import (
	"fmt"
	"foo"
)

func main() {
	t := &foo.T{}
	fmt.Println(t.Method1())
	fmt.Println(t.Field1)
	// fmt.Println(t.method2()) // コンパイルエラー
	// fmt.Println(t.field2) // コンパイルエラー
}
