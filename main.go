package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	. "arieltaiwnalb/models"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				userText := message.Text
				userID := message.ID
				if userText == "CLOCK IN/OUT" {
					//let user clock in/ out
				} else if 0 <= strings.Index(userText, "register") {
					//register
					if hasUser(strings.Fields(userText)[1]) {

					}

				} else {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(userID+":"+userText+" OK!")).Do(); err != nil {
						log.Print(err)
					}
				}

			}

		}
	}
}

func hasUser(lineID string) bool {
	p := md.Person{userLineID: lineID}

	return
}
