package main

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type sorcererService struct{}

func (s *sorcererService) FiringSolution(ctx context.Context, in *Target) (*Solution, error) {
	log.Printf("sorcerer: received request for a target distance %0.2f m in %0.2f m/s wind", in.Distance, in.Wind)
	return &Solution{Angle: 24.4 /* degrees */, Mass: 10.0 /* kg */}, nil
}

func main() {
	sorcererPort := os.Getenv("SORCERER_PORT")
	if sorcererPort == "" {
		sorcererPort = "5000"
	}

	lis, err := net.Listen("tcp", ":"+sorcererPort)
	if err != nil {
		log.Fatal("sorcerer: Failed to listen on port... error: ", err)
	}

	server := grpc.NewServer()
	RegisterSorcererServiceServer(server, &sorcererService{})
	if err := server.Serve(lis); err != nil {
		log.Fatal(errors.Wrap(err, "sorcerer: Failed to start server!"))
	}
}
