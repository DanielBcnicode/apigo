package test

import (
	"testing"
	cart "github.com/DanielBcnicode/apigo/Cart"
)

func TestFindByID(t *testing.T) {
	respository := new(cart.ItemMemoryRepository)
	item, itemErr := cart.CreateItem("itemid", "identifier", "nameTest", "description test", 12.12)
	if itemErr != nil {
		t.Errorf("Item cannot be created")
	}
	respository.Persist(item)

	item, itemErr = cart.CreateItem("itemid2", "identifier2", "nameTest2", "description test2", 2.2)
	if itemErr != nil {
		t.Errorf("Item cannot be created")
	}
	respository.Persist(item)
	
	itemFromRepository, findErr := respository.FindByID("itemid")
	if findErr != nil {
		t.Errorf("Item not found")
	}

	if itemFromRepository.ID != "itemid" || itemFromRepository.Name != "nameTest" {
		t.Errorf("FindByID has returned an invalid object")

	}
}

func TestFindByCart(t *testing.T) {
	repository := new(cart.ItemMemoryRepository)
	item, itemErr := cart.CreateItem("itemid", "identifier", "nameTest", "description test", 12.12)
	if itemErr != nil {
		t.Errorf("Item cannot be created")
	}
	repository.Persist(item)

	item, itemErr = cart.CreateItem("itemid2", "identifier2", "nameTest2", "description test2", 2.2)
	if itemErr != nil {
		t.Errorf("Item cannot be created")
	}
	repository.Persist(item)
	
	cartObject, _ := cart.CreateCart("identifier")

	itemsFromRepository := repository.FindByCart(cartObject)
	
	if len(itemsFromRepository) == 0 {
		t.Errorf("Item not found")
	}

	if itemsFromRepository[0].ID != "itemid" || itemsFromRepository[0].Name != "nameTest" {
		t.Errorf("FindByCart has returned an invalid object")
	}
}