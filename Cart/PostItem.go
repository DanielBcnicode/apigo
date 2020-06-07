package cart

// PostItem struct used into the post body to add a item to a cart
type PostItem struct {
	ID          string  `json:"id"`
	CartID      string  `json:"cart_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
