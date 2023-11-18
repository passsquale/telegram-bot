package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (c *ProductItemCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	commandParts := strings.SplitN(args, " ", 4)
	if len(commandParts) != 4 {
		log.Println("wrong args", args)
		return
	}
	idx, err := strconv.Atoi(commandParts[0])
	if err != nil {
		log.Printf("fail to edit package with idx %d: %v", idx, err)
		return
	}
	item, err := c.parseItem(commandParts[1:])
	if err != nil {
		log.Printf("fail to parse string %v: %v", args, err)
		return
	}
	err = c.itemService.Update(uint64(idx), *item)
	if err != nil {
		log.Println(err.Error())
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Success")
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductItemCommander.Describe: error sending reply message to chat - %v", err)
	}
}
