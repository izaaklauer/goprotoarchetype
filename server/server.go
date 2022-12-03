package server

import (
	"context"
	"log"

	"github.com/izaaklauer/goprotoarchetype/config"
	goprotoarchetypev1 "github.com/izaaklauer/goprotoarchetype/gen/proto/go/goprotoarchetype/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GoprotoarchetypeServer struct {
	goprotoarchetypev1.UnimplementedGoprotoarchetypeServiceServer

	config config.Goprotoarchetype
}

// NewGoprotoarchetypeServer initializes a new server from config
func NewGoprotoarchetypeServer(config config.Goprotoarchetype) (*GoprotoarchetypeServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := GoprotoarchetypeServer{
		config: config,
	}

	return &server, nil
}

func (s *GoprotoarchetypeServer) HelloWorld(
	ctx context.Context,
	req *goprotoarchetypev1.HelloWorldRequest,
) (*goprotoarchetypev1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &goprotoarchetypev1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
