package packageApi

import (
	"github.com/passsquale/telegram-bot/internal/app/path"
	"github.com/passsquale/telegram-bot/internal/service/logistic/mypackage"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PackageCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type LogisticPackageCommander struct {
	bot            *tgbotapi.BotAPI
	packageService *mypackage.DummyPackageService
}

func NewLogisticPackageCommander(bot *tgbotapi.BotAPI) *LogisticPackageCommander {
	packageService := mypackage.NewService()
	return &LogisticPackageCommander{
		bot:            bot,
		packageService: packageService,
	}
}

func (c *LogisticPackageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LogisticPackageCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
