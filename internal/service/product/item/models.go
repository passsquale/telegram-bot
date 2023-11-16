package item

import "fmt"

var allItems = []Item{
	{Id: 0, OwnerId: 5, ProductId: 18, Title: "zero"},
	{Id: 1, OwnerId: 7, ProductId: 4, Title: "one"},
	{Id: 2, OwnerId: 1, ProductId: 45, Title: "two"},
	{Id: 3, OwnerId: 9, ProductId: 34, Title: "three"},
	{Id: 4, OwnerId: 4, ProductId: 19, Title: "four"},
	{Id: 5, OwnerId: 8, ProductId: 52, Title: "five"},
}

type Item struct {
	Id        uint64
	OwnerId   uint64
	ProductId uint64
	Title     string
}

func (i *Item) String() string {
	return fmt.Sprintf("Item{id:%v, ownerId:%v, pruductId:%v, title:%v}",
		i.Id, i.OwnerId, i.ProductId, i.Title)
}
func NewItem(id, ownerId, productId uint64, title string) *Item {
	return &Item{Id: id, OwnerId: ownerId, ProductId: productId, Title: title}
}
