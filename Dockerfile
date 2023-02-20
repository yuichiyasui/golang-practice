FROM golang:1.20.1-alpine

WORKDIR /go/src/app

RUN apk update && apk add git

COPY ./app .

CMD ["go", "run", "main.go" ]