# syntax=docker/dockerfile:1

ARG GO_VERSION="1.23-alpine"

FROM golang:${GO_VERSION} AS build

RUN apk add --no-cache --update build-base

ARG TEMPL_VERSION="v0.2.778"
RUN go install github.com/a-h/templ/cmd/templ@${TEMPL_VERSION}

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/mattn/go-sqlite3

COPY . .
RUN templ generate && CGO_ENABLED=1 GOOS=linux go build -o ./build/server ./cmd/server
RUN go test -v ./...

FROM oven/bun:1 AS web

WORKDIR /app

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
RUN bun build:js && bun build:css && bun build:static && bun build:fonts

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=build /app/migrations /app/migrations
COPY --from=build /app/build/server /app/gostart-crm
COPY --from=web /app/build/static /app/static

CMD ["/app/gostart-crm"]
