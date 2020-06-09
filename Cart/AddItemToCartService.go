package cart

import "errors"

//AddItemToCartService struct for the service
type AddItemToCartService struct {
	ItemRepository ItemRepository
	CartRepository CartRepository
}

//AddItemToCart service
func (service AddItemToCartService) AddItemToCart(postItem PostItem) error {
	if serviceError := service.checkIntegrity(); serviceError != nil {
		return serviceError
	}
	if postItemErr := service.checkPostItem(postItem); postItemErr != nil {
		return postItemErr
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
func (service AddItemToCartService) checkIntegrity() error {
	if service.CartRepository == nil || service.ItemRepository == nil {
		return errors.New("AddItemToCartService has not been instantiated properly")
	}

	return nil
}

func (service AddItemToCartService) checkPostItem(postItem PostItem) error {
	if len(postItem.CartID) == 0 || len(postItem.ID) == 0 || len(postItem.Name) == 0 {
		return errors.New("Parameter wrong")
	}
	return nil
}
