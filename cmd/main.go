package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	errChan := make(chan error)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	Run(errChan, gatewayServer, pingServer)

	select {
	case err := <-errChan:
		log.Fatalf("fatal error when start the app %v", err)
	case <-ctx.Done():
		log.Print("context cancels")
	case sig := <-sigChan:
		log.Printf("received os signal: %v", sig)
	}
}

func Run(errChan chan<- error, runners ...*http.Server) {
	for _, r := range runners {
		r := r
		go func() {
			errChan <- r.ListenAndServe()
		}()
	}
}
