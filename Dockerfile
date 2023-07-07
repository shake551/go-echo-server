FROM golang:1.19-alpine as builder

WORKDIR /app
COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -installsuffix cgo -o main main.go

FROM alpine:latest

RUN apk update \
  && apk add --no-cache bash git curl

WORKDIR /

COPY --from=builder /app/main /api

RUN chmod +x /api

EXPOSE 8080
ENTRYPOINT ["/api"]
