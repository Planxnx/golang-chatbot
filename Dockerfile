FROM golang:1.12

RUN mkdir -p /go/src/golang-chatbot
WORKDIR /go/src/golang-chatbot

ADD . /go/src/golang-chatbot

RUN go get -v

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8080

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o golang-chatbot" -command="./golang-chatbot"