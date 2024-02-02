package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name: "my-cookie",
			Value: "Some value 1",
		})
		fmt.Fprintln(w, "New Cookie is added, please check the devtools -> application -> cookies")
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("my-cookie")
		if err != nil {
			log.Fatal("Fail to get cookie")
		}
		io.WriteString(w, cookie.Value)
	})

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}