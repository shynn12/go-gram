package main

import (
	"cmd-gram-blockchain/internal/api"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	api := api.New(&mux.Router{})

	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
