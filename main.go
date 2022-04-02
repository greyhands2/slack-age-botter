package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}
}

func main() {
	//set environment variable
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-316496250898-3323795439380-rQe4D990PtHM1kTVSWh9LF65")

	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A039HNC78BW-3334037830961-c0483d44d6ed3a8f8097ddb1082ca3025b7b4e11c682225e057386c81c5bc156")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	//use a goroutine to print command events

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "Year of birth calculator",
		Example:     "my year of birth is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				fmt.Println("error")
			}
			currentTime := time.Now()
			currentYear := currentTime.Year()
			age := currentYear - yob

			r := fmt.Sprintf("age is %d", age)

			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}
