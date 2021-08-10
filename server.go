// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	names := r.URL.Query()["name"]
	var name string
	if len(names) == 1 {
		name = names[0]
	}
	m := map[string]string{"name": name}
	enc := json.NewEncoder(w)
	enc.Encode(m)

	_, err := fmt.Fprintf(w, "Hi there, I love %s!", m["name"])

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
