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
			var replyMsg string
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				
				userText := message.Text
				userLineID := event.Source.UserID
				p := md.Person{UserLineID:userLineID, UserName:strings.Fields(userText)[1]}
				if !p.HadAUser(){
					replyMsg = "You must to register first."						
				} else if "CLOCK IN/OUT" == userText {
					replyMsg = "CLOCK IN/OUT;"
					//let user clock in/ out
					
					
				} else if 0 <= strings.Index(userText, "register") {
					replyMsg = checkUserStatus(userLineID, strings.Fields(userText)[1]) 
				} else {
					replyMsg = "not register;"
					}
				}

				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg)).Do(); err != nil {
					log.Print(err)
					
			}

		}
	}
}

func checkUserStatus(lineID string, userSetName string) string {
	// log.Println("@hasUser===, userSetName:", userSetName, " lineID:", lineID)
	p := md.Person{UserLineID:lineID, UserName:userSetName}
	// log.Println("p:",p)
	
	if p.HadAUser() { return "You had registered." }
	
	id, err := p.AddAPerson()
	if nil != err { log.Fatal(err) }
	// log.Println("addPerson rs:", id)
	if 0 < id { return "Your registration was successful!" }
	return "Registration Failed"
}

// func clockInNOut(lineID string){
// 	ts := md.TimeSheet{}
// 	ts.Has
// }
