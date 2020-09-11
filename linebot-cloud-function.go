package linebotcf

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	channelSecret      = "Your channel secret"
	channelAccessToken = "Your channel access token"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

// HTTPHandler is a main entry
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.Header.Get("X-Line-Signature") != "" {
			lineProcess(r)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello~"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func lineProcess(r *http.Request) {
	bot, err := linebot.New(channelSecret, channelAccessToken, linebot.WithHTTPClient(client))
	if err != nil {
		log.Println(err)
		return
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Println(err)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

			switch event.Source.Type {
			case linebot.EventSourceTypeUser:
				fmt.Printf("From user: %s\n", event.Source.UserID)
			case linebot.EventSourceTypeGroup:
				fmt.Printf("From group: %s, userID: %s\n", event.Source.GroupID, event.Source.UserID)
			case linebot.EventSourceTypeRoom:
				fmt.Printf("From room: %s, userID: %s\n", event.Source.RoomID, event.Source.UserID)
			}

			var txtMsg *linebot.TextMessage
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Printf("User message: %s\n", event.Message)
				txtMsg = linebot.NewTextMessage(message.Text + " OK!")
			default:
				txtMsg = linebot.NewTextMessage("ok")
			}

			_, err := bot.ReplyMessage(event.ReplyToken, txtMsg).Do()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// PushMessage pushes the message to the target user or group.
func PushMessage(msg string, to string) {
	bot, err := linebot.New(channelSecret, channelAccessToken, linebot.WithHTTPClient(client))
	if err != nil {
		log.Println(err)
		return
	}

	txtMsg := linebot.NewTextMessage(msg)

	_, err = bot.PushMessage(to, txtMsg).Do()
	if err != nil {
		log.Println(err)
	}
}
