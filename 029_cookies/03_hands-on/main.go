package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/readcount", readCount)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	countCookie, err := req.Cookie("visited-count")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "visited-count",
			Value: "1",
			Path:  "/",
		})
	} else {
		previousCount, err := strconv.Atoi(countCookie.Value)
		if err != nil {
			log.Fatal("Something's wrong when get the previous count")
		}

		http.SetCookie(w, &http.Cookie{
			Name: "visited-count",
			Value: strconv.Itoa(previousCount+1),
			Path: "/",
		})
	}

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func readCount(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("visited-count")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Visited count:", c.Value)
}
