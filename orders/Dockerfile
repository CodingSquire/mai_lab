FROM golang:alpine AS builder

WORKDIR /app

ADD . .
RUN go build -o orders

FROM alpine AS runner

WORKDIR /app
COPY --from=builder /app/orders .

EXPOSE 8080
COPY .env .
ENTRYPOINT ./orders
