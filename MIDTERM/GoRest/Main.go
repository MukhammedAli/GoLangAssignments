package main

import (
	"GoRest/interpack/controller"
	"GoRest/interpack/domain"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	domain.Create()

	router.HandleFunc("/store/:{key}", controller.GetValueByKey).Methods("GET")

	//It works like a PUT . When I tried to implement PUT it gave me 405 error . That's why I used GET again
	//It says me that I have not allowed put method.
	//Nevertheless it works like put method but uses get keyword
	router.HandleFunc("/store/:{key}/:{value}", controller.UpdateValueByKey).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
