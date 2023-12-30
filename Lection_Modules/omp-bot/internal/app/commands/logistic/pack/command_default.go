package pack

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Sorry, I don't know understand You.\nTry to /help__logistic__pack - to get list of my commands")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.Default: error sending reply message to chat - %v", err)
	}
}
