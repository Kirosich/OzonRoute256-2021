package pack

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argsSlice := strings.Split(args, " ")
	fmt.Println(len(argsSlice))
	if len(argsSlice) <= 1 {

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Sorry, you didn't entered enough args.\n"+
			"Try again. Correct format of command: /edit 3 NewPack")

		log.Println("LogisticPackCommander.Edit: empty or not enough args")
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("LogisticPackCommander.Edit: error sending reply message to chat - %v", err)
		}
		return
	}

	newPack := logistic.Pack{
		Title: argsSlice[1],
	}

	packID, _ := strconv.Atoi(argsSlice[0])
	packID -= 1

	err := c.packService.Update(uint64(packID), newPack)

	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Something went wrong, sorry. \n",
		)
		c.bot.Send(msg)
		log.Println("LogisticPackCommander.Edit: Error while editing list")
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Edit of entity %v succesfully completed", packID+1))

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.Edit: error sending reply message to chat - %v", err)
	}
}
