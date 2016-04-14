package main

import (
	"fmt"
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
		fmt.Println(message)
	}
}
