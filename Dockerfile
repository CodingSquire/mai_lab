FROM golang:alpine AS b
WORKDIR /app
ADD . /app
RUN cd /app && go build -o users ./cmd/main.go

FROM alpine
WORKDIR /a
COPY --from=b /app/users /app

EXPOSE 8080
ENTRYPOINT .
