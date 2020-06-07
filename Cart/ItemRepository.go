package cart

// ItemRepository repository interface
type ItemRepository interface {
	FindByID(id string) Item
	FindByCart(cart Cart) []Item
	Persist(item Item)
}
