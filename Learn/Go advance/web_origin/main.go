package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler) //设置路由
	http.HandleFunc("/greet", greetHandler) //设置路由

	fmt.Println("start server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "I am hello handler,now is my time")
	if err != nil {
		return
	}
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
