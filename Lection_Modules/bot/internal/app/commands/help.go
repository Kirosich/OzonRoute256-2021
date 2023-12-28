package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - print list of commands\n"+
			"/get - get a game\n"+
			"/list - get a list of games \n"+
			"/delete - delete an existing game from list \n"+
			"/new - create a new game in list \n"+
			"/edit - edit a game \n",
	)
	c.bot.Send(msg)
}
