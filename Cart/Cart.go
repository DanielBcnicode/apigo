package cart

import (
	"errors"
	"time"
)

// Cart basic struct of a cart
type Cart struct {
	ID        string    `json:"id"`
	Items     []Item    `json:"items"`
	CreatedAt time.Time `json:"created_at"`
}

//CreateCart like a name constructor
func CreateCart(id string) (Cart, error) {
	cart := new(Cart)
	cart.ID = id
	cart.CreatedAt = time.Now()
	cart.Items = make([]Item, 0)

	return *cart, nil
}

//AddItem add a item to the Cart
func (cart *Cart) AddItem(item Item) error {
	if cart.ItemExist(item.ID) {
		return errors.New("Item already exists")
	}
	if cart.ID != item.CartID {
		return errors.New("Item is not for the current Cart")
	}
	cart.Items = append(cart.Items, item)
	return nil
}

//ItemExist verify the existence of a item in the cart
func (cart Cart) ItemExist(itemID string) bool {
	for _, item := range cart.Items {
		if itemID == item.ID {
			return true
		}
	}

	return false
}
