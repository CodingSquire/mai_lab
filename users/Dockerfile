FROM golang:alpine AS builder
WORKDIR /app
ADD . /app
RUN cd /app && go build -o users .

FROM alpine
WORKDIR /app
COPY --from=builder /app/users /app

EXPOSE 80
ENV PORT_MAIN=80
ENTRYPOINT ./users
