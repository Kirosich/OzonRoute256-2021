package pack

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const numberOfPositions = 5

func (c *LogisticPackCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products := c.packService.List()

	// Максимум страниц с проверкой на остаток

	var maxPages int

	if len(products)%numberOfPositions > 0 {
		maxPages = len(products)/numberOfPositions + 1
	} else {
		maxPages = len(products) / numberOfPositions
	}

	// Добавление в сообщение указание страниц
	outputMsgText += fmt.Sprintf("Page: 1/%v \n", maxPages)

	// Добавление в сообщение списка продуктов
	for i := 0; i < numberOfPositions; i++ {
		outputMsgText += products[i].Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Page:   0,
		Vector: true,
	})

	callbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "pack",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.List: error sending reply message to chat - %v", err)
	}
}
