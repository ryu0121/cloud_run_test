package main

import (
	"log"
	"net/http"
	"sampleapp/handler"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", handler.Handler))
}
