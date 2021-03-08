// protoc --go_out=output_directory input_directory/file.go
//go:generate protoc -I proto/ --go_out=plugins=grpc:pkg/proto proto/goate_types.proto
//nolint
//go:generate protoc -I proto/ -I third_party/ --go_out=plugins=grpc:pkg/proto --grpc-gateway_out=logtostderr=true:pkg/proto ./proto/goate_api.proto

package goate
