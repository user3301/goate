
#build stage
FROM golang:alpine AS builder
LABEL maintainer="stan_gai@hotmail.com"
LABEL version="v0.1"
WORKDIR /app
COPY . .
RUN go build -o grpclab ./cmd/.

CMD [ "./grpclab" ]
EXPOSE 80 8082
