package logistic

import (
	"log"

	"github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/app/commands/logistic/pack"
	"github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
		packCommander: pack.NewPackCommander(bot),
	}
}

// Реализация Интерфейса
func (c *LogisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "pack":
		c.packCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "pack":
		c.packCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
		commandPath.CommandName = "default"
		c.packCommander.HandleCommand(msg, commandPath)
	}
}
