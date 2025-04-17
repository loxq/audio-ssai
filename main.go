package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/stream", streamHandler)
	fmt.Println("🎧 Middleware audio SSAI à l'écoute sur :8080/stream")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
