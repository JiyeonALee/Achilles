package main

import (
	"context"
	"log"
	"net"
	"sort"
	"time"

	"github.com/backendservice/Achilles/achilles"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
}

func SortScores(score_map map[string][]int) (*achilles.AchillesReply, error) {
	_, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	pl := make(achilles.PairList, len(score_map))
	i := 0
	for k, v := range score_map {
		totalScore := int32(0)
		for _, score := range v {
			totalScore += int32(score)
		}
		temp_p := achilles.EmptyStruct{}
		pl[i] = &achilles.Pair{k, totalScore, temp_p, nil, 0}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	result := &achilles.AchillesReply{Pair: pl}
	return result, nil
}

func (s *server) Compute(context.Context, *achilles.AchillesRequest) (*achilles.AchillesReply, error) {
	x := make(map[string][]int)

	// x["F:Korea"] = append(x["Korea"], 0)
	// x["F:Korea"] = append(x["Korea"], 0)

	// x["F:Mexico"] = append(x["Korea"], 3)
	// x["F:Mexico"] = append(x["Korea"], 3)

	// x["F:Sweden"] = append(x["Korea"], 3)
	// x["F:Sweden"] = append(x["Korea"], 0)

	// x["F:Germany"] = append(x["Korea"], 0)
	// x["F:Germany"] = append(x["Korea"], 3)

	result, err := SortScores(x)
	return result, err
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
