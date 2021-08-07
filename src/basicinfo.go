package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/basic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("call basic")
		w.Write([]byte(`{"name":"kubernetes in action","price":100}`))
	})
	http.ListenAndServe(":8080", nil)
}
