package packageApi

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *LogisticPackageCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	deleted, err := c.packageService.Delete(idx)
	if err != nil {
		log.Printf("fail to remove package with idx %d: %v", idx, err)
		return
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("%v", deleted),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.Get: error sending reply message to chat - %v", err)
	}
}
