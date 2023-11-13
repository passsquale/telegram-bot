package packageApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *LogisticPackageCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - get a list of packages\n"+
			"/get - get a package\n"+
			"/new - create a new package\n"+
			"/delete - delete a package\n"+
			"/edit - edit a package",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.Help: error sending reply message to chat - %v", err)
	}
}
