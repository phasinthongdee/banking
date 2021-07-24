package main

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()
	//define rountes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
