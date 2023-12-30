package pack

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c LogisticPackCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)
	idx -= 1
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"You didn't type number of entity. \n"+
				"Correct form of command: /delete__logistic__pack 3",
		)
		c.bot.Send(msg)
		log.Println("LogisticPackCommander.Delete: wrong args", args)
		return
	}

	_, err = c.packService.Remove(uint64(idx))
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"You typed incorreect number. \n"+
				"Try another.",
		)
		c.bot.Send(msg)
		log.Println("LogisticPackCommander.Delete: delete was failed", args)
		return
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("The entity %v was delete", idx+1),
	)
	c.bot.Send(msg)
}
