package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	md "github.com/yoshiW86/ArielTaiwanLB/models"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var debugMsg string


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
			replyMsg := ""
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				
				userText := message.Text
				userID := event.Source.UserID
				if userText == "CLOCK IN/OUT" {
					replyMsg += "CLOCK IN/OUT;"
					//let user clock in/ out
				} else if 0 <= strings.Index(userText, "register") {
					//register
					// replyMsg += "register userID=" + userID + " " + strings.Fields(userText)[1]
					replyMsg += hasUser(userID, strings.Fields(userText)[1]) 
				} else {
					replyMsg += "not register;"
					}
				}

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg)).Do(); err != nil {
					log.Print(err)
					
			}

		}
	}
}

func hasUser(lineID string, userSetName string) string {
	log.Println("@hasUser===, userSetName:", userSetName, " lineID:", lineID)
	p := md.Person{UserLineID:lineID, UserName:userSetName}
	log.Println("p:",p)
	ra, err := p.GetPersonsByLID()
	if nil != err {
		log.Fatal(err)
	}
	log.Println("ra[0]:", ra)
	if nil != ra {
		return "You are registered."
	}
	
	id, err := p.AddPerson()
	if nil != err {
		log.Fatal(err)
	}
	log.Println("addPerson rs:", id)
	if 0 < id {
		return "success registration."
	}
	return "Registration Failed"

	
}
