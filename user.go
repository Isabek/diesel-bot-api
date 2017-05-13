package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
)

type User struct {
	ID int64
}

func (user *User) getId() int64 {
	return user.ID
}

func NewUser(id int64) *User {
	return &User{ID:id}
}

var usersCurrentMenu = make(map[int64]Menu)
var usersFilter = make(map[int64]map[string]string)

func (user *User) SetCurrentMenu(m *Menu) {
	usersCurrentMenu[user.getId()] = *m
}

func (user *User) GetCurrentMenu() *Menu {
	menu, ok := usersCurrentMenu[user.getId()]
	if !ok {
		return nil
	}
	return &menu
}

func (user *User) GetCurrentReplyMessage() tg.Chattable {
	menu := user.GetCurrentMenu()
	var keyboard []tg.KeyboardButton
	for _, item := range menu.Items {
		keyboard = append(keyboard, tg.KeyboardButton{Text:item.Title})
	}
	markup := tg.NewReplyKeyboard(keyboard)
	msg := tg.NewMessage(user.getId(), menu.Title)
	msg.ReplyMarkup = markup
	return msg
}

func (user *User) InitMenu(menu *Menu) {
	m := user.GetCurrentMenu()
	if m == nil {
		user.SetCurrentMenu(menu)
	}
}

func (user *User) SetFilter(key, value string) {
	mm, ok := usersFilter[user.getId()]
	if !ok {
		mm = make(map[string]string)
		usersFilter[user.getId()] = mm
	}
	mm[key] = value
}

func (user *User) ClearFilter() {
	usersFilter[user.getId()] = map[string]string{}
}

func (user *User) GetFilter() map[string]string {
	return usersFilter[user.getId()]
}

func (user *User) GetFormattedFilter() string {
	m := user.GetFilter()
	var message string
	for key, value := range m {
		message += fmt.Sprintf("%s - %s\n", key, value)
	}
	return message
}