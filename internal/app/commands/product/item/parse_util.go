package item

import (
	"errors"
	"github.com/passsquale/telegram-bot/internal/service/product/item"
	"strconv"
)

func (c *ProductItemCommander) parseItem(commandParts []string) (*item.Item, error) {
	if len(commandParts) != 3 {
		return nil, errors.New("wrong arguments")
	}
	ownerId, err := strconv.ParseInt(commandParts[0], 10, 64)
	if err != nil {
		return nil, err
	}
	productId, err := strconv.ParseInt(commandParts[1], 10, 64)
	if err != nil {
		return nil, err
	}
	title := commandParts[2]

	return &item.Item{OwnerId: uint64(ownerId), ProductId: uint64(productId), Title: title}, nil
}
