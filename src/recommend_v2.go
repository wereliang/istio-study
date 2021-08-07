package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/recommend", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("call recommend v2")
		resp, err := http.Get("http://star/star")
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(fmt.Sprintf("get star error %s", err)))
			return
		}
		defer resp.Body.Close()
		star, err := ioutil.ReadAll(resp.Body)
		w.Write([]byte(fmt.Sprintf(`{"comment":"what a fine book!","star":"%s"}`, string(star))))
	})
	http.ListenAndServe(":8080", nil)
}
