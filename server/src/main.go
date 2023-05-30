package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello world!"))
		if err != nil {
			log.Printf("could not write response: %v", err)
		}
	})

	log.Println("starting server...")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(fmt.Errorf("could not start server: %w", err))
	}
}
