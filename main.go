package main

import (
	server "awesome-grpc/protos"
	"awesome-grpc/protos/protos"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	log := hclog.Default()
	gRPCServer := grpc.NewServer()
	currencyServer := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gRPCServer, currencyServer)
	reflection.Register(gRPCServer)
	listener, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to start", err)
		os.Exit(1)
	}
	log.Info("Starting listening on ", "Address ", listener.Addr())
	gRPCServer.Serve(listener)
}
