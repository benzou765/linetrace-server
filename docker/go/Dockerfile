FROM golang:1.12-alpine

RUN apk add --no-cache git
RUN go get github.com/labstack/echo/...
RUN go get github.com/labstack/echo-contrib/session
RUN go get github.com/goml/gobrain

WORKDIR /app 
