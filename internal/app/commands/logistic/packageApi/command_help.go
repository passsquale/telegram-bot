package packageApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *LogisticPackageCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list packages\n"+
			"/get - get package",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.Help: error sending reply message to chat - %v", err)
	}
}
