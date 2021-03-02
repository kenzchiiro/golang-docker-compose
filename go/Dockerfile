FROM golang:alpine

ENV GO111MODULE=on
WORKDIR /gitlab.com/kenzchiro/bitkub_challange/go/src
#WORKDIR /go/src

RUN apk update && apk add --no-cache git


COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

#RUN go build

CMD ["go", "run", "main.go"]