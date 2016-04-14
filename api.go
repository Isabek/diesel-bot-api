package main

import (
	"log"
	"os"
	"time"

	"github.com/tucnak/telebot"
)

var dieselBot *telebot.Bot

func main() {

	token := os.Getenv("SECRET")
	if token == "" {
		log.Panic("Please run this program with SECRET parameter")
	}

	if bot, err := telebot.NewBot(token); err != nil {
		log.Panic(err)
	} else {
		dieselBot = bot
	}

	dieselBot.Messages = make(chan telebot.Message, 1000)
	go messages()

	dieselBot.Start(1 * time.Second)
}

func messages() {
	for message := range dieselBot.Messages {
		if message.Text == "/searchflat" {
			dieselBot.SendMessage(message.Sender, "Количество комнат в квартире?", &telebot.SendOptions{
				ReplyMarkup: telebot.ReplyMarkup{
					ForceReply:         true,
					Selective:          true,
					OneTimeKeyboard:    true,
					ResizeKeyboard:     true,
					HideCustomKeyboard: true,
					CustomKeyboard: [][]string{
						[]string{"1", "2", "3"},
						[]string{"1-2", "1-3", "3-..."},
					},
				},
			})
		}
	}
}
