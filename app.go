package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

func MessengerVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		challenge := r.URL.Query().Get("hub.challenge")
		verify_token := r.URL.Query().Get("hub.verify_token")
		if len(verify_token) > 0 && len(challenge) > 0 && verify_token == "planx-golang" {
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(w, challenge)
			return
		}
	} else if r.Method == "POST" {
		defer r.Body.Close()
		input := new(MessengerInput)
		if err := json.NewDecoder(r.Body).Decode(input); err == nil {
			log.Println("got message:", input.Entry[0].Messaging[0].Message.Text)
			log.Println("INPUT :", input[0])
			reply := input.Entry[0].Messaging[0]
			reply.Sender, reply.Recipient = reply.Recipient, reply.Sender

			reply.Message.Text = input.Entry[0].Messaging[0].Message.Text
			reply.Message.Seq = 0
			reply.Message.Mid = ""

			b, _ := json.Marshal(reply)
			http.Post("https://graph.facebook.com/v2.6/me/messages?access_token=EAAf5m6IkbZAUBAGfEZALvFtyRj8ahcT0kE9FTKiCz7owzkdwthWDabMHrfhZCEhouIj9pzbYJBnbbWzhZBycU8QVLGyBoxqzflv7AgLQLAD7m0zSCtLdjz1nWdB1JZAiFYMzbpPZCfdZAoUAYKgHGhizfD8j71cgTrlaOdj4EYBZBtHeZBQZCtC6XU",
				"application/json",
				bytes.NewReader(b))
			return
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(400)
	fmt.Fprintf(w, "Bad Request")
}
