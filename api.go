package main

import (
	"log"

	_ "github.com/vteremasov/diesel-scraper"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	START = "/start"
	BACK = "Назад"
)

func main() {

	token := "201936101:AAF2-w_JWM4B3COqv6nsWXfwPl4vsIn4QiM"
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	bot.Debug = !true
	log.Printf("Authorization account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	menu := InitMenu()

	updates, err := bot.GetUpdatesChan(u); for update := range updates {
		if update.Message == nil {
			continue
		}
		user := NewUser(update.Message.Chat.ID)
		user.InitMenu(menu)

		if user.GetCurrentMenu().Slug == MAIN_MENU_TYPE {
			user.ClearFilter()
		}

		if update.Message.Text == START {
			user.ClearFilter()
			user.SetCurrentMenu(menu)
			msg := user.GetCurrentReplyMessage()
			bot.Send(msg)
		} else if update.Message.Text == BACK {
			user.SetCurrentMenu(user.GetCurrentMenu().Prev())
			msg := user.GetCurrentReplyMessage()
			bot.Send(msg)
		} else {
			user.SetFilter(user.GetCurrentMenu().Slug, update.Message.Text)
			if user.GetCurrentMenu().IsExistNextMenu(update.Message.Text) {
				user.SetCurrentMenu(user.GetCurrentMenu().Next(update.Message.Text))
				msg := user.GetCurrentReplyMessage()
				bot.Send(msg)
			} else {
				msg := tg.NewMessage(user.getId(), user.GetFormattedFilter())
				msg.ReplyMarkup = tg.NewRemoveKeyboard(false)
				bot.Send(msg)
			}
		}
	}
}