package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/passsquale/telegram-bot/internal/app/commands/product/item"
	"github.com/passsquale/telegram-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath)
}

type ProductCommander struct {
	bot           *tgbotapi.BotAPI
	itemCommander Commander
}

func NewProductCommander(bot *tgbotapi.BotAPI) *ProductCommander {
	return &ProductCommander{
		bot:           bot,
		itemCommander: item.NewProductItemCommander(bot),
	}
}
func (c *ProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "item":
		c.itemCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ProductCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "item":
		c.itemCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ProductCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
