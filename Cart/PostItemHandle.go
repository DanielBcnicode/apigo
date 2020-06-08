package cart

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//PostItemHandleService service to handle the endpoint POST getcart/item
type PostItemHandleService struct {
	Service AddItemToCartService
}

// PostItemHandle handle the POST request for the endpoint /getcart/item
// This add a item to a existing cart or create a new cart and add the item
func (handle PostItemHandleService) PostItemHandle(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	cartBody := PostItem{}
	err = json.Unmarshal(body, &cartBody)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = checkPostItemBody(cartBody); err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = handle.Service.AddItemToCart(cartBody); err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	JSON(w, http.StatusCreated, cartBody)
}

func checkPostItemBody(body PostItem) error {
	withError := false
	errorMessage := ""

	if len(body.CartID) == 0 {
		withError = true
		errorMessage += "Cart Id can't be empty.\n"
	}
	if len(body.Description) == 0 {
		withError = true
		errorMessage += "Description can't be empty.\n"
	}
	if len(body.ID) == 0 {
		withError = true
		errorMessage += "Item Id can't be empty.\n"
	}
	if len(body.Name) == 0 {
		withError = true
		errorMessage += "Name can't be empty.\n"
	}

	if withError {
		return errors.New(errorMessage)
	}

	return nil
}
