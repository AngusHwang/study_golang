FROM golang:1.18-alpine

ENV TZ=Asia/Seoul

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod verify

COPY . .

RUN go build -o main

EXPOSE 3000

CMD ["./main"]