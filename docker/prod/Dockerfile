FROM golang:1.25-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./

COPY --exclude="tests" . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/cli ./cmd/cli

FROM alpine:3.20
WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=build /out/api /usr/local/bin/api
COPY --from=build /out/cli /usr/local/bin/cli

ENTRYPOINT ["api"]