FROM golang:alpine AS builder

WORKDIR /app

ADD . .
RUN go build -o orders ./cmd/main.go

FROM alpine AS runner

WORKDIR /app
COPY --from=builder /app/orders .

ENTRYPOINT ["./orders"]
