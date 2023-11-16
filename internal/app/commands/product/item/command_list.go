package item

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/passsquale/telegram-bot/internal/app/path"
	"log"
)

func (c *ProductItemCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the items: \n\n"

	items, err := c.itemService.List()
	if err != nil {
		log.Println("cannot get a list of items")
		return
	}
	for _, p := range items {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Subdomain:    "item",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductItemCommander.List: error sending reply message to chat - %v", err)
	}
}
