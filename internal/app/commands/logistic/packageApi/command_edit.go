package packageApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (c *LogisticPackageCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	commandParts := strings.SplitN(args, " ", 2)
	if len(commandParts) != 2 {
		log.Println("wrong args", args)
		return
	}
	idx, err := strconv.Atoi(commandParts[0])
	if err != nil {
		log.Printf("fail to edit package with idx %d: %v", idx, err)
		return
	}
	title := commandParts[1]
	err = c.packageService.Update(idx, title)
	if err != nil {
		log.Println(err.Error())
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Success")
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.Get: error sending reply message to chat - %v", err)
	}
}
