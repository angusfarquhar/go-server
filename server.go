// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
