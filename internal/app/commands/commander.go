package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/passsquale/telegram-bot/internal/service/product"
	"log"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}
func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil { // If we got a message
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.CallbackQuery != nil {
			msg := tgbotapi.NewMessage(
				update.CallbackQuery.Message.Chat.ID,
				"Data: "+update.CallbackQuery.Data,
			)
			c.bot.Send(msg)
			return
		}

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)

		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		default:
			c.Default(update.Message)
		}
	}
}
