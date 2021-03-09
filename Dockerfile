
#build stage

FROM golang:1.15-buster AS builder
LABEL maintainer="stan_gai@hotmail.com"
LABEL version="v0.1"
WORKDIR /app
COPY . .
RUN go build -o goate ./cmd

ENTRYPOINT [ "/app/goate" ]
EXPOSE 8080 8082
