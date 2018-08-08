package web

import (
	"../components"
	"../generator"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	minLength := components.ConvertToInt(params["min"])
	noNumbers := components.ConvertToInt(params["numbers"])
	noSymbols := components.ConvertToInt(params["symbols"])
	noPasswords := components.ConvertToInt(params["passwords"])

	if minLength == 0 || noNumbers == 0 || noSymbols == 0 || noPasswords == 0 {
		ValidationError(w, r)
		return
	}

	passwords := make([]string, noPasswords)
	for i := 0; i < noPasswords; i++ {
		data, err := generator.GeneratePassword(minLength, noNumbers, noSymbols)
		if err != nil {
			log.Print(err)
			ServerError(w, r)
			return
		} else {
			passwords[i] = data
		}
	}

	components.SetJsonHeader(w, http.StatusOK)
	json.NewEncoder(w).Encode(passwords)
}

func NotFoundError(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"error":   "Not Found",
		"type":    "404",
		"message": fmt.Sprintf("The defined route '%s' has not been found", r.URL.Path),
	}

	components.SetJsonHeader(w, http.StatusNotFound)
	json.NewEncoder(w).Encode(data)
}

func ValidationError(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"error":   "Validation Error",
		"type":    "422",
		"message": "The defined params are not valid or not allowed",
	}

	components.SetJsonHeader(w, http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(data)
}

func ServerError(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"error":   "Server Error",
		"type":    "500",
		"message": "Unable to process the request",
	}

	components.SetJsonHeader(w, http.StatusInternalServerError)
	json.NewEncoder(w).Encode(data)
}
