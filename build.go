//go:generate protoc --go_out=plugins=grpc:pkg ./proto/entitler_api.proto
//go:generate protoc --go_out=plugins=grpc:pkg ./proto/entitler_types.proto
package gRPCServerClient