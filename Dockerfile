FROM golang:1.16.0-alpine

ENV GO111MODULE=auto

WORKDIR /go/src/app

COPY . .

RUN go build -o main .

CMD ["/go/src/app/main"]

VOLUME /go/src/app/fasta
