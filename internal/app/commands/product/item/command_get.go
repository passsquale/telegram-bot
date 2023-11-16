package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *ProductItemCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.ParseInt(args, 0, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	item, err := c.itemService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get item with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		item.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductItemCommander.Describe: error sending reply message to chat - %v", err)
	}
}
