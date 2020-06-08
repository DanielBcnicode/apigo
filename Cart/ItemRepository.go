package cart

// ItemRepository repository interface
type ItemRepository interface {
	FindByID(id string) (Item, error)
	FindByCart(cart Cart) []Item
	Persist(item Item) error
}
