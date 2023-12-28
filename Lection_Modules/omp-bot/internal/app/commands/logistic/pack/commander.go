package pack

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/pack"
)

type LogisticPackCommander struct {
	bot         *tgbotapi.BotAPI
	packService *pack.Service
}

func NewLogisticPackCommander(
	bot *tgbotapi.BotAPI,
) *LogisticPackCommander {
	packService := pack.NewService()

	return &LogisticPackCommander{
		bot:         bot,
		packService: packService,
	}
}

func (c *LogisticPackCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LogisticPackCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LogisticPackCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
