package cart

// CartRepository repository interface
type CartRepository interface {
	FindByID(id string) Item
	Persist(item Item)
}
