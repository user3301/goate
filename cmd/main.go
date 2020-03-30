package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/user3301/grpclab/pkg/proto"
)

var config = flag.String("config", "", "Config file to load, leave blank to use defaults")

func main() {
	flag.Parse()
	config, err := loadConfig(*config)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = config.Validate()
	if err != nil {
		log.Fatalf("required config is missing %v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.AppConfig.Port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterGRPCLabAPIServiceServer(grpcServer, &proto.UnimplementedGRPCLabAPIServiceServer{})
	// ... determine whether to use TLS
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve grpc server %v", err)
	}
}
