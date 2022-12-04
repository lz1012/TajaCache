package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("初始化项目")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("init"))
	})
	http.ListenAndServe(":80", nil)
}
