package cart

import (
	"net/http"

	"github.com/gorilla/mux"
)

//GetCartHandle Get handle to obtain a Cart for v*/getcart/{item}
type GetCartHandle struct {
	Service GetCartService
}

// Handle handle the GET request for the endpoint v*/getcart/{item}
func (handle GetCartHandle) Handle(w http.ResponseWriter, r *http.Request) {
	// an example API handler
	vars := mux.Vars(r)
	requestID := vars["cartId"]

	cart, err := handle.Service.getCart(requestID)

	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	JSON(w, http.StatusCreated, cart)
}
