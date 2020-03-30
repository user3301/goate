//go:generate protoc -I proto/ --go_out=pkg/ proto/grpclab_types.proto
//go:generate protoc -I proto/ --go_out=pkg/ proto/grpclab_api.proto

package grpclab
