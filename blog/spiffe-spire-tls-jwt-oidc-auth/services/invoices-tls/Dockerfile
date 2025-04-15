FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /data

COPY --from=builder /app/app .

RUN chmod +x /data/app

EXPOSE 3000

ENTRYPOINT ["/data/app"]
