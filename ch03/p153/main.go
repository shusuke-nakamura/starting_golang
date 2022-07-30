package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("NumberCPU: %dn\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Println("Version: %s\n", runtime.Version())
}
