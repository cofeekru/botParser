package tg_bot

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	exitCode = 1
)

type unit struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	URL   string `json:"url"`
}

func BotStart(botToken string, parserHOST string, parserPORT string) {
	ctx := context.Background()

	bot, err := telego.NewBot(botToken)

	if err != nil {
		log.Println("Error to start bot", err)
		os.Exit(exitCode)
	} else {
		log.Println("Starting bot..")
	}

	updates, err := bot.UpdatesViaLongPolling(ctx, nil)

	if err != nil {
		log.Println("Error to get updates")
	}

	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()

	smartphonesKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("iPhone 16").
				WithCallbackData("/catalog/iphone-16"),
			tu.InlineKeyboardButton("iPhone 16e").
				WithCallbackData("/catalog/iphone-16e"),
		),
	)

	laptopsKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("MacBook Air 13 M4").
				WithCallbackData("/catalog/macbook-air-13-M4-2025"),
		),
	)

	inlineKeyboardMain := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Смартфоны").
				WithCallbackData("smartphonesKeyboard"),
			tu.InlineKeyboardButton("Ноутбуки").
				WithCallbackData("laptopsKeyboard"),
		),
	)

	mapKeyboards := map[string](*telego.InlineKeyboardMarkup){
		"laptopsKeyboard":     laptopsKeyboard,
		"smartphonesKeyboard": smartphonesKeyboard,
	}

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := update.Message.Chat.ID
		message := tu.Message(
			tu.ID(chatID),
			"Привет! Я подскажу цену. Выбери категорию товара...",
		).WithReplyMarkup(inlineKeyboardMain)

		_, err := bot.SendMessage(ctx, message)
		if err != nil {
			log.Println("Error to send message")
		}
		return err
	}, th.CommandEqual("start"))

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := update.CallbackQuery.Message.GetChat().ID
		message := tu.EditMessageReplayMarkup(
			tu.ID(chatID),
			update.CallbackQuery.Message.GetMessageID(),
			mapKeyboards[update.CallbackQuery.Data],
		)
		_, err := bot.EditMessageReplyMarkup(ctx, message)
		if err != nil {
			log.Println("Error to send message")
		}
		return err
	}, th.CallbackDataSuffix("Keyboard"))

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := update.CallbackQuery.Message.GetChat().ID
		host := parserHOST + parserPORT + update.CallbackQuery.Data

		response, err := http.Get(host)
		if err != nil {
			log.Println("Error to get request")
			bot.SendMessage(ctx, tu.Message(
				tu.ID(chatID),
				"Данные не получены. Повторите запрос через некоторое время.",
			))
			return err
		}

		body, err := io.ReadAll(response.Body)
		var textResponse []unit
		err = json.Unmarshal(body, &textResponse)
		if err != nil {
			log.Println("Error to decode json", err)
		}

		var messageText string
		for index, unit := range textResponse {
			messageText += unit.Name + " " + unit.Price + "\n" + unit.URL + "\n\n"

			if index > 0 && (index%10 == 9 || index == len(textResponse)-1) {
				message := tu.Message(
					tu.ID(chatID),
					messageText,
				)
				_, err = bot.SendMessage(ctx, message)
				if err != nil {
					log.Println("Error to send message")
				}
				messageText = ""
			}

		}

		message := tu.Message(
			tu.ID(chatID),
			"Выбери категорию товара...",
		).WithReplyMarkup(inlineKeyboardMain)

		_, err = bot.SendMessage(ctx, message)
		if err != nil {
			log.Println("Error to send message")
		}
		return err
	}, th.CallbackDataPrefix("/"))
	bh.Start()

}
