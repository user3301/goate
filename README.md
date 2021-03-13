# goate

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9c92631cb282462fb2fa5737516ad851)](https://app.codacy.com/manual/user3301/goate?utm_source=github.com&utm_medium=referral&utm_content=user3301/goate&utm_campaign=Badge_Grade_Dashboard)

An API server serves both HTTP and gRPC services on same port.

## Prerequisites

### Required

* __Go__ (>= 1.15)
* __protobuf__
* __gRPC__
  
### Optional

* __grpcurl__
* __GNU Make__
* __Docker__

## Getting Started

### Using **Docker**

You can build image on your local machine using the `Dockerfile` by running:

```shell
docker built -t [image_name] .
```

or if you have GNU Make you can simply run:

```make
make image
```

or pull the image from Docker Hub:

```shell
docker pull user3301/goate:latest
```
