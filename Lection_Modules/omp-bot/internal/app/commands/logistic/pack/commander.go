package pack

import (
	"log"

	"github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/app/path"
	pack "github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/service/logistic/pack"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type LogisticPackCommander struct {
	bot         *tgbotapi.BotAPI
	packService *pack.DummyPackService
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
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
