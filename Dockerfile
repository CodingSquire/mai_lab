FROM golang:alpine AS builder
WORKDIR /app
ADD . /app
RUN cd /app && go build -o gsl ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/gsl /app

EXPOSE 8000

CMD ["./gsl"]