FROM golang:1.16.6-alpine AS dev

COPY main.go /go/src

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS="-mod=vendor"

ENTRYPOINT [ "sh" ]

FROM dev AS build

RUN (([ ! -d "/go/bin/vendor" ] && go mod download && go mod vendor) || true)

RUN go build -ldflags="-s -w" -mod vendor -o /go/bin/fetch /go/src/main.go

RUN chmod +x /go/bin/fetch && mkdir /go/bin/fasta

FROM scratch

COPY --from=build /go/bin /
COPY --from=build etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "/fetch" ]

VOLUME /fasta
