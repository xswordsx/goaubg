package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	log.Println("Listening on http://localhost:8888/hello")

	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler called", r.URL)
	fmt.Fprintf(w, "Hello, AUBG! - The time is %v", time.Now())
}
