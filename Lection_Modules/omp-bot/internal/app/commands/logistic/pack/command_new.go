package pack

import (
	"fmt"
	"log"

	"github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments() // Title of entity
	newPack := logistic.Pack{
		Title: args,
	}

	if args == "" {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Вы не ввели название сущности. \n"+
				"Правильный пример команды: /new__logistic__pack kek",
		)

		c.bot.Send(msg)
		log.Println("LogisticPackCommander.Get: wrong args", args)
		return
	}

	_, err := c.packService.Create(newPack)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Something went wrong, sorry. \n",
		)
		c.bot.Send(msg)
		log.Println("LogisticPackCommander.New: Error while creating")
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Сущность %v добавлена. \n", newPack.Title),
	)

	c.bot.Send(msg)

}
