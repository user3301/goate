// protoc --go_out=output_directory input_directory/file.go
//go:generate protoc -I proto/ --go_out=plugins=grpc:pkg/proto proto/grpclab_types.proto
//go:generate protoc -I proto/ -I third_party/ --go_out=plugins=grpc:pkg/proto --grpc-gateway_out=logtostderr=true:pkg/proto ./proto/grpclab_api.proto

package grpclab
