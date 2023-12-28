package pack

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__logistic_pack - print list of commands\n"+
			"/get__logistic_pack - get an entity \n"+
			"/list__logistic_pack - get a list of your entity \n"+
			"/delete__logistic__pack - delete an existing entity \n"+
			"/new__logistic__pack - create a new entity \n"+
			"/edit__logistic__pack - edit an entity",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.Help: error sending reply message to chat - %v", err)
	}
}
