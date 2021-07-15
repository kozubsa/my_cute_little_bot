package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var token = ""

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		// check Message, group, user
		if update.ChannelPost != nil {
			continue
		}
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.Chat == nil {
			continue
		}
		if update.Message.From == nil {
			continue
		}
		if update.Message.ForwardFromChat != nil &&
			update.Message.From != nil {
			integrationMsg := fmt.Sprintf("Новая интеграция:\n")
			integrationMsg += fmt.Sprintf("User id:\n%d\n", update.Message.From.ID)
			integrationMsg += fmt.Sprintf("User first_name:\n%s\n", update.Message.From.FirstName)
			integrationMsg += fmt.Sprintf("User last_name:\n%s\n", update.Message.From.LastName)
			integrationMsg += fmt.Sprintf("User username:\n%s\n", update.Message.From.UserName)
			integrationMsg += fmt.Sprintf("Forward chat id:\n%d\n", update.Message.ForwardFromChat.ID)
			integrationMsg += fmt.Sprintf("Forward chat title:\n%s\n", update.Message.ForwardFromChat.Title)

			msg := tgbotapi.NewMessage(292572266, integrationMsg)

			if _, err := bot.Send(msg); err != nil {
				return
			}
			continue
		}

		if update.Message.NewChatMembers != nil &&
			update.Message.From != nil {
			integrationMsg := fmt.Sprintf("Новая интеграция:\n")
			integrationMsg += fmt.Sprintf("User id:\n%d\n", update.Message.From.ID)
			integrationMsg += fmt.Sprintf("User first_name:\n%s\n", update.Message.From.FirstName)
			integrationMsg += fmt.Sprintf("User last_name:\n%s\n", update.Message.From.LastName)
			integrationMsg += fmt.Sprintf("User username:\n%s\n", update.Message.From.UserName)
			integrationMsg += fmt.Sprintf("New chat id:\n%d\n", update.Message.Chat.ID)
			integrationMsg += fmt.Sprintf("Chat title:\n%s\n", update.Message.Chat.Title)
			integrationMsg += fmt.Sprintf("Chat type:\n%s\n", update.Message.Chat.Type)

			msg := tgbotapi.NewMessage(292572266, integrationMsg)

			if _, err := bot.Send(msg); err != nil {
				return
			}
			continue
		}

		// var maxSizePicture tgbotapi.PhotoSize

		// // телеграм возвращает фотографии в нескольких версиях
		// // сохраняем самую большую
		// for _, photo := range *update.Message.Photo {
		// 	if maxSizePicture.FileSize > photo.FileSize {
		// 		continue
		// 	}
		// 	maxSizePicture = photo
		// }
		//
		// fileUrl, err := bot.GetFileDirectURL(maxSizePicture.FileID)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		//
		// file := "./files/"+uuid.New().String()+".png"
		// if err := DownloadFile(file, fileUrl); err != nil {
		// 	log.Fatal(err)
		// }
		//
		// log.Println("Downloaded: " + fileUrl)
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
