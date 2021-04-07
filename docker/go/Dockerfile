FROM golang:latest

ENV GO111MODULE=on

COPY src/ /go/src/app/

WORKDIR /go/src/app

# cosmtrek/airのインストール
RUN go get -u github.com/cosmtrek/air

RUN go mod download

CMD air -c .air.toml