package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/user3301/grpclab/internal/service"

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

	store, err := config.GetStore(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	userService := service.NewUserService(store)
	pingServer := &http.Server{
		Addr: fmt.Sprintf(":%d", config.PingServerConfig.Port),
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
			_, err = rw.Write([]byte("received"))
			if err != nil {
				log.Print(err)
			}
		}),
	}

	gatewayServer, err := gatewayserver.NewGatewayServer(ctx, config.AppConfig.Port, userService)
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
