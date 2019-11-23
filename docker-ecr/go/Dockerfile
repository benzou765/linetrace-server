FROM golang:1.12-alpine

RUN apk add --no-cache git
RUN go get github.com/labstack/echo/...
RUN go get github.com/labstack/echo-contrib/session

ADD src.tar.gz /app
WORKDIR /app
CMD ["go","run","main.go"]
