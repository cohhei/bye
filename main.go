package main

import (
	"github.com/cohhei/bye/settings"
	"github.com/nlopes/slack"
)

func main() {
	if err := postToSlack(settings.Token); err != nil {
		panic(err)
	}
}

func postToSlack(token string) error {
	c := slack.New(token)
	_, _, err := c.PostMessage(settings.Channel, settings.Message, slack.PostMessageParameters{})
	return err
}
