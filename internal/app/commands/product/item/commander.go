package item

import (
	"github.com/passsquale/telegram-bot/internal/app/path"
	"github.com/passsquale/telegram-bot/internal/service/product/item"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ItemCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type ProductItemCommander struct {
	bot         *tgbotapi.BotAPI
	itemService *item.DummyItemService
}

func NewProductItemCommander(bot *tgbotapi.BotAPI) *ProductItemCommander {
	itemService := item.NewService()
	return &ProductItemCommander{
		bot:         bot,
		itemService: itemService,
	}
}

func (c *ProductItemCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ProductItemCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ProductItemCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
