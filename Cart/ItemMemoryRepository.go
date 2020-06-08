package cart

import "errors"

//ItemMemoryRepository implementation in memory for the Item Repository
type ItemMemoryRepository struct{}

var itemArray = make([]Item, 0)

//FindByID Cart
func (repository ItemMemoryRepository) FindByID(id string) (Item, error) {
	for _, item := range itemArray {
		if id == item.ID {
			return item, nil
		}
	}
	emptyItem := Item{}
	return emptyItem, errors.New("Item not found")
}

//FindByCart return Items slice for a cart
func (repository ItemMemoryRepository) FindByCart(cart Cart) []Item {
	items := make([]Item, 0)
	for _, item := range itemArray {
		if cart.ID == item.CartID {
			items = append(items, item)
		}
	}
	return items
}

//Persist Item
func (repository ItemMemoryRepository) Persist(item Item) error {
	_, err := repository.FindByID(item.ID)
	if err == nil {
		removeItemFromSlice(item.ID)
	}
	itemArray = append(itemArray, item)
	return nil
}

func removeItemFromSlice(id string) {
	newArray := make([]Item, 0)
	for _, item := range itemArray {
		if item.ID == id {
			newArray = append(newArray, item)
		}
	}
	itemArray = newArray
}
