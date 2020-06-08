package test

import (
	"testing"

	cart "github.com/DanielBcnicode/apigo/Cart"
)

func TestCreateCart(t *testing.T) {
	cartTest, err := cart.CreateCart("identifier")
	if err != nil {
		t.Errorf("Some error ocurred ")
	}
	if len(cartTest.Items) != 0 {
		t.Errorf("New cart can't have any item inside" + string(len(cartTest.Items)))
	}

	if cartTest.ID != "identifier" {
		t.Errorf("The cart ID has changed")
	}
}

func TestAddItem(t *testing.T) {
	cartTest, err := cart.CreateCart("identifier")
	item, itemErr := cart.CreateItem("itemid", "identifier", "nameTest", "description test", 12.12)

	if err != nil || itemErr != nil {
		t.Errorf("Cant instantiate objects " + err.Error() + itemErr.Error())
	}
	cartTest.AddItem(item)

	if len(cartTest.Items) != 1 {
		t.Errorf("We can't add a item into cart.")
	}

	addItemErr := cartTest.AddItem(item)
	if addItemErr == nil {
		t.Errorf("cart.AddItem has not been sent an error")
	}

}

func TestItemExist(t *testing.T) {
	cartTest, err := cart.CreateCart("identifier")
	item, itemErr := cart.CreateItem("itemid", "identifier", "nameTest", "description test", 12.12)

	if err != nil || itemErr != nil {
		t.Errorf("Cant instantiate objects " + err.Error() + itemErr.Error())
	}
	cartTest.AddItem(item)

	if cartTest.ItemExist(("itemid")) == false {
		t.Errorf("The function cart.ItemExist has not found an existing item")
	}

}
