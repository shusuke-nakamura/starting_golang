package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	/* POSTに送信するデータを生成 */
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	res, err := http.PostForm("https://example.com/comment/post", vs)
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
