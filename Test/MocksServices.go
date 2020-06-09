package test

import (
	cart "github.com/DanielBcnicode/apigo/Cart"
	"github.com/stretchr/testify/mock"
)

//MockItemRepository mock
type MockItemRepository struct {
	mock.Mock
}

//FindByID mock
func (m *MockItemRepository) FindByID(id string) (cart.Item, error) {
	args := m.Called(id)
	result := args.Get(0)
	return result.(cart.Item), args.Error(1)
}

//FindByCart mock
func (m *MockItemRepository) FindByCart(c cart.Cart) []cart.Item {
	args := m.Called(c)
	result := args.Get(0)
	return result.([]cart.Item)
}

//Persist mock
func (m *MockItemRepository) Persist(item cart.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

//MockCartRepository mock
type MockCartRepository struct {
	mock.Mock
}

//FindByID mock
func (m *MockCartRepository) FindByID(id string) (cart.Cart, error) {
	args := m.Called(id)

	return args.Get(0).(cart.Cart), args.Error(1)
}

//Persist mock
func (m *MockCartRepository) Persist(item cart.Cart) error {
	args := m.Called(item)
	return args.Error(0)
}
