package imagine

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

	// only reply to messages that have a mention
	// note: slack also pings the endpoint when the bot sends messages
	if strings.Contains(body.Event.Text, "@") {
		// asynchronously generate the image and reply to the user
		go generateAndReply(h, body.Event)
	}

	// respond
	out := api.ImagineResponse{Accepted: true}
	response.Write(w, http.StatusOK, out)
}
