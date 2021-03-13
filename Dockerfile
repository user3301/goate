
#build stage

FROM golang:1.15-buster AS builder
LABEL maintainer="stan_gai@hotmail.com"
LABEL version="v0.1"

# create directory permission denied issue resolution
# https://code.visualstudio.com/docs/remote/containers-advanced#_adding-a-nonroot-user-to-your-dev-container
ARG USERNAME=goate
ARG USER_UID=1000
ARG USER_GID=${USER_UID}

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME
# -----------END OF RESOLUTION--------

WORKDIR /app

COPY . .
RUN go build -o goate ./cmd

FROM debian:buster-slim

COPY --from=builder /app/goate /bin/goate

USER ${USERNAME}

ENTRYPOINT [ "goate" ]
EXPOSE 8080 8082
