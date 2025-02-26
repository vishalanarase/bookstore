FROM golang:1.21-alpine3.17 AS builder
LABEL maintainer="Vishal Anarase"

WORKDIR /go/src/github.com/bookstore

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/api cmd/api/main.go

ENV USER=api UID=1000
RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid "${UID}" "${USER}" && \
    chmod a+x /go/bin/api

FROM alpine:3.17
RUN apk --no-cache add curl

COPY --from=builder /go/bin/api /go/bin/api
COPY --from=builder /etc/passwd /etc/passwd 

USER api
WORKDIR /go/bin

EXPOSE 8080/tcp
ENTRYPOINT ["/go/bin/api"]