package main

import (
	"github.com/backendservice/Achilles/achilles"
	grpc "google.golang.org/grpc"
	"context"
	"log"
	"net"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	
}

func (s *server) Compute(context.Context, *achilles.AchillesRequest) (*achilles.AchillesReply, error) {
	result := &achilles.AchillesReply {Score: "3"}
	return result, nil
}

func getScore() []nationScore{
	return scores
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	achilles.RegisterComputingMinScoreServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

