package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Message struct {
	From string
	Text string
}

func bot(input <-chan Message, output chan<- Message) {
	for msg := range input {
		var reply string

		switch strings.ToLower(msg.Text) {
		case "привет":
			reply = "Привет, чем могу помочь?"
		case "как дела":
			reply = "Нормально, как у Ваc?"
		case "пойдет, расскажи как прошел твой день":
			reply = "Впринципе неплохо"
		case "мне пора идти":
			reply = "Удачи хорошего дня"
		default:
			reply = "Не понял о чем Вы"

		}
		output <- Message{
			From: "Бот",
			Text: reply,
		}
	}
}

func readUser(input chan<- Message) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Напиши сообщение:")

	for scanner.Scan() {
		text := scanner.Text()

		input <- Message{
			From: "Вы",
			Text: text,
		}
		if strings.ToLower(text) == "пока" {
			close(input)
			return
		}
	}
}

func printMessages(output <-chan Message) {
	for msg := range output {
		fmt.Printf("[%s]: %s\n", msg.From, msg.Text)
	}
}

func main() {
	UserToBot := make(chan Message)
	BotToUser := make(chan Message)

	go bot(UserToBot, BotToUser)
	go printMessages(BotToUser)

	readUser(UserToBot)
}
