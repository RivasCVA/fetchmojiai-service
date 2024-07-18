package slack

import (
	"fmt"
	"os"

	slackExternal "github.com/slack-go/slack"
)

type SlackClientInterface interface {
	// Sends a message to a given user.
	SendMessage(userId string, message string) error
	// Sends an image to a given user.
	SendImage(userId string, image string, alt string) error
}

type SlackClient struct {
	slack *slackExternal.Client
}

func NewClient() SlackClientInterface {
	// create a new external slack client
	token := os.Getenv("SLACK_TOKEN")
	slack := slackExternal.New(token)

	return &SlackClient{
		slack: slack,
	}
}

func (c *SlackClient) SendMessage(userId string, message string) error {
	// post the message on slack
	_, _, err := c.slack.PostMessage(
		userId,
		slackExternal.MsgOptionText(message, false),
		slackExternal.MsgOptionAsUser(true),
	)
	if err != nil {
		return fmt.Errorf("sendMessage: unable to send a message to user %s: %w", userId, err)
	}

	return nil
}

func (c *SlackClient) SendImage(userId string, image string, alt string) error {
	// post the image on slack
	_, _, err := c.slack.PostMessage(
		userId,
		slackExternal.MsgOptionAttachments(slackExternal.Attachment{ImageURL: image, Fallback: alt}),
		slackExternal.MsgOptionAsUser(true),
	)
	if err != nil {
		return fmt.Errorf("SendImage: unable to send an image to user %s: %w", userId, err)
	}

	return nil
}
