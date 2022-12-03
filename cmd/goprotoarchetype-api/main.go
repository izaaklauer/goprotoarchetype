package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/izaaklauer/goprotoarchetype/config"
	goprotoarchetypev1 "github.com/izaaklauer/goprotoarchetype/gen/proto/go/goprotoarchetype/v1"
	"github.com/izaaklauer/goprotoarchetype/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("starting goprotoarchetype.......")
	defer fmt.Println("goprotoarchetype exiting!")

	err := serve()
	if err != nil {
		log.Fatal("failed serving", err)
	}
}

func serve() error {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s config.hcl", os.Args[0])
		os.Exit(1)
	}

	configPath := os.Args[1]

	config, err := config.GetConfig(configPath)
	if err != nil {
		return errors.Wrapf(err, "failed to load config from %q", configPath)
	}

	// Start the service
	goprotoarchetypeServer, err := server.NewGoprotoarchetypeServer(config.Goprotoarchetype)
	if err != nil {
		return errors.Wrapf(err, "failed to start goprotoarchetype server")
	}

	listener, err := net.Listen("tcp", config.Server.BindAddr)
	if err != nil {
		return errors.Wrapf(err, "failed to listen on %s", config.Server.BindAddr)
	}
	grpcServer := grpc.NewServer()

	goprotoarchetypev1.RegisterGoprotoarchetypeServiceServer(grpcServer, goprotoarchetypeServer)
	reflection.Register(grpcServer)

	log.Printf("Serving on %q", config.Server.BindAddr)
	if err := grpcServer.Serve(listener); err != nil {
		return errors.Wrapf(err, "gRPC server exited")
	}

	return nil
}
