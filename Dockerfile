FROM golang:1.19.3-alpine

LABEL maintainer="andrew@nijmeh.cloud"
LABEL description="a text adventure game written for my unit project"

WORKDIR /hexa

COPY go.mod ./

COPY . .

RUN go mod download

RUN go build . 
