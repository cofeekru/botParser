package main

import (
	"telegramBot/src/parser"
	"telegramBot/src/tg_bot"

	"github.com/joho/godotenv"
)

var ParserHOST, ParserPORT string

func main() {
	godotenv.Load(".env")

	go parser.ServerStart()
	tg_bot.BotStart()
}
