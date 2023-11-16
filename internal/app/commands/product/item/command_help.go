package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ProductItemCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list_product_item - get a list of items\n"+
			"/get_product_item - get a item\n"+
			"/new_product_item - create a new item\n"+
			"/delete_product_item - delete a item\n"+
			"/edit_product_item - edit a item",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductItemCommander.Help: error sending reply message to chat - %v", err)
	}
}
