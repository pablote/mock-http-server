package main

import (
	"fmt"
	"github.com/pablote/mock-server/src"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "9090"
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	log.Printf("mock-server starting on port %v\n", port)
	http.HandleFunc("/", src.NewHandler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}