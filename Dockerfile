FROM golang:1.20.1-alpine


RUN apk update && apk add git

WORKDIR /go/src/app
COPY ./app /go/src/app

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]