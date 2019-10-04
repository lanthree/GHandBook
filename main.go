package main

import (
	"handbook/handlers"
	"log"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/random_cards", handlers.RandomCardsHandler)
	http.HandleFunc("/", handlers.NotFoundHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}
