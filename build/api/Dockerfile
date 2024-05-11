############################
# STEP 1 build executable binary
############################
# Start from golang base image
FROM golang:1.20-alpine3.17 as builder

# Add Maintainer info
LABEL maintainer="Vishal Anarase"

# Set destination for COPY
WORKDIR /go/src/github.com/bookstore

# Copy code
COPY . .

# Download all the dependencies
ENV GO111MODULE=on
RUN go mod download
RUN go mod verify
RUN go mod tidy

# Build
RUN CGO_ENABLE=0 GOOS=linux go build -o /go/bin/api cmd/api/main.go

# Create appuser
ENV USER=api
ENV UID=1000
# See https://stackoverflow.com/a/55757473/13429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
RUN chmod a+x /go/bin/api

############################
# STEP 2 build a small image
############################
FROM alpine:3.17

RUN apk --no-cache add curl

# Copy our static executable
COPY --from=builder /go/bin/api /go/bin/api
COPY --from=builder /etc/passwd /etc/passwd

# Copy the configs directory
COPY internal/config /go/bin/config

# Use an unprivileged user
USER api
WORKDIR /go/bin

EXPOSE 8080/tcp

# Run binary.
ENTRYPOINT ["/go/bin/api"]