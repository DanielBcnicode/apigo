package cart

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON return the json response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Del("Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR return an Error response
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
