FROM golang:1.23 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /app .

FROM alpine:latest

COPY --from=builder /app /usr/local/bin/app

CMD ["app"]
