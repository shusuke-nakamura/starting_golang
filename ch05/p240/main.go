package main

import (
	"fmt"
	"foo"
)

func main() {
	t := foo.NewI()
	fmt.Println(t.Method1())
	// t.method2()
}
