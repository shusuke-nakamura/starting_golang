package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bs, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	/* bsは[]byte型でファイルのすべてのバイト列が入る */
	fmt.Println(bs)
}
