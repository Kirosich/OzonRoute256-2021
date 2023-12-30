package pack

import (
	"encoding/json"
	"log"

	"github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const limit uint64 = 5

func (c *LogisticPackCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	cursorI := 0

	products, err := c.packService.List(uint64(cursorI), limit)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Something went wrong, sorry. \n",
		)
		c.bot.Send(msg)
		log.Println("LogisticPackCommander.List: Error while getting list")
		return
	}

	// Максимум страниц с проверкой на остаток

	// var maxPages int

	// if len(products)%numberOfPositions > 0 {
	// 	maxPages = len(products)/numberOfPositions + 1
	// } else {
	// 	maxPages = len(products) / numberOfPositions
	// }

	// Добавление в сообщение указание страниц
	// outputMsgText += fmt.Sprintf("Page: 1/%v \n", maxPages)

	// Добавление в сообщение списка продуктов
	for _, val := range products {
		outputMsgText += val.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		CursorI: cursorI,
		Vector:  true,
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

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.List: error sending reply message to chat - %v", err)
	}
}
