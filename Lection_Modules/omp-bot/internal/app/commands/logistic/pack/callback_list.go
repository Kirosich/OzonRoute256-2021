package pack

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Page   int  `json:"page"`
	Vector bool `json:"vector"`
}

func (c *LogisticPackCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("LogisticPackCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	// На какой странице находимся
	page := parsedData.Page

	// Направление вверх = true | вниз = false
	vector := parsedData.Vector

	if vector {
		page += 1
	} else {
		page -= 1
	}
	// Список сущностей
	products := c.packService.List()

	// Итоговое сообщение
	listMsg := "Here all the products: \n\n"

	// Максимум страниц с проверкой на остаток

	var maxPages int

	if len(products)%numberOfPositions > 0 {
		maxPages = len(products)/numberOfPositions + 1
	} else {
		maxPages = len(products) / numberOfPositions
	}

	listMsg += fmt.Sprintf("Page: %v/%v\n", page+1, maxPages)

	FromElem := page * numberOfPositions
	ToElem := FromElem + numberOfPositions

	for i := FromElem; i < ToElem; i++ {
		if i <= len(products)-1 { // -1 так как в len идёт конкретная длина
			listMsg += fmt.Sprintf("%v", products[i].Title+"\n")
		} else {
			break
		}
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		listMsg,
	)

	serializedDataNext, _ := json.Marshal(CallbackListData{
		Page:   page,
		Vector: true,
	})

	serializedDataPrev, _ := json.Marshal(CallbackListData{
		Page:   page,
		Vector: false,
	})

	callbackPathNext := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "pack",
		CallbackName: "list",
		CallbackData: string(serializedDataNext),
	}

	callbackPathPrev := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "pack",
		CallbackName: "list",
		CallbackData: string(serializedDataPrev),
	}

	if page != 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
				tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathPrev.String()),
			),
		)
	} else {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
