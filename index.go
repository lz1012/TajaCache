package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("初始化项目")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("init"))
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		encoder.Encode(map[string]interface{}{
			"code":    http.StatusOK,
			"message": "登陆成功",
		})
	})
	http.ListenAndServe(":80", nil)
}
