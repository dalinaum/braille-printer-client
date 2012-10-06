package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	/*resp, _ := http.Get("http://localhost:8080/")
	  defer resp.Body.Close()
	  body, _ := ioutil.ReadAll(resp.Body)
	  fmt.Printf(string(body))*/

	resp2, _ := http.PostForm("http://localhost:8080/braille",
		url.Values{"input": {"hello"}, "lang": {"en"}})
	defer resp2.Body.Close()
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Printf(string(body2))
}
