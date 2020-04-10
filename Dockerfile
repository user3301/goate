
#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o grpclab ./cmd/.

CMD ["./grpclab"]
EXPOSE 8080 8082
