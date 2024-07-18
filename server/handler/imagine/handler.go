package imagine

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RivasCVA/fetchmojiai-service/api"
	"github.com/RivasCVA/fetchmojiai-service/client/openai"
	"github.com/RivasCVA/fetchmojiai-service/client/slack"
	"github.com/RivasCVA/fetchmojiai-service/server/response"
)

type ImagineHandlerInterface interface {
	Imagine(w http.ResponseWriter, r *http.Request)
}

type ImagineHandler struct {
	openaiClient openai.OpenAIClientInterface
	slackClient  slack.SlackClientInterface
}

func NewHandler(openaiClient openai.OpenAIClientInterface, slackClient slack.SlackClientInterface) ImagineHandlerInterface {
	return &ImagineHandler{
		openaiClient: openaiClient,
		slackClient:  slackClient,
	}
}

func (h *ImagineHandler) Imagine(w http.ResponseWriter, r *http.Request) {
	// decode the request body
	var body api.ImagineJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(fmt.Errorf("Imagine: unable to decode the request body: %w", err))
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// extract the properties
	event := body.Event
	prompt := h.slackClient.StripUserMentions(event.Text)
	userId := event.User
	timestamp := event.Ts

	// ask openai
	url, err := h.openaiClient.GenerateImage(prompt)
	if err != nil {
		fmt.Println(fmt.Errorf("Imagine: unable to generate the image: %w", err))
		response.WriteError(w, http.StatusInternalServerError, "unable to generate the image")
		return
	}

	// send reply image and message via slack
	err = h.slackClient.ReplyImageWithMessage(userId, timestamp, url, prompt, fmt.Sprintf("<@%s> %s", userId, prompt))
	if err != nil {
		fmt.Println(fmt.Errorf("Imagine: unable to send the image to slack: %w", err))
		response.WriteError(w, http.StatusInternalServerError, "unable to send the image to slack")
		return
	}

	// respond
	out := api.ImagineResponse{
		Image: url,
		User:  userId,
	}
	response.Write(w, http.StatusOK, out)
}
