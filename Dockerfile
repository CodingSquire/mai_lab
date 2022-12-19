FROM golang:alpine AS builder


WORKDIR /app
ADD . /app

RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/main.go

FROM alpine:latest

COPY --from=builder /app/users /app

CMD ["./app"]
