package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (c *ProductItemCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	commandParts := strings.Split(args, " ")
	item, err := c.parseItem(commandParts)
	if err != nil {
		log.Printf("fail to parse string %v: %v", args, err)
		return
	}

	id, err := c.itemService.Create(*item)
	if err != nil {
		log.Printf("fail to add item %v: %v", args, err)
		return
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Created item! id: "+strconv.FormatInt(int64(id), 10),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductItemCommander.Describe: error sending reply message to chat - %v", err)
	}
}
