package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
