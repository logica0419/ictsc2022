package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, ICTSC2022 Contestant!\n"))
		if err != nil {
			log.Printf("error: %v", err)
		}
		log.Printf("successfully sent response")
	})

	log.Panic(http.ListenAndServe(":8080", nil))
}
