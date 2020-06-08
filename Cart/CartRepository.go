package cart

// CartRepository repository interface
type CartRepository interface {
	FindByID(id string) (Cart, error)
	Persist(item Cart) error
}
