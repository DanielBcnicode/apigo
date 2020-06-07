package cart

//AddItemToCartService struct for the service
type AddItemToCartService struct {
	itemRepository ItemRepository
	cartRepository CartRepository
}

//AddItemToCart service
func (service AddItemToCartService) AddItemToCart(postItem PostItem) error {
	cart, err := service.cartRepository.FindByID(postItem.CartID)
	item, err = CreateItem(cart.ID, postItem.Name, postItem.Description, postItem.Price)
	cart.Items = append(cart.Items, item)

	service.itemRepository.Persist(item)
	service.cartRepository.Persist(cart)

	return nil
}
