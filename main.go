package main

import (
	"os"
	"telegramBot/src/parser"
	"telegramBot/src/tg_bot"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	var parserPORT = os.Getenv("PORT")
	parserHOST := os.Getenv("HOST")
	botToken := os.Getenv("TOKEN")

	go parser.ServerStart(parserHOST, parserPORT)
	tg_bot.BotStart(botToken, parserHOST, parserPORT)
}
