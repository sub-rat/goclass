package exmphttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("testing"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println(string(body))
			w.Write([]byte(`{"success": true}`))
		} else {
			if value, ok := r.URL.Query()["id"]; ok {
				fmt.Println(value)
			}
			w.Write([]byte(`{"id":1,"name":"home"}`))
		}
	})
	http.ListenAndServe(":8081", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	if value, ok := req.Header["Accept"]; ok {
		fmt.Println("header = ", value)
	}
	w.Write([]byte("hello"))
}
