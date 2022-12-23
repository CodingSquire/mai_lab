FROM golang:alpine AS builder

WORKDIR /app

ADD . .
RUN go build -mod=readonly -o migrations ./migrations.go

FROM alpine AS runner

WORKDIR /app
COPY --from=builder /app/migrations .

ENTRYPOINT ["./migrations"]
