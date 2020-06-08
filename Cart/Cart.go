package cart

import "time"

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

//ItemExist verify the existence of a item in the cart
func (cart Cart) ItemExist(itemID string) bool {
	for _, item := range cart.Items {
		if itemID == item.ID {
			return true
		}
	}

	return false
}
