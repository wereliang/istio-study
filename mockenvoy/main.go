package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	BuildVersion string
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "istio-iptables" {
		fmt.Println("hello istio iptables")
		return
	}

	go xdsServer()

	go func() {
		server := http.NewServeMux()
		server.HandleFunc("/healthz/ready", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})
		http.ListenAndServe(":15021", nil)
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name, _ := os.Hostname()
		fmt.Println(name, BuildVersion)

		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println("body:", string(body))
		for k, v := range r.Header {
			fmt.Printf("Header: %s : %s\n", k, v[0])
		}
		// w.Write([]byte(`{"ret":0,"msg":"ok","data":{"mesh_route_tag":"v1"}}`))
		w.Write([]byte(fmt.Sprintf("host:%s version :%s", name, BuildVersion)))
	})

	fmt.Println("hello envoy")
	http.ListenAndServe(":8080", nil)
}
