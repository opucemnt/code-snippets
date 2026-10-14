// 简单的 HTTP Hello 服务。
//
// 运行: go run hello_server.go
// 访问: curl "http://localhost:8080/hello?name=gopher"
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "hello, %s!\n", name)
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}
