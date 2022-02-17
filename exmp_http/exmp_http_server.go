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
	// ctx := req.Context()
	// _, err := http.Post("http://exampl.com", "application/json", nil)
	// http.Error(w, "error in hello", 505)

	// select {
	// case <-time.After(3 * time.Second):
	// 	if value, ok := req.Header["Accept"]; ok {
	// 		fmt.Println("header = ", value)
	// 	}
	// case <-ctx.Done():
	// 	err := ctx.Err()
	// 	fmt.Println("server error:", err)
	// 	// http.Error(w, err.Error(), http.StatusInternalServerError)

	// }

	// if value, ok := req.Header["Accept"]; ok {
	// 	fmt.Println("header = ", value)
	// }
	w.Write([]byte("hello"))
}
