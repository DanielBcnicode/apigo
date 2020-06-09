package test

import (
	"testing"

	cart "github.com/DanielBcnicode/apigo/Cart"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddItemToCart(t *testing.T) {
	MockCartR := new(MockCartRepository)
	MockItemR := new(MockItemRepository)

	Service := cart.AddItemToCartService{}
	Service.CartRepository = MockCartR
	Service.ItemRepository = MockItemR

	returnedCart, _ := cart.CreateCart("c1")
	returnedItem, _ := cart.CreateItem("i2",
		"c1", "name1", "descrip", 10.0)

	MockCartR.On("FindByID", mock.Anything).Return(returnedCart, nil)
	MockCartR.On("Persist", mock.Anything).Return(nil)
	MockItemR.On("FindByID", mock.Anything).Return(returnedItem, nil)
	MockItemR.On("Persist", mock.Anything).Return(nil)

	postItem := cart.PostItem{
		ID:          "ItemID",
		CartID:      "cartID",
		Name:        "Name",
		Description: "Description",
		Price:       100.50}

	err := Service.AddItemToCart(postItem)
	assert.Nil(t, err)
}
func TestAddItemToCartCheckIntegrity(t *testing.T) {
	//MockCartR := new(MockCartRepository)
	MockItemR := new(MockItemRepository)

	Service := cart.AddItemToCartService{}
	//Service.CartRepository = MockCartR
	Service.ItemRepository = MockItemR

	postItem := cart.PostItem{
		ID:          "",
		CartID:      "cartID",
		Name:        "Name",
		Description: "Description",
		Price:       100.50}

	err := Service.AddItemToCart(postItem)
	assert.EqualError(t, err, "AddItemToCartService has not been instantiated properly")
}

func TestAddItemToCartCheckPostItem(t *testing.T) {
	MockCartR := new(MockCartRepository)
	MockItemR := new(MockItemRepository)

	Service := cart.AddItemToCartService{}
	Service.CartRepository = MockCartR
	Service.ItemRepository = MockItemR

	postItem := cart.PostItem{
		ID:          "",
		CartID:      "cartID",
		Name:        "Name",
		Description: "Description",
		Price:       100.50}

	err := Service.AddItemToCart(postItem)
	assert.EqualError(t, err, "Parameter wrong")
}
