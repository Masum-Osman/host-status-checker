package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next(w, r)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, Welcome Web Address!")

	// fs := http.FileServer(http.Dir("./static/example.html"))
}

func main() {
	http.HandleFunc("/status", logging(status))
	http.ListenAndServe(":40040", nil)
}
