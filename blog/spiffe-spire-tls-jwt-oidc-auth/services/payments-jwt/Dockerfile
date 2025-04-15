FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o jwt-service .

FROM alpine:latest

WORKDIR /data

COPY --from=builder /app/jwt-service .

RUN chmod +x /data/jwt-service

EXPOSE 3000

ENTRYPOINT ["/data/jwt-service"]
