package item

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *ProductItemCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseInt(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	deleted, err := c.itemService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to remove item with idx %d: %v", idx, err)
		return
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("%v", deleted),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductItemCommander.Describe: error sending reply message to chat - %v", err)
	}
}
