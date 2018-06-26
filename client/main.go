package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/backendservice/Achilles/achilles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	//	address     = "ec2-18-191-204-27.us-east-2.compute.amazonaws.com:50051"
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewComputingMinScoreClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	reply, err := Compute(c, context.Background(), time.Minute, name)

	for _, v := range reply.GetPair() {
		group_team := strings.Split(v.Key, ":")
		fmt.Printf("Group %s, Team %s: %d", group_team[0], group_team[1], v.Value)
		fmt.Println()
	}

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func Compute(c pb.ComputingMinScoreClient, ctx context.Context, timeout time.Duration, name string) (*pb.AchillesReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	reply, err := c.Compute(ctx, &pb.AchillesRequest{Name: ""})
	if err != nil {
		return nil, err
	}

	return reply, nil
}
