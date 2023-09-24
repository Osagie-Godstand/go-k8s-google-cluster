package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", myName)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func myName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My name is Osagie Godstand. You can also call me Desmond.")

}
