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
	var productLen int
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
		productLen = len(products)
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
		productLen = len(products)
		fmt.Printf("Vector: %v, Cursor: %v \n", vector, cursor)
	}

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

	fmt.Println(productLen)
	// Switch для отображение страниц
	switch {
	case cursor == 0: // если первая страница, то только кнопка вперед
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
			),
		)
	case uint64(productLen) != limit: // если страница равна максимальному количеству страниц, то кнопка только назад
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathPrev.String()),
			),
		)
	case cursor != 0: // остальное можно и вперед и назад
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
