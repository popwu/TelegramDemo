package libs

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GoListen(botToken string, args *[]string) {

	// 创建新的 bot 实例
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("已授权的账号 %s", bot.Self.UserName)

	// 设置更新配置
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// 获取更新通道
	updates := bot.GetUpdatesChan(u)

	// 遍历所有更新
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// 打印接收到的消息
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// 可选：回复一个确认消息
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "我已收到你的消息："+update.Message.Text)
		bot.Send(msg)
	}
}
