FROM golang:1.20.4-bullseye

COPY .devcontainer/cert.cr[t] /usr/local/share/ca-certificates/

RUN update-ca-certificates