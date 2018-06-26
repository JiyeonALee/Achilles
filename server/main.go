package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sort"
	"strings"
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

func getScore(nationTxt string) (nation_group string, nation string, nation_scores []int) {
	raw, err := ioutil.ReadFile("./score.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var data WorldCup_Data
	json.Unmarshal(raw, &data)

	nation_id, nation := getNationData(data, nationTxt)
	isValid := false

	nation_group, nation_scores, isValid = getScoreFromGroupA(data, nation_id)

	if !isValid {
		nation_group, nation_scores, isValid = getScoreFromGroupA(data, nation_id)
	}

	if !isValid {
		nation_group, nation_scores, isValid = getScoreFromGroupB(data, nation_id)
	}

	if !isValid {
		nation_group, nation_scores, isValid = getScoreFromGroupC(data, nation_id)
	}

	if !isValid {
		nation_group, nation_scores, isValid = getScoreFromGroupD(data, nation_id)
	}

	if !isValid {
		nation_group, nation_scores, isValid = getScoreFromGroupE(data, nation_id)
	}

	if !isValid {
		nation_group, nation_scores, isValid = getScoreFromGroupF(data, nation_id)
	}

	if !isValid {
		nation_group, nation_scores, isValid = getScoreFromGroupG(data, nation_id)
	}

	if !isValid {
		nation_group, nation_scores, isValid = getScoreFromGroupH(data, nation_id)
	}

	return
}

func getNationData(data WorldCup_Data, nationTxt string) (nation_id int, nation string) {

	for _, team := range data.Teams {
		if strings.Contains(strings.ToLower(team.Name), strings.ToLower(nationTxt)) {
			nation_id = team.ID
			nation = team.Name
			break
		}
	}

	return
}

func getScoreFromGroupA(data WorldCup_Data, nation_id int) (nation_group string, nation_scores []int, isValid bool) {

	for _, match := range data.Groups.A.Matches {
		if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
			nation_group = "A"
			break
		}
	}

	isValid = nation_group != ""
	if !isValid {
		return
	}

	if nation_group != "" {
		for _, match := range data.Groups.A.Matches {
			if match.Finished {
				if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
					isHomeTeam := match.HomeTeam == nation_id

					if match.HomeResult > match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 3)
						} else {
							nation_scores = append(nation_scores, 0)
						}
					} else if match.HomeResult < match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 0)
						} else {
							nation_scores = append(nation_scores, 3)
						}
					} else {
						nation_scores = append(nation_scores, 1)
					}
				}
			}
		}
	}

	return
}

func getScoreFromGroupB(data WorldCup_Data, nation_id int) (nation_group string, nation_scores []int, isValid bool) {

	for _, match := range data.Groups.B.Matches {
		if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
			nation_group = "B"
			break
		}
	}

	isValid = nation_group != ""
	if !isValid {
		return
	}

	if nation_group != "" {
		for _, match := range data.Groups.B.Matches {
			if match.Finished {
				if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
					isHomeTeam := match.HomeTeam == nation_id

					if match.HomeResult > match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 3)
						} else {
							nation_scores = append(nation_scores, 0)
						}
					} else if match.HomeResult < match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 0)
						} else {
							nation_scores = append(nation_scores, 3)
						}
					} else {
						nation_scores = append(nation_scores, 1)
					}
				}
			}
		}
	}

	return
}

func getScoreFromGroupC(data WorldCup_Data, nation_id int) (nation_group string, nation_scores []int, isValid bool) {

	for _, match := range data.Groups.C.Matches {
		if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
			nation_group = "C"
			break
		}
	}

	isValid = nation_group != ""
	if !isValid {
		return
	}

	if nation_group != "" {
		for _, match := range data.Groups.C.Matches {
			if match.Finished {
				if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
					isHomeTeam := match.HomeTeam == nation_id

					if match.HomeResult > match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 3)
						} else {
							nation_scores = append(nation_scores, 0)
						}
					} else if match.HomeResult < match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 0)
						} else {
							nation_scores = append(nation_scores, 3)
						}
					} else {
						nation_scores = append(nation_scores, 1)
					}
				}
			}
		}
	}

	return
}

