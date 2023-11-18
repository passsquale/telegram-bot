package item

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *ProductItemCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help_product_item - help\n"+
			"/list_product_item - get a list of items\n"+
			"/get_product_item {id}- get an item\n"+
			"/new_product_item {ownerID} {productID} {title}- create a new item\n"+
			"/delete_product_item {id} - delete an item\n"+
			"/edit_product_item {id} {ownerID} {productID} {title} - edit an item",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductItemCommander.Help: error sending reply message to chat - %v", err)
	}
}
