package main

import (
	"ascii-art-web/internal/server"
	"log"
	"net/http"
)

const adress = ":8080"

func main() {
	server.RegisterRoutes()

	log.Printf("Ascii-art-web running on http://localhost%s\n", adress)
	if err := http.ListenAndServe(adress, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}

}
