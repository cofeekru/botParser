package tg_bot

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func handlerQuery(bot *telego.Bot, updates <-chan telego.Update, parserHOST string, parserPORT string) {
	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()

	bh.Handle(func(ctx *th.Context, update telego.Update) error {
		chatID := update.Message.Chat.ID
		message := tu.Message(
			tu.ID(chatID),
			"Привет! Я подскажу цену. Выбери категорию товара...",
		).WithReplyMarkup(mainKeyboard)

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
		host := parserHOST + ":" + parserPORT + update.CallbackQuery.Data

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
		).WithReplyMarkup(mainKeyboard)

		_, err = bot.SendMessage(ctx, message)
		if err != nil {
			log.Println("Error to send message")
		}
		return err
	}, th.CallbackDataPrefix("/"))

	bh.Start()
}
