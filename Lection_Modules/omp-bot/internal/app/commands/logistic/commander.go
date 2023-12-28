package logistic

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/logistic/pack"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

// Интерфейс Коммандера
type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

// Структура коммандера
type LogisticCommander struct {
	bot           *tgbotapi.BotAPI
	packCommander Commander
}

// Конструктор коммандера
func NewLogisticCommander(
	bot *tgbotapi.BotAPI,
) *LogisticCommander {
	return &LogisticCommander{
		bot: bot,
		// packageCommander
		packCommander: pack.NewLogisticPackCommander(bot),
	}
}

// Реализация Интерфейса
func (c *LogisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "package":
		c.packCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "pack":
		c.packCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