func getScoreFromGroupD(data WorldCup_Data, nation_id int) (nation_group string, nation_scores []int, isValid bool) {

	for _, match := range data.Groups.D.Matches {
		if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
			nation_group = "D"
			break
		}
	}

	isValid = nation_group != ""
	if !isValid {
		return
	}

	if nation_group != "" {
		for _, match := range data.Groups.D.Matches {
			if match.Finished {
				if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
					isHomeTeam := match.HomeTeam == nation_id

					if match.HomeResult > match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 3)
						} else {
							nation_scores = append(nation_scores, 0)
						}
					} else if match.HomeResult < match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 0)
						} else {
							nation_scores = append(nation_scores, 3)
						}
					} else {
						nation_scores = append(nation_scores, 1)
					}
				}
			}
		}
	}

	return
}

func getScoreFromGroupE(data WorldCup_Data, nation_id int) (nation_group string, nation_scores []int, isValid bool) {

	for _, match := range data.Groups.E.Matches {
		if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
			nation_group = "E"
			break
		}
	}

	isValid = nation_group != ""
	if !isValid {
		return
	}

	if nation_group != "" {
		for _, match := range data.Groups.E.Matches {
			if match.Finished {
				if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
					isHomeTeam := match.HomeTeam == nation_id

					if match.HomeResult > match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 3)
						} else {
							nation_scores = append(nation_scores, 0)
						}
					} else if match.HomeResult < match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 0)
						} else {
							nation_scores = append(nation_scores, 3)
						}
					} else {
						nation_scores = append(nation_scores, 1)
					}
				}
			}
		}
	}

	return
}

func getScoreFromGroupF(data WorldCup_Data, nation_id int) (nation_group string, nation_scores []int, isValid bool) {

	for _, match := range data.Groups.F.Matches {
		if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
			nation_group = "F"
			break
		}
	}

	isValid = nation_group != ""
	if !isValid {
		return
	}

	if nation_group != "" {
		for _, match := range data.Groups.F.Matches {
			if match.Finished {
				if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
					isHomeTeam := match.HomeTeam == nation_id

					if match.HomeResult > match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 3)
						} else {
							nation_scores = append(nation_scores, 0)
						}
					} else if match.HomeResult < match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 0)
						} else {
							nation_scores = append(nation_scores, 3)
						}
					} else {
						nation_scores = append(nation_scores, 1)
					}
				}
			}
		}
	}

	return
}

func getScoreFromGroupG(data WorldCup_Data, nation_id int) (nation_group string, nation_scores []int, isValid bool) {

	for _, match := range data.Groups.G.Matches {
		if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
			nation_group = "G"
			break
		}
	}

	isValid = nation_group != ""
	if !isValid {
		return
	}

	if nation_group != "" {
		for _, match := range data.Groups.G.Matches {
			if match.Finished {
				if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
					isHomeTeam := match.HomeTeam == nation_id

					if match.HomeResult > match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 3)
						} else {
							nation_scores = append(nation_scores, 0)
						}
					} else if match.HomeResult < match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 0)
						} else {
							nation_scores = append(nation_scores, 3)
						}
					} else {
						nation_scores = append(nation_scores, 1)
					}
				}
			}
		}
	}

	return
}

func getScoreFromGroupH(data WorldCup_Data, nation_id int) (nation_group string, nation_scores []int, isValid bool) {

	for _, match := range data.Groups.H.Matches {
		if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
			nation_group = "H"
			break
		}
	}

	isValid = nation_group != ""
	if !isValid {
		return
	}

	if nation_group != "" {
		for _, match := range data.Groups.H.Matches {
			if match.Finished {
				if match.HomeTeam == nation_id || match.AwayTeam == nation_id {
					isHomeTeam := match.HomeTeam == nation_id

					if match.HomeResult > match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 3)
						} else {
							nation_scores = append(nation_scores, 0)
						}
					} else if match.HomeResult < match.AwayResult {
						if isHomeTeam {
							nation_scores = append(nation_scores, 0)
						} else {
							nation_scores = append(nation_scores, 3)
						}
					} else {
						nation_scores = append(nation_scores, 1)
					}
				}
			}
		}
	}

	return
}

func main() {

	nation_group, nation, nation_scores := getScore("korea")

	fmt.Println(nation_group)
	fmt.Println(nation)
	for _, score := range nation_scores {
		fmt.Println(score)
	}

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
