package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// placeholder front page
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, it's me.")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not set.")
	}
	http.HandleFunc("/", handler) // placeholder front page
	log.Fatal(http.ListenAndServe(":"+port, nil))
}