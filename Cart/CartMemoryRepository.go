package cart

import "errors"

//CartMemoryRepository implementation in memory for CartRepository
type CartMemoryRepository struct{}

var cartArray = make([]Cart, 0)

//FindByID Cart
func (repository CartMemoryRepository) FindByID(id string) (Cart, error) {
	for _, cart := range cartArray {
		if id == cart.ID {
			return cart, nil
		}
	}
	emptyCart := Cart{}
	return emptyCart, errors.New("Cart not found")
}

//Persist Cart If some cart with same ID is found, it will delete
func (repository CartMemoryRepository) Persist(item Cart) error {
	_, err := repository.FindByID(item.ID)
	if err == nil {
		removeCartFromSlice(item.ID)
	}
	cartArray = append(cartArray, item)
	return nil
}

func removeCartFromSlice(id string) {
	newArray := make([]Cart, 0)
	for _, cart := range cartArray {
		if cart.ID != id {
			newArray = append(newArray, cart)
		}
	}
	cartArray = newArray
}
