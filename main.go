package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var developMessegeCounter = 0

type MessengerInput struct {
	Entry []struct {
		Time      uint64 `json:"time,omitempty"`
		Messaging []struct {
			Sender struct {
				Id string `json:"id"`
			} `json:"sender,omitempty"`
			Recipient struct {
				Id string `json:"id"`
			} `json:"recipient,omitempty"`
			Timestamp uint64 `json:"timestamp,omitempty"`
			Message   *struct {
				Mid  string `json:"mid,omitempty"`
				Seq  uint64 `json:"seq,omitempty"`
				Text string `json:"text"`
			} `json:"message,omitempty"`
		} `json:"messaging"`
	}
}

func main() {
	http.HandleFunc("/webhook", MessengerVerify)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MessengerVerify(res http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		challenge := req.URL.Query().Get("hub.challenge")
		verify_token := req.URL.Query().Get("hub.verify_token")
		if len(verify_token) > 0 && len(challenge) > 0 && verify_token == "planx-golang" {
			res.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(res, challenge)
			return
		}
	} else if req.Method == "POST" {
		defer req.Body.Close()
		input := new(MessengerInput)
		if err := json.NewDecoder(req.Body).Decode(input); err == nil {

			log.Println("Sender ID :", input.Entry[0].Messaging[0].Sender.Id)
			log.Println("Message :", input.Entry[0].Messaging[0].Message.Text)

			reply := input.Entry[0].Messaging[0]
			reply.Sender, reply.Recipient = reply.Recipient, reply.Sender
			if input.Entry[0].Messaging[0].Sender.Id == strconv.Itoa(2222082947888679) {
				developMessegeCounter++
				reply.Message.Text = "(" + strconv.Itoa(developMessegeCounter) + ")" + input.Entry[0].Messaging[0].Message.Text
			} else {
				reply.Message.Text = input.Entry[0].Messaging[0].Message.Text
			}
			reply.Message.Seq = 0
			reply.Message.Mid = ""

			b, _ := json.Marshal(reply)
			http.Post("https://graph.facebook.com/v2.6/me/messages?access_token=EAAf5m6IkbZAUBAGfEZALvFtyRj8ahcT0kE9FTKiCz7owzkdwthWDabMHrfhZCEhouIj9pzbYJBnbbWzhZBycU8QVLGyBoxqzflv7AgLQLAD7m0zSCtLdjz1nWdB1JZAiFYMzbpPZCfdZAoUAYKgHGhizfD8j71cgTrlaOdj4EYBZBtHeZBQZCtC6XU",
				"application/json",
				bytes.NewReader(b))
			return
		}
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(400)
	fmt.Fprintf(res, "Bad Request")
}
