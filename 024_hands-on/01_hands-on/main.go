package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.html")
}

func main() {
	http.HandleFunc("/", foo)	
	http.HandleFunc("/dog/", dog)	
	http.Handle("/assets/images.jpeg", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))	

	http.ListenAndServe(":8080", nil)
}