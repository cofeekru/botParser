package tg_bot

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var (
	mainKeyboard = tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Apple").
				WithCallbackData("appleKeyboard"),
			tu.InlineKeyboardButton("Samsung").
				WithCallbackData("samsungKeyboard"),
		),
	)

	appleKeyboard = tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("iPhone").
				WithCallbackData("iphoneKeyboard"),
			tu.InlineKeyboardButton("iPad").
				WithCallbackData("ipadKeyboard"),
			tu.InlineKeyboardButton("MacBook").
				WithCallbackData("macbookKeyboard"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Назад").
				WithCallbackData("mainKeyboard"),
		),
	)

	iphoneKeyboard = tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("iPhone 16").
				WithCallbackData("/catalog/iphone-16"),
			tu.InlineKeyboardButton("iPhone 16e").
				WithCallbackData("/catalog/iphone-16e"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Назад").
				WithCallbackData("appleKeyboard"),
		),
	)

	ipadKeyboard = tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Air 11\" (2025)").
				WithCallbackData("/catalog/ipad-air-11-2025"),
			tu.InlineKeyboardButton("Mini (2024)").
				WithCallbackData("/catalog/ipad-mini-2024"),
			tu.InlineKeyboardButton("iPad (2025)").
				WithCallbackData("/catalog/ipad-2025"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Назад").
				WithCallbackData("appleKeyboard"),
		),
	)

	samsungKeyboard = tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Galaxy S").
				WithCallbackData("galaxySKeyboard"),
			tu.InlineKeyboardButton("Galaxy A").
				WithCallbackData("galaxyAKeyboard"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Назад").
				WithCallbackData("mainKeyboard"),
		),
	)

	galaxySKeyboard = tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("S23").
				WithCallbackData("/catalog/samsung-galaxy-s23"),
			tu.InlineKeyboardButton("S24").
				WithCallbackData("/catalog/samsung-galaxy-s24"),
			tu.InlineKeyboardButton("S25").
				WithCallbackData("/catalog/samsung-galaxy-s25"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Назад").
				WithCallbackData("samsungKeyboard"),
		),
	)

	galaxyAKeyboard = tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("A35 5G").
				WithCallbackData("/catalog/samsung-galaxy-a35-5g"),
			tu.InlineKeyboardButton("A36 5G").
				WithCallbackData("/catalog/samsung-galaxy-a36-5g"),
			tu.InlineKeyboardButton("A55 5G").
				WithCallbackData("/catalog/samsung-galaxy-a55-5g"),
			tu.InlineKeyboardButton("A56 5G").
				WithCallbackData("/catalog/samsung-galaxy-a56-5g"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Назад").
				WithCallbackData("samsungKeyboard"),
		),
	)

	mapKeyboards = map[string](*telego.InlineKeyboardMarkup){
		"mainKeyboard":    mainKeyboard,
		"appleKeyboard":   appleKeyboard,
		"iphoneKeyboard":  iphoneKeyboard,
		"ipadKeyboard":    ipadKeyboard,
		"samsungKeyboard": samsungKeyboard,
		"galaxySKeyboard": galaxySKeyboard,
		"galaxyAKeyboard": galaxyAKeyboard,
	}
)
