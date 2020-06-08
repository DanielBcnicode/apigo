package test

import (
	"testing"

	cart "github.com/DanielBcnicode/apigo/Cart"
)

func TestCartMemoryFindByID(t *testing.T) {
	repository := new(cart.CartMemoryRepository)

	cartObject, _ := cart.CreateCart("identifier")
	repository.Persist(cartObject)

	cartObject2, _ := cart.CreateCart("identifier2")
	repository.Persist(cartObject2)

	cartByFind, findErr := repository.FindByID("identifier")
	if findErr != nil || cartByFind.ID != "identifier" {
		t.Errorf("Cart not found")
	}
}

func TestCartMemoryFindPersist(t *testing.T) {
	repository := new(cart.CartMemoryRepository)

	cartObject, _ := cart.CreateCart("identifier")
	repository.Persist(cartObject)

	cartObject2, _ := cart.CreateCart("identifier")
	repository.Persist(cartObject2)

	cartByFind, findErr := repository.FindByID("identifier")
	if findErr != nil || cartByFind.ID != "identifier" {
		t.Errorf("Cart not found")
	}
}
