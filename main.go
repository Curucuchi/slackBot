package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/slack-go/slack"
)

type Quote struct {
	Quote string `json:"quote"`
}

func main() {

	hit, err := http.Get("https://api.kanye.rest/")
	if err != nil {
		log.Fatal("There was an issue hitting kanyerest", err)
	}

	read, err := ioutil.ReadAll(hit.Body)
	if err != nil {
		log.Fatal("There was an issue reading the body!", err)
	}

	var q Quote
	json.Unmarshal(read, &q)

	OAUTH_TOKEN := "xoxb-2485850931159-2500608011667-q5VkARnrpZiKyUEeO0AfiER6"
	CHANNEL_ID := "C02EQGWC6HY"

	api := slack.New(OAUTH_TOKEN)
	message := slack.Attachment{
		Text: q.Quote,
	}

	channelId, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText("This is the main message", false),
		slack.MsgOptionAttachments(message),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)
}
