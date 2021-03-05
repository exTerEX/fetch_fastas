FROM golang:1.16.0-alpine3.13 AS build

WORKDIR /go/src/app/

COPY main.go /go/src/app/main.go

RUN go build -o /bin/app /go/src/app/main.go

FROM alpine:3.13

COPY --from=build /bin/app /bin/app

CMD ["/bin/app"]

VOLUME /bin/fasta
