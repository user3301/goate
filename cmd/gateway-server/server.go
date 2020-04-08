package gatewayserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/user3301/grpclab/pkg/proto"
)

func NewGatewayServer(ctx context.Context, port int, service pb.GRPCLabAPIServiceServer) (*http.Server, error) {
	grpcHandler := grpc.NewServer()
	pb.RegisterGRPCLabAPIServiceServer(grpcHandler, service)
	reflection.Register(grpcHandler)

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterGRPCLabAPIServiceHandlerServer(ctx, gwmux, service); err != nil {
		return nil, err
	}
	handler := requestDispatcher(ctx, grpcHandler, gwmux)
	log.Println(ctx, "api server running on port %d", port)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}, nil
}

func requestDispatcher(ctx context.Context, grpcHandler, otherHandler http.Handler) http.Handler {
	hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			log.Println(ctx, "dispatch: grpc")
			grpcHandler.ServeHTTP(w, r)
		} else {
			log.Println(ctx, "dispatch: http gateway")
			otherHandler.ServeHTTP(w, r)
		}
	})
	return h2c.NewHandler(hf, &http2.Server{})
}
