package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body)
	base, _ := url.Parse("http://example.com") // check invalid url
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String() // base urlにたいしてreferenceをつける
	fmt.Println(endpoint)
	req, _ := http.NewRequest("GET", endpoint, nil) // structをつくっているだけ
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())       // &アンパサンドが入っている場合一度Encodeする必要がある
	req.URL.RawQuery = q.Encode() // エンコードされたものを戻す

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
