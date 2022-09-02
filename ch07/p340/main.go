package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	/* GetメソッドでURLにアクセス */
	res, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
