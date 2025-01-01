package main

import (
	_ "embed"
	"log"
	"net/http"
)

func main() {
	server := NewSixelServer(NewInMemorySixelStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
