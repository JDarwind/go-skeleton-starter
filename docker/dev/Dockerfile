FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git ca-certificates bash && \
    go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .air.toml ./.air.toml


CMD ["air", "-c", ".air.toml"]