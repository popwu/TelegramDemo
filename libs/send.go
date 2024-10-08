package libs

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GoSend(botToken string, args *[]string) {
	if len(*args) != 2 {
		log.Fatal("输入： send <chatID> <message>")
		os.Exit(1)
	}
	chatIDStr := (*args)[0]
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Printf("无法将 chatID 转换为 int64: %v", err)
		return
	}
	message := (*args)[1]

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	if len(os.Args) < 2 {
		log.Fatal("Please provide an alert message as a command-line argument")
	}

	msg := tgbotapi.NewMessage(chatID, message)
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Alert message sent successfully")
}
