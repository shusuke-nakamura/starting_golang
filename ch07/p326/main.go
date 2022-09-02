package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f, err := ioutil.TempFile(os.TempDir(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.WriteString("Hello World!\n")
	fmt.Println(f.Name())
}
