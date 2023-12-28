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
	// Максимум страниц с проверкой на остаток

	var maxPages int

	if len(products)%numberOfPositions > 0 {
		maxPages = len(products)/numberOfPositions + 1
	} else {
		maxPages = len(products) / numberOfPositions
	}

	// Строчка все сущности
	listMsg := "Here all the products: \n\n"

	// Добавка строки с отображением страниц
	listMsg += fmt.Sprintf("Page: %v/%v\n", page+1, maxPages)

	// Отображение сущностей по страницам
	FromElem := page * numberOfPositions
	ToElem := FromElem + numberOfPositions

	for i := FromElem; i < ToElem; i++ {
		if i <= len(products)-1 { // -1 так как в len идёт конкретная длина
			listMsg += fmt.Sprintf("%v", products[i].Title+"\n")
		} else {
			break
		}
	}

	// Создание сообщения
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		listMsg,
	)

	// Json для кнопки следующей страницы
	serializedDataNext, _ := json.Marshal(CallbackListData{
		Page:   page,
		Vector: true,
	})

	// Json для кнопки предыдущей страницы
	serializedDataPrev, _ := json.Marshal(CallbackListData{
		Page:   page,
		Vector: false,
	})

	// Соответствующие callback'и
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

	// Switch для отображение страниц
	switch {
	case page == 0: // если первая страница, то только кнопка вперед
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
			),
		)
	case page == maxPages-1: // если страница равна максимальному количеству страниц, то кнопка только назад
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathPrev.String()),
			),
		)
	case page != 0: // остальное можно и вперед и назад
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
				tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathPrev.String()),
			),
		)
	}

	// Отправка сообщений
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
