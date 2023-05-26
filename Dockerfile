FROM golang:1.20.4-bullseye as builder

WORKDIR /

COPY cmd/ cmd/
COPY go.mod go.mod
COPY go.sum go.sum
