package main

import (
	"fmt"
	"log"
	"os"
	"telegram-demo/libs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run main.go <操作> [参数...]")
		return
	}

	action := os.Args[1]
	args := os.Args[2:]

	// 从环境变量获取 bot token
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN 环境变量未设置")
		os.Exit(1)
	}

	switch action {
	// case "hello":
	// 	libs.DoHello(&args)
	// case "upload":
	// 	libs.DoUpload(&args)
	case "send":
		libs.GoSend(botToken, &args)
	case "listen":
		libs.GoListen(botToken, &args)
	// case "anotherthing":
	// 	anotherThing(args)
	default:
		fmt.Printf("未知操作: %s\n", action)
	}
}
