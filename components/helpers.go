package components

import (
	"net/http"
	"strconv"
)

func ConvertToInt(input string) int {
	number, err := strconv.Atoi(input)
	if err == nil {
		return number
	}
	return 0
}

func SetJsonHeader(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
}
