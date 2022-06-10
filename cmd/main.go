package main

import (
	services "github.com/Xanonymous-GitHub/imgur-services"
	"log"
	"net/http"
	"os"
)

func main() {
	// add all api endpoint function here.
	http.HandleFunc("/upload", services.Image)

	// you can also custom dev server's port, default is 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
