package item

import (
	"errors"
	"fmt"
)

type ItemService interface {
	Describe(itemID uint64) (*Item, error)
	List(cursor, limit uint64) ([]Item, error)
	Create(item Item) (uint64, error)
	Update(itemID uint64, item Item) error
	Remove(itemID uint64) (bool, error)
}

var curId uint64 = 6

type DummyItemService struct {
}

func NewService() *DummyItemService {
	return &DummyItemService{}
}
func (s *DummyItemService) Describe(itemID uint64) (*Item, error) {
	if itemID < 0 || itemID >= uint64(len(allItems)) {
		return nil, errors.New("invalid index")
	}
	return &allItems[itemID], nil
}

func (s *DummyItemService) List(cursor, limit uint64) ([]Item, error) {
	totalCount := uint64(len(allItems))
	if cursor > totalCount {
		return nil, errors.New(fmt.Sprintf(
			"Incorrect cursor position %v, correct cursor positions are [0..%v]",
			cursor, totalCount,
		))
	}
	right := cursor + limit
	if right > totalCount {
		right = totalCount
	}
	return allItems[cursor:right], nil
}
func (s *DummyItemService) Create(item Item) (uint64, error) {
	if item.Title == "" || len(item.Title) > 20 {
		return 0, errors.New("invalid title")
	}

	item.Id = curId
	curId++
	allItems = append(allItems, item)
	return item.Id, nil
}
func (s *DummyItemService) Update(itemID uint64, item Item) error {
	if itemID < 0 || itemID >= uint64(len(allItems)) {
		return errors.New("invalid itemID")
	}
	allItems[itemID] = item
	return nil
}
func (s *DummyItemService) Remove(itemID uint64) (bool, error) {
	if itemID < 0 || itemID >= uint64(len(allItems)) {
		return false, errors.New("invalid index")
	}
	allItems = append(allItems[:itemID], allItems[itemID+1:]...)
	return true, nil
}
