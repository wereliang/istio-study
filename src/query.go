package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func GetData(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func main() {
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("call query")
		var basic, recommend string
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			basic = GetData("http://basicinfo/basic")
		}()

		go func() {
			defer wg.Done()
			recommend = GetData("http://recommend/recommend")
		}()

		wg.Wait()
		w.Write([]byte(fmt.Sprintf(`{"basic":%s,"recommend":%s}`, basic, recommend)))
	})
	http.ListenAndServe(":8080", nil)

}
