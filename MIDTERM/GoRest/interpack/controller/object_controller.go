package controller

import (
	"GoRest/interpack/domain"

	"github.com/gorilla/mux"

	"encoding/json"
	"net/http"
	"sync"
)

var (
	mu sync.Mutex
)

func GetValueByKey(response http.ResponseWriter, request *http.Request) {
	Getparameters := mux.Vars(request)
	mu.Lock()
	getval := domain.GetValue(Getparameters["key"])
	mu.Unlock()

	if len(getval) == 0 {
		json.NewEncoder(response).Encode("No value")
		return
	}

	json.NewEncoder(response).Encode(getval)
}

func UpdateValueByKey(response http.ResponseWriter, request *http.Request) {
	Getparameters := mux.Vars(request)

	key := Getparameters["key"]
	value := Getparameters["value"]

	if len(key) == 0 {
		json.NewEncoder(response).Encode("Empty key")
		return
	}

	mu.Lock()
	domain.PostValue(key, value)
	mu.Unlock()

	json.NewEncoder(response).Encode("New Key: " + key + "new value: " + value)
}
