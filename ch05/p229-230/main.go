package main

import (
	"fmt"
	"foo"
)

func main() {
	t := &foo.T{}
	fmt.Println(t.Method1())
	fmt.Println(t.Field1)
	// t.method2
	// t.field2
}
