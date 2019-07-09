package config

import (
	"log"
	"os"
)

var (
	// Port to be listened by application
	Port string
	// SlackToken Token
	SlackToken string
	// SlackAppName name
	SlackAppName string
	// AppID id
	AppID string
	// ChannelID id
	ChannelID string
	// DirectMessage string contains User.Id that received all orders
	DirectMessage string
	// MagicWord string
	MagicWord string
	// Suggestion string
	Suggestion string
	// MenuText string
	MenuText string
	// HelloMessage string
	HelloMessage string
	// MenuImageText string
	MenuImageText string
	// MenuImageLink string
	MenuImageLink string
	//SelectMenu constains the json with the menu items
	SelectMenu string
)

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	Port = os.Getenv("SERVER_PORT")
	if Port == "" {
		log.Fatal("variable SERVER_PORT not defined")
	}
	SlackToken = os.Getenv("SLACK_TOKEN")
	if SlackToken == "" {
		log.Fatal("variable SLACK_TOKEN not defined")
	}
	SlackAppName = os.Getenv("SLACK_APP_NAME")
	if SlackAppName == "" {
		log.Fatal("variable SLACK_APP_NAME not defined")
	}
	AppID = os.Getenv("SLACK_APP_ID")
	if AppID == "" {
		log.Fatal("variable SLACK_APP_ID not defined")
	}
	ChannelID = os.Getenv("SLACK_CHANNEL_ID")
	if ChannelID == "" {
		log.Fatal("variable SLACK_CHANNEL_ID not defined")
	}
	DirectMessage = os.Getenv("SLACK_DIRECT_MESSAGE")
	if DirectMessage == "" {
		log.Fatal("variable SLACK_DIRECT_MESSAGE not defined")
	}
	MagicWord = getEnv("SLACK_MAGIC_WORD", "hey")
	Suggestion = getEnv("SLACK_SUGGESTION", "Please, use this")
	MenuText = getEnv("SLACK_MENU_TEXT", "Pick an item from the dropdown list")
	HelloMessage = getEnv("SLACK_HELLO_MESSAGE", "Good Morning, Im ready to receive orders")
	MenuImageText = getEnv("SLACK_IMAGE_TEXT", "My Menu")
	MenuImageLink = getEnv("SLACK_MENU_LINK_IMAGE", "https://d3itj9t5jzykfd.cloudfront.net/ui/451238/image_5be4762c7d4c8.jpg")
	SelectMenu = getEnv("SLACK_SELECT_MENU", "[{\"Text\":\"1x beef\",\"Value\": \"1-beef\"},{\"Text\":\"1x pork\",\"Value\":\"1-pork\"},{\"Text\":\"1x chicken\",\"Value\":\"1-chicken\"},{\"Text\":\"1x Wine\",\"Value\":\"1-wine\"},{\"Text\":\"1x Beer\",\"Value\":\"1-beer\"}]")
}
