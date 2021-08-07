package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/recommend", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("call recommend v1")
		w.Write([]byte(`{"comment":"what a fucking book!"}`))
	})
	http.ListenAndServe(":8080", nil)
}
