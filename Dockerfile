FROM golang:alpine AS builder
WORKDIR /app
ADD . /app
RUN cd /app && go build -o users ./cmd/main.go

FROM alpine
WORKDIR /app
COPY app.env .
COPY --from=builder /app/users /app

EXPOSE 8080
CMD ["./users"]
