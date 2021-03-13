
#build stage

FROM golang:1.15-buster AS builder
LABEL maintainer="stan_gai@hotmail.com"
LABEL version="v0.1"
WORKDIR /app
COPY . .
RUN go build -o goate ./cmd

FROM debian:buster-slim

COPY --from=builder /app/goate /bin/goate

RUN addgroup goate && useradd -g goate goate

USER goate

ENTRYPOINT [ "goate" ]
EXPOSE 8080 8082
