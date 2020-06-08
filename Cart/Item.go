package cart

import (
	"time"
)

// Item basic structure of an item
type Item struct {
	ID          string    `json:"id"`
	CartID      string    `json:"cart_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateItem Item constructor
func CreateItem(itemID string, cartID string, name string, description string, price float32) (Item, error) {
	item := new(Item)

	item.ID = itemID //uuid.NewV4().String()
	item.CartID = cartID
	item.Name = name
	item.Description = description
	item.Price = price
	item.CreatedAt = time.Now()

	return *item, nil
}
