package pack

import (
	"log"

	"github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/app/path"
	pack "github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/service/logistic/pack"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type PackCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath)
}

func NewPackCommander(bot *tgbotapi.BotAPI) PackCommander {
	return NewLogisticPackCommander(bot)
}

type LogisticPackCommander struct {
	bot         *tgbotapi.BotAPI
	packService pack.PackService
}

func NewLogisticPackCommander(
	bot *tgbotapi.BotAPI,
) *LogisticPackCommander {
	packService := pack.NewDummyPackService()

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
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
