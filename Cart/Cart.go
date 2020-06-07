package cart

import "time"

// Cart basic struct of a cart
type Cart struct {
	ID        string    `json:"id"`
	Items     []Item    `json:"items"`
	CreatedAt time.Time `json:"created_at"`
}
