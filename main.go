package main

import (
	"log"
	"os"

	"github.com/yanzay/tbot"
)

// CharactersCount is default value
const CharactersCount int = 140

func main() {

	token := os.Getenv("MESSAGE_SHORTENER_BOT")

	bot, err := tbot.NewServer(token)
	if err != nil {
		log.Fatalf("Can't connect to telegram: %v\n", err)
	}

	bot.HandleDefault(echoHandle)

	log.Println("Starting bot...")

	err = bot.ListenAndServe()
	if err != nil {
		log.Fatalf("Can't start telegram bot: %v\n", err)
	}
}

func echoHandle(message *tbot.Message) {

	// if message length grater than 140 characters
	if len([]rune(message.Text())) > CharactersCount {

		// getting telegra.ph auth token
		data, err := GetAuth(message.From)
		if err != nil {
			log.Fatal(err)
		}

		result, _ := data.(map[string]interface{})
		result, _ = result["result"].(map[string]interface{})

		// Creating telegra.ph page
		page := Page{
			AccessToken:   result["access_token"].(string),
			Title:         "Shorted by MessageShortenerBot",
			AuthorName:    result["author_name"].(string),
			AuthorURL:     "",
			Content:       PrepareContent(message.Text()),
			ReturnContent: false,
		}

		data, err = CreatePage(page)
		if err != nil {
			log.Fatal(err)
		}

		result, _ = data.(map[string]interface{})
		result, _ = result["result"].(map[string]interface{})

		// printing result URL
		message.Reply(result["url"].(string))
		return
	}

	message.Reply(message.Text())
}
