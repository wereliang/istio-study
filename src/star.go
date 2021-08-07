package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/star", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("call start")
		w.Write([]byte("*****"))
	})
	http.ListenAndServe(":8080", nil)
}
