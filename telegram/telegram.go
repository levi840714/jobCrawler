package telegram

import (
	"jobCrawler/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI = nil

func Init() {
	var err error
	bot, err = tgbotapi.NewBotAPI(config.Config.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func Send(message string) {
	channel := int64(config.Config.Telegram.Channel)
	msg := tgbotapi.NewMessage(channel, message)
	// msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)

}

func OnMessage() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	}
}
