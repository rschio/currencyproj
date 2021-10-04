package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rschio/currencyproj/currency/cmd/util"
	gw "github.com/rschio/currencyproj/currency/currency/proto"
	"google.golang.org/grpc"
)

var (
	serverPort  = util.MustEnv("PORT")
	gatewayPort = util.MustEnv("GATEWAY_PORT")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterExchangerHandlerFromEndpoint(ctx, mux, "localhost:"+serverPort, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":"+gatewayPort, mux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
