package cart

import (
	"errors"
)

//GetCartService service to obtain a cart
type GetCartService struct {
	CartRepository CartRepository
}

func (service GetCartService) getCart(id string) (Cart, error) {
	emptyCart := Cart{}
	if errID := service.checkParameter(id); errID != nil {
		return emptyCart, errID
	}

	cart, findErr := service.CartRepository.FindByID(id)
	if findErr != nil {
		return emptyCart, findErr
	}

	return cart, nil
}

func (service GetCartService) checkParameter(id string) error {
	if len(id) == 0 {
		return errors.New("Cart Id can't be empty")
	}

	return nil
}
