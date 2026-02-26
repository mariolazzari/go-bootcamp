package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := fmt.Fprintf(w, "Ciao Mario!")
		if err != nil {
			log.Fatalf("Error writing response", err)
		}
		log.Println("Bytes writte:", bytes)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
