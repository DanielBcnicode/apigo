package cart

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetCartHandle handle the GET request for the endpoint /getcart/{item}
func GetCartHandle(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	vars := mux.Vars(r)
	requestID := vars["cartId"]
	json.NewEncoder(w).Encode(map[string]string{"Valor enviado ": requestID})
}
