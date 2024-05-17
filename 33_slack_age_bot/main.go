package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
	"github.com/joho/godotenv"
	//"go.mongodb.org/mongo-driver/event"
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
	// Load SLACK_BOT_TOKEN & SLACK_APP_TOKEN in .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		//Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Println("Error converting year to integer:", err)
				err := response.Reply("Invalid year format. Please provide a valid year.")
				if err != nil {
					log.Println("Error replying to message:", err)
				}
				return
			}
			age := 2024 - yob
			r := fmt.Sprintf("Age is %d", age)
			err = response.Reply(r)
			if err != nil {
				log.Println("Error replying to message:", err)
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal("Slacker bot listen error:", err)
	}
}
