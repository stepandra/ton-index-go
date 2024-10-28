FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .
RUN go build -o ton-index-go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/ton-index-go .
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 4100
ENTRYPOINT ["./entrypoint.sh"]
