package item

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/passsquale/telegram-bot/internal/app/path"
	"log"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *ProductItemCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("ProductItemCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v\n", parsedData),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductItemCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
