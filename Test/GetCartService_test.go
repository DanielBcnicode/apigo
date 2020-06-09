package test

import (
	"testing"

	cart "github.com/DanielBcnicode/apigo/Cart"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//Mocks are defined in MocksServices.go
func TestCheckParameter(t *testing.T) {
	MockCartR := new(MockCartRepository)
	Service := cart.GetCartService{}
	Service.CartRepository = MockCartR

	_, err := Service.GetCart("")

	assert.EqualError(t, err, "Cart Id can't be empty")

}

func TestCGetCart(t *testing.T) {
	MockCartR := new(MockCartRepository)
	Service := cart.GetCartService{}
	Service.CartRepository = MockCartR

	cartTest, _ := cart.CreateCart("1")
	itemTest, _ := cart.CreateItem("2", "1", "name", "desc", 34.3)
	cartTest.AddItem(itemTest)
	MockCartR.On("FindByID", mock.Anything).Return(cartTest, nil)

	cartReturned, _ := Service.GetCart("1")

	assert.Equal(t, cartTest, cartReturned)

}
