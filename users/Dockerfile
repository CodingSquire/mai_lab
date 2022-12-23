FROM golang:alpine AS builder

WORKDIR /app
ADD . /app
RUN cd /app && go build -o users ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/users /app

ENTRYPOINT ./users
