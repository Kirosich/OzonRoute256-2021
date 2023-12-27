package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - print list of commands\n"+
			"/get - get a game\n"+
			"/list - get a list of games"+
			"/delete - delete an existing game from list"+
			"/new - create a new game in list"+
			"/edit - edit a game",
	)
	c.bot.Send(msg)
}
