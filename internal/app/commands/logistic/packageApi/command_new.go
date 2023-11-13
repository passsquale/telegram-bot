package packageApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *LogisticPackageCommander) New(inputMessage *tgbotapi.Message) {
	title := inputMessage.CommandArguments()

	mypackage, err := c.packageService.New(title)
	if err != nil {
		log.Printf("fail to add package with idx %d: %v", title, err)
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"New package: "+mypackage.Title,
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.Get: error sending reply message to chat - %v", err)
	}
}
