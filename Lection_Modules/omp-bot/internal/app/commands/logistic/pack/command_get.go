package pack

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Sorry, you entered wrong number of entity. \n"+
				"Correct form of command: /get__logistic__pack 3",
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("LogisticPackCommander.Get: error sending reply message to chat - %v", err)
		}
		log.Println("LogisticPackCommander.Get: wrong args", args)
		return
	}

	product, err := c.packService.Describe(uint64(idx - 1))
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Sorry, You entered wrong number of entity. \n"+
				"Try to enter another one.",
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("LogisticPackCommander.Get: error sending reply message to chat - %v", err)
		}
		log.Printf("LogisticPackCommander.Get: fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.Get: error sending reply message to chat - %v", err)
	}
}
