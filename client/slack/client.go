package slack

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	slackExternal "github.com/slack-go/slack"
)

type SlackClientInterface interface {
	// Sends a message to a channel. A user id can be used for direct messages.
	SendMessage(channelId string, message string) error
	// Sends an image to a channel. A user id can be used for direct messages.
	SendImage(channelId string, image string, alt string) error
	// Replies a message as a thread in a channel. A user id can be used for direct messages.
	ReplyMessage(channelId string, timestamp string, message string) error
	// Replies an image with a message as a thread in a channel. A user id can be used for direct messages.
	ReplyImageWithMessage(channelId string, timestamp string, image string, alt string, message string) error
	// Removes any user mentions from the message.
	StripUserMentions(message string) string
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

func (c *SlackClient) SendMessage(channelId string, message string) error {
	// post the message on slack
	_, _, err := c.slack.PostMessage(
		channelId,
		slackExternal.MsgOptionText(message, false),
		slackExternal.MsgOptionAsUser(true),
	)
	if err != nil {
		return fmt.Errorf("SendMessage: unable to send a message to channel %s: %w", channelId, err)
	}

	return nil
}

func (c *SlackClient) SendImage(channelId string, image string, alt string) error {
	// post the image on slack
	_, _, err := c.slack.PostMessage(
		channelId,
		slackExternal.MsgOptionAttachments(slackExternal.Attachment{ImageURL: image, Fallback: alt}),
		slackExternal.MsgOptionAsUser(true),
	)
	if err != nil {
		return fmt.Errorf("SendImage: unable to send an image to channel %s: %w", channelId, err)
	}

	return nil
}

func (c *SlackClient) ReplyMessage(channelId string, timestamp string, message string) error {
	// post the reply message on slack
	_, _, err := c.slack.PostMessage(
		channelId,
		slackExternal.MsgOptionTS(timestamp),
		slackExternal.MsgOptionText(message, false),
		slackExternal.MsgOptionAsUser(true),
	)
	if err != nil {
		return fmt.Errorf("ReplyMessage: unable to reply to channel %s with timestamp %s: %w", channelId, timestamp, err)
	}

	return nil
}

func (c *SlackClient) ReplyImageWithMessage(channelId string, timestamp string, image string, alt string, message string) error {
	// post the reply image and message on slack
	_, _, err := c.slack.PostMessage(
		channelId,
		slackExternal.MsgOptionTS(timestamp),
		slackExternal.MsgOptionAttachments(slackExternal.Attachment{ImageURL: image, Fallback: alt}),
		slackExternal.MsgOptionText(message, false),
		slackExternal.MsgOptionAsUser(true),
	)
	if err != nil {
		return fmt.Errorf("ReplyImageWithMessage: unable to reply to channel %s with timestamp %s: %w", channelId, timestamp, err)
	}

	return nil
}

func (c *SlackClient) StripUserMentions(message string) string {
	regex := regexp.MustCompile("<@[^>]+>")
	return strings.TrimSpace(regex.ReplaceAllString(message, ""))
}
