package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	pb "github.com/user3301/grpclab/pkg/proto"

	gatewayserver "github.com/user3301/grpclab/cmd/gateway-server"
)

var config = flag.String("config", "", "Config file to load, leave blank to use defaults")

func main() {
	flag.Parse()
	ctx := context.Background()
	config, err := loadConfig(*config)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = config.Validate()
	if err != nil {
		log.Fatalf("required config is missing %v", err)
	}

	pingServer := &http.Server{
		Addr: fmt.Sprintf(":%d", config.PingServerConfig.Port),
		Handler: http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
			log.Println("it's all good baby baby")
		}),
	}

	gatewayServer, err := gatewayserver.NewGatewayServer(ctx, config.AppConfig.Port, &pb.UnimplementedGRPCLabAPIServiceServer{})
	if err != nil {
		log.Fatalf("failed to initialize gateway server %v", err)
	}

	errStream := make(chan error)
	go func() {
		errStream <- gatewayServer.ListenAndServe()
	}()
	go func() {
		errStream <- pingServer.ListenAndServe()
	}()

	if err := <-errStream; err != nil {
		log.Fatalf("fatal error when start servers %v", err)
	}
}
