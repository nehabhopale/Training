package item

type Item struct {
	itemName     string
	prize        uint32
	availability bool
	rating       int
}

func (i *Item) GetItemName() string {
	return i.itemName

}
func (i *Item) Getprize() uint32 {
	return i.prize

}
func (i *Item) Getavailability() bool {
	return i.availability
}
func (i *Item) Getrating() int {
	return i.rating

}
func New(itemName string, prize uint32, availability bool, rating int) *Item {
	return &Item{
		itemName:     itemName,
		prize:        prize,
		availability: availability,
		rating:       rating,
	}
}
