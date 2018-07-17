package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	slackToken := os.Getenv("SLACK_TOKEN")
	if err := postToSlack(slackToken); err != nil {
		panic(err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func postToSlack(token string) error {
	c := slack.New(token)
	_, _, err := c.PostMessage(os.Getenv("SLACK_CHANNEL"), os.Getenv("SLACK_MESSAGE"), slack.PostMessageParameters{})
	return err
}
