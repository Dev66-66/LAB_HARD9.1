package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloHandler handles requests to the /hello endpoint.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Lab 3 Go Server!")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	port := ":8082"
	fmt.Printf("Simple Go server listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
