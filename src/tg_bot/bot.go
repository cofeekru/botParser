package tg_bot

import (
	"context"

	"log"

	"os"

	"github.com/mymmrac/telego"
)

const (
	exitCode = 1
)

type unit struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	URL   string `json:"url"`
}

func BotStart(botToken string, parserHOST string, parserPOST string) {
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

	handlerQuery(bot, updates, parserHOST, parserPOST)

}
