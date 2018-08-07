package main

import (
	"./generator"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func getPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		generator.GeneratePassword(
			validateInputs(params["min"]),
			validateInputs(params["numbers"]),
			validateInputs(params["symbols"])))

}

func validateInputs(input string) int {
	i2, err := strconv.Atoi(input)
	if err == nil {
		return i2
	}
	return 0
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/password/{min}/{numbers}/{symbols}", getPassword).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
