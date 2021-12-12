package Restaurant

import (
	"fmt"
	item "hotel/item"
)

type Restaurant struct {
	restName    string
	itemDetails []*item.Item
}

func (r *Restaurant) AddNewItem(itemName string, prize uint32, availability bool, rating int) {
	r.itemDetails = append(r.itemDetails, item.New(itemName, prize, availability, rating))

}
func (r *Restaurant) GetRestDetails() {

	fmt.Println(r.restName)
	for _, v := range r.itemDetails {
		fmt.Println("*************")
		fmt.Println("name of dish", v.GetItemName())
		fmt.Println("prize of dish", v.Getprize())
		fmt.Println("availability of dish", v.Getavailability())
		fmt.Println("ratings for dish", v.Getrating())
		fmt.Println("*************")

	}

}
func (r *Restaurant) GetItemDetails(dish string) (uint32, bool) {
	for _, v := range r.itemDetails {
		if v.GetItemName() == dish {
			return v.Getprize(), v.Getavailability()
		}
	}
	return 0, false

}

func New(restName string) *Restaurant {
	return &Restaurant{
		restName: restName,
	}
}
