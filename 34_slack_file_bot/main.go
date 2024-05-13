package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/joho/godotenv"
)

func main(){
	// Load SLACK_BOT_TOKEN & CHANNEL_ID in .env file
	godotenv.Load(".env")


	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"ram_gopal_siddh_8209820704.pdf"}

	for i := 0; i<len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil{
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}

}