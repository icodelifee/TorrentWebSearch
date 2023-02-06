FROM golang:1.18.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN apk add git bash

RUN go mod download

COPY . .

RUN go build -o itorrent-gofiber

EXPOSE 5050

RUN chmod +x start.sh
CMD ["bash", "start.sh"]

