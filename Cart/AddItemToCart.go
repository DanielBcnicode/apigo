package cart

import "errors"

//AddItemToCartService struct for the service
type AddItemToCartService struct {
	ItemRepository ItemRepository
	CartRepository CartRepository
}

//AddItemToCart service
func (service AddItemToCartService) AddItemToCart(postItem PostItem) error {
	if serviceError := checkIntegrity(service); serviceError != nil {
		return serviceError
	}

	cart, notFoundError := service.CartRepository.FindByID(postItem.CartID)
	if notFoundError != nil {
		cart, _ = CreateCart(postItem.CartID)
	}

	item, createItemError := CreateItem(postItem.ID, cart.ID, postItem.Name, postItem.Description, postItem.Price)
	if createItemError != nil {
		return createItemError
	}

	if addItemError := cart.AddItem(item); addItemError != nil {
		return addItemError
	}

	service.ItemRepository.Persist(item)
	service.CartRepository.Persist(cart)

	return nil
}

//checkIntegrity check the proper instantiation of service
func checkIntegrity(service AddItemToCartService) error {
	if service.CartRepository == nil || service.ItemRepository == nil {
		return errors.New("AddItemToCartService has not been instantiated properly")
	}

	return nil
}
