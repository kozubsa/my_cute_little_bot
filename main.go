package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/introphin/envopt"
	"log"
)

func main() {
	envopt.Validate("envopt.json")

	bot, err := tgbotapi.NewBotAPI(envopt.GetEnv("TELEGRAM_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.MessageConfig{}

		switch update.Message.Command() {
		case "info":
			text := ""

			if update.Message.From != nil {
				fromTemplate := "From:\n\t\tid: %d\n\t\tusername: %s\n\t\tfirst name: %s\n\t\tlast name: %s\n\t\tlanguage code: %s\n\n"
				text += fmt.Sprintf(fromTemplate,
					update.Message.From.ID,
					update.Message.From.UserName,
					update.Message.From.FirstName,
					update.Message.From.LastName,
					update.Message.From.LanguageCode,
				)
			}
			if update.Message.Chat != nil {
				chatTemplate := "Chat:\n\t\tid: %d\n\t\ttype: %s\n\t\ttitle: %s\n\t\tusername: %s\n\t\tfirst name: %s\n\t\tlast name: %s\n\t\tAllMembersAreAdmins: %v\n\t\tDescription: %s\n\t\tInviteLink: %s\n\n"
				text += fmt.Sprintf(chatTemplate,
					update.Message.Chat.ID,
					update.Message.Chat.Type,
					update.Message.Chat.Title,
					update.Message.Chat.UserName,
					update.Message.Chat.FirstName,
					update.Message.Chat.LastName,
					update.Message.Chat.AllMembersAreAdmins,
					update.Message.Chat.Description,
					update.Message.Chat.InviteLink,
				)
			}

			text += fmt.Sprintf("MessageID: %d\n", update.Message.MessageID)

			msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
		case "group":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("chat id is: %#v", update.Message.Chat.ID))
		case "user":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("user id is: %#v", update.Message.From.ID))
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "unknown command")
		}

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
