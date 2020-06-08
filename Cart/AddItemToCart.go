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
	cart, err := service.CartRepository.FindByID(postItem.CartID)
	if err != nil {
		cart, _ = CreateCart(postItem.CartID)
	}
	item, err2 := CreateItem(postItem.ID, cart.ID, postItem.Name, postItem.Description, postItem.Price)
	if err2 != nil {
		return err2
	}

	if cart.ItemExist(item.ID) {
		return errors.New("The Item already exists")
	}

	cart.Items = append(cart.Items, item)

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
