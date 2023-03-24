package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func devices() {
	url := "http://swcm1-chel2.is74.ru:8090/api/gateways?limit=50"
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGlfa2V5X2lkIjoiNDY4ODk0OWUtYWQ4My00Y2VmLTgxYWMtMDAyZTI1NDM5NmJmIiwiYXVkIjoiYXMiLCJpc3MiOiJhcyIsIm5iZiI6MTY3OTQ5MDczNCwic3ViIjoiYXBpX2tleSJ9.cB5QSAWckp92L86iiXX6rSyjzKQzNaU8lZlJjPbhYdI"
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	//for _, b := range body {
	//	log.Print(string((b)))
	//}
	log.Println(string([]byte(body)))

}
func telegramBot() {
	//create bot
	bot, err := tgbotapi.NewBotAPI("6022325853:AAF0JDEJ_NUuLubWI-niOfCfIHGCrlJicw8")
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	//install time upd
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	//Get update from bot
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi i`m a lora_bot")
				bot.Send(msg)
			case "/getdevices":
				devices()
			}
		}
	}
}
func main() {
	log.Println("start bot")
	telegramBot()
}
