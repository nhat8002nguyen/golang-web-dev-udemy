package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Dog</h1>")
}

func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Nhat Nguyen</h1>")
}