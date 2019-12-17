# gRPC Server-Client with Protoc-gen-go
1. install `protoc` by download the binary file from website and place the file under your PATH
2. install gRPC by running `go get -g google.golang.org/grpc`
3. define http request in `.proto` file under `proto` directory
4. place `google/protobuf` under `vendor` directory
5. once the `.proto` file defines the http rules, use protoc by running `protoc --proto_path=<proto directory> --proto_path=<protobuf directory> --go_out=plugines=grpc:proto <your .proto file name>`,
for this particular instance, the command is `protoc --proto_path=proto --proto+path=vendor --go_out=plugins=grpc:proto service.proto`
