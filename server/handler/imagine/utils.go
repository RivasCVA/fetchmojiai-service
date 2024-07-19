package imagine

import (
	"fmt"

	"github.com/RivasCVA/fetchmojiai-service/api"
)

const (
	EMOJI_LOADING = ":loading-cat:"
	EMOJI_CHECK   = ":check-check:"
)

// Generates an emoji through OpenAI and replies to the user via Slack.
func processEmojiRequest(h *ImagineHandler, event api.SlackEvent) {
	// extract the properties
	prompt := h.slackClient.StripUserMentions(event.Text)
	userId := event.User
	timestamp := event.Ts

	// acknowledge the user via slack
	err := h.slackClient.ReplyMessage(userId, timestamp, fmt.Sprintf("Generating your emoji %s", EMOJI_LOADING))
	if err != nil {
		fmt.Println(fmt.Errorf("processEmojiRequest: unable to send the acknowledgement to slack: %w", err))
		return
	}

	// ask openai
	// note: curate the prompt to have emoji in its suffix
	url, err := h.openaiClient.GenerateImage(fmt.Sprintf("%s emoji", prompt))
	if err != nil {
		fmt.Println(fmt.Errorf("processEmojiRequest: unable to generate the image on openapi: %w", err))
		return
	}

	// send reply image and message via slack
	err = h.slackClient.ReplyImageWithMessage(userId, timestamp, url, prompt, fmt.Sprintf("%s %s", EMOJI_CHECK, prompt))
	if err != nil {
		fmt.Println(fmt.Errorf("processEmojiRequest: unable to send the image to slack: %w", err))
		return
	}
}
