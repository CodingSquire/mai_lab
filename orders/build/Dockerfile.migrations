FROM golang:alpine AS builder

WORKDIR /app

ADD . .
RUN go build -o migrations ./tools/migrations/migrations.go

FROM alpine AS runner

WORKDIR /app
COPY --from=builder /app/migrations .

ENTRYPOINT ["./migrations"]
