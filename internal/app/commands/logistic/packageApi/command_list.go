package packageApi

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/passsquale/telegram-bot/internal/app/path"
	"log"
)

func (c *LogisticPackageCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the packages: \n\n"

	packages := c.packageService.List()
	for _, p := range packages {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "demo",
		Subdomain:    "subdomain",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.List: error sending reply message to chat - %v", err)
	}
}
