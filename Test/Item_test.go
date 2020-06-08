package test

import (
	"testing"
	"time"

	cart "github.com/DanielBcnicode/apigo/Cart"
)

func TestCreateItem(t *testing.T) {
	item, itemErr := cart.CreateItem("itemid", "identifier", "nameTest", "description test", 12.12)

	if itemErr != nil {
		t.Errorf("Cant instantiate item objects " + itemErr.Error())
	}

	if item.CreatedAt.Add(time.Second*3).Unix() < time.Now().Unix() || item.ID != "itemid" || item.CartID != "identifier" ||
		item.Name != "nameTest" || item.Description != "description test" ||
		item.Price != 12.12 {
		t.Errorf("Some attribute has been changed")
	}
}
