package pack

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CallbackListData struct {
	CursorI int  `json:"cursori"`
	Vector  bool `json:"vector"`
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

	// Где находится курсор
	cursor := parsedData.CursorI
	vector := parsedData.Vector

	// Строчка все сущности
	listMsg := "Here all the products: \n\n"

	// Список сущностей
	if vector {
		cursor += int(limit)

		products := c.packService.List(uint64(cursor), limit)
		for _, val := range products {
			listMsg += val.Title
			listMsg += "\n"
		}
		fmt.Printf("Vector: %v, Cursor: %v \n", vector, cursor)
	} else {
		cursorI := cursor
		fmt.Println(cursor, cursor-int(limit))
		for i := cursor; i > cursorI-int(limit); i-- {
			fmt.Println(i)
			if i > 0 {
				cursor--
			} else {
				break
			}
		}
		products := c.packService.List(uint64(cursor), limit)
		for _, val := range products {
			listMsg += val.Title
			listMsg += "\n"
		}
		fmt.Printf("Vector: %v, Cursor: %v \n", vector, cursor)
	}

	// // Максимум страниц с проверкой на остаток

	// var maxPages int

	// if len(products)%numberOfPositions > 0 {
	// 	maxPages = len(products)/numberOfPositions + 1
	// } else {
	// 	maxPages = len(products) / numberOfPositions
	// }

	// // Добавка строки с отображением страниц
	// listMsg += fmt.Sprintf("Page: %v/%v\n", page+1, maxPages)

	// Отображение сущностей по страницам
	// FromElem := page * numberOfPositions
	// ToElem := FromElem + numberOfPositions

	// for i := FromElem; i < ToElem; i++ {
	// 	if i <= len(products)-1 { // -1 так как в len идёт конкретная длина
	// 		listMsg += fmt.Sprintf("%v", products[i].Title+"\n")
	// 	} else {
	// 		break
	// 	}
	// }

	// Создание сообщения
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		listMsg,
	)

	serializedDataNext, _ := json.Marshal(CallbackListData{
		CursorI: cursor,
		Vector:  true,
	})

	serializedDataPrev, _ := json.Marshal(CallbackListData{
		CursorI: cursor,
		Vector:  false,
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

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
			tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathPrev.String()),
		),
	)

	// Switch для отображение страниц
	// switch {
	// case page == 0: // если первая страница, то только кнопка вперед
	// 	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
	// 		),
	// 	)
	// case page == maxPages-1: // если страница равна максимальному количеству страниц, то кнопка только назад
	// 	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathPrev.String()),
	// 		),
	// 	)
	// case page != 0: // остальное можно и вперед и назад
	// 	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
	// 			tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathPrev.String()),
	// 		),
	// 	)
	// }

	// Отправка сообщений
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
