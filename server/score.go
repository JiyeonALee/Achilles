package main

import "time"

// type nationScore struct {
// 	Group  string
// 	Nation string
// 	Score  []int32
// }

// var scores = []nationScore{
// 	{
// 		"F",
// 		"KOR",
// 		[]int32{0, 0},
// 	},
// 	{
// 		"F",
// 		"MEX",
// 		[]int32{3, 3},
// 	}, {
// 		"F",
// 		"SWE",
// 		[]int32{3, 0},
// 	}, {
// 		"F",
// 		"GER",
// 		[]int32{0, 3},
// 	},
// }

type WorldCup_Data struct {
	Stadiums []struct {
		ID    int     `json:"id"`
		Name  string  `json:"name"`
		City  string  `json:"city"`
		Lat   float64 `json:"lat"`
		Lng   float64 `json:"lng"`
		Image string  `json:"image"`
	} `json:"stadiums"`
	Tvchannels []struct {
		ID      int      `json:"id"`
		Name    string   `json:"name"`
		Icon    string   `json:"icon"`
		Country string   `json:"country"`
		Iso2    string   `json:"iso2"`
		Lang    []string `json:"lang"`
	} `json:"tvchannels"`
	Teams []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		FifaCode    string `json:"fifaCode"`
		Iso2        string `json:"iso2"`
		Flag        string `json:"flag"`
		Emoji       string `json:"emoji"`
		EmojiString string `json:"emojiString"`
	} `json:"teams"`
	Groups struct {
		A struct {
			Name     string `json:"name"`
			Winner   int    `json:"winner"`
			Runnerup int    `json:"runnerup"`
			Matches  []struct {
				Name       int       `json:"name"`
				Type       string    `json:"type"`
				HomeTeam   int       `json:"home_team"`
				AwayTeam   int       `json:"away_team"`
				HomeResult int       `json:"home_result"`
				AwayResult int       `json:"away_result"`
				Date       time.Time `json:"date"`
				Stadium    int       `json:"stadium"`
				Channels   []int     `json:"channels"`
				Finished   bool      `json:"finished"`
				Matchday   int       `json:"matchday"`
			} `json:"matches"`
		} `json:"a"`
		B struct {
			Name     string `json:"name"`
			Winner   int    `json:"winner"`
			Runnerup int    `json:"runnerup"`
			Matches  []struct {
				Name       int       `json:"name"`
				Type       string    `json:"type"`
				HomeTeam   int       `json:"home_team"`
				AwayTeam   int       `json:"away_team"`
				HomeResult int       `json:"home_result"`
				AwayResult int       `json:"away_result"`
				Date       time.Time `json:"date"`
				Stadium    int       `json:"stadium"`
				Channels   []int     `json:"channels"`
				Finished   bool      `json:"finished"`
				Matchday   int       `json:"matchday"`
			} `json:"matches"`
		} `json:"b"`
		C struct {
			Name     string      `json:"name"`
			Winner   interface{} `json:"winner"`
			Runnerup interface{} `json:"runnerup"`
			Matches  []struct {
				Name       int       `json:"name"`
				Type       string    `json:"type"`
				HomeTeam   int       `json:"home_team"`
				AwayTeam   int       `json:"away_team"`
				HomeResult int       `json:"home_result"`
				AwayResult int       `json:"away_result"`
				Date       time.Time `json:"date"`
				Stadium    int       `json:"stadium"`
				Channels   []int     `json:"channels"`
				Finished   bool      `json:"finished"`
				Matchday   int       `json:"matchday"`
			} `json:"matches"`
		} `json:"c"`
		D struct {
			Name     string      `json:"name"`
			Winner   interface{} `json:"winner"`
			Runnerup interface{} `json:"runnerup"`
			Matches  []struct {
				Name       int       `json:"name"`
				Type       string    `json:"type"`
				HomeTeam   int       `json:"home_team"`
				AwayTeam   int       `json:"away_team"`
				HomeResult int       `json:"home_result"`
				AwayResult int       `json:"away_result"`
				Date       time.Time `json:"date"`
				Stadium    int       `json:"stadium"`
				Channels   []int     `json:"channels"`
				Finished   bool      `json:"finished"`
				Matchday   int       `json:"matchday"`
			} `json:"matches"`
		} `json:"d"`
		E struct {
			Name     string      `json:"name"`
			Winner   interface{} `json:"winner"`
			Runnerup interface{} `json:"runnerup"`
			Matches  []struct {
				Name       int       `json:"name"`
				Type       string    `json:"type"`
				HomeTeam   int       `json:"home_team"`
				AwayTeam   int       `json:"away_team"`
				HomeResult int       `json:"home_result"`
				AwayResult int       `json:"away_result"`
				Date       time.Time `json:"date"`
				Stadium    int       `json:"stadium"`
				Channels   []int     `json:"channels"`
				Finished   bool      `json:"finished"`
				Matchday   int       `json:"matchday"`
			} `json:"matches"`
		} `json:"e"`
		F struct {
			Name     string      `json:"name"`
			Winner   interface{} `json:"winner"`
			Runnerup interface{} `json:"runnerup"`
			Matches  []struct {
				Name       int       `json:"name"`
				Type       string    `json:"type"`
				HomeTeam   int       `json:"home_team"`
				AwayTeam   int       `json:"away_team"`
				HomeResult int       `json:"home_result"`
				AwayResult int       `json:"away_result"`
				Date       time.Time `json:"date"`
				Stadium    int       `json:"stadium"`
				Channels   []int     `json:"channels"`
				Finished   bool      `json:"finished"`
				Matchday   int       `json:"matchday"`
			} `json:"matches"`
		} `json:"f"`
		G struct {
			Name     string      `json:"name"`
			Winner   interface{} `json:"winner"`
			Runnerup interface{} `json:"runnerup"`
			Matches  []struct {
				Name       int       `json:"name"`
				Type       string    `json:"type"`
				HomeTeam   int       `json:"home_team"`
				AwayTeam   int       `json:"away_team"`
				HomeResult int       `json:"home_result"`
				AwayResult int       `json:"away_result"`
				Date       time.Time `json:"date"`
				Stadium    int       `json:"stadium"`
				Channels   []int     `json:"channels"`
				Finished   bool      `json:"finished"`
				Matchday   int       `json:"matchday"`
			} `json:"matches"`
		} `json:"g"`
		H struct {
			Name     string      `json:"name"`
			Winner   interface{} `json:"winner"`
			Runnerup interface{} `json:"runnerup"`
			Matches  []struct {
				Name       int       `json:"name"`
				Type       string    `json:"type"`
				HomeTeam   int       `json:"home_team"`
				AwayTeam   int       `json:"away_team"`
				HomeResult int       `json:"home_result"`
				AwayResult int       `json:"away_result"`
				Date       time.Time `json:"date"`
				Stadium    int       `json:"stadium"`
				Channels   []int     `json:"channels"`
				Finished   bool      `json:"finished"`
				Matchday   int       `json:"matchday"`
			} `json:"matches"`
		} `json:"h"`
	} `json:"groups"`
	Knockout struct {
		Round16 struct {
			Name    string `json:"name"`
			Matches []struct {
				Name        int         `json:"name"`
				Type        string      `json:"type"`
				HomeTeam    int         `json:"home_team"`
				AwayTeam    int         `json:"away_team"`
				HomeResult  interface{} `json:"home_result"`
				AwayResult  interface{} `json:"away_result"`
				HomePenalty interface{} `json:"home_penalty"`
				AwayPenalty interface{} `json:"away_penalty"`
				Winner      interface{} `json:"winner"`
				Date        time.Time   `json:"date"`
				Stadium     int         `json:"stadium"`
				Channels    []int       `json:"channels"`
				Finished    bool        `json:"finished"`
				Matchday    int         `json:"matchday"`
			} `json:"matches"`
		} `json:"round_16"`
		Round8 struct {
			Name    string `json:"name"`
			Matches []struct {
				Name        int         `json:"name"`
				Type        string      `json:"type"`
				HomeTeam    int         `json:"home_team"`
				AwayTeam    int         `json:"away_team"`
				HomeResult  interface{} `json:"home_result"`
				AwayResult  interface{} `json:"away_result"`
				HomePenalty interface{} `json:"home_penalty"`
				AwayPenalty interface{} `json:"away_penalty"`
				Winner      interface{} `json:"winner"`
				Date        time.Time   `json:"date"`
				Stadium     int         `json:"stadium"`
				Channels    []int       `json:"channels"`
				Finished    bool        `json:"finished"`
				Matchday    int         `json:"matchday"`
			} `json:"matches"`
		} `json:"round_8"`
		Round4 struct {
			Name    string `json:"name"`
			Matches []struct {
				Name        int         `json:"name"`
				Type        string      `json:"type"`
				HomeTeam    int         `json:"home_team"`
				AwayTeam    int         `json:"away_team"`
				HomeResult  interface{} `json:"home_result"`
				AwayResult  interface{} `json:"away_result"`
				HomePenalty interface{} `json:"home_penalty"`
				AwayPenalty interface{} `json:"away_penalty"`
				Winner      interface{} `json:"winner"`
				Date        time.Time   `json:"date"`
				Stadium     int         `json:"stadium"`
				Channels    []int       `json:"channels"`
				Finished    bool        `json:"finished"`
				Matchday    int         `json:"matchday"`
			} `json:"matches"`
		} `json:"round_4"`
		Round2Loser struct {
			Name    string `json:"name"`
			Matches []struct {
				Name        int         `json:"name"`
				Type        string      `json:"type"`
				HomeTeam    int         `json:"home_team"`
				AwayTeam    int         `json:"away_team"`
				HomeResult  interface{} `json:"home_result"`
				AwayResult  interface{} `json:"away_result"`
				HomePenalty interface{} `json:"home_penalty"`
				AwayPenalty interface{} `json:"away_penalty"`
				Winner      interface{} `json:"winner"`
				Date        time.Time   `json:"date"`
				Stadium     int         `json:"stadium"`
				Channels    []int       `json:"channels"`
				Finished    bool        `json:"finished"`
				Matchday    int         `json:"matchday"`
			} `json:"matches"`
		} `json:"round_2_loser"`
		Round2 struct {
			Name    string `json:"name"`
			Matches []struct {
				Name        int         `json:"name"`
				Type        string      `json:"type"`
				HomeTeam    int         `json:"home_team"`
				AwayTeam    int         `json:"away_team"`
				HomeResult  interface{} `json:"home_result"`
				AwayResult  interface{} `json:"away_result"`
				HomePenalty interface{} `json:"home_penalty"`
				AwayPenalty interface{} `json:"away_penalty"`
				Winner      interface{} `json:"winner"`
				Date        time.Time   `json:"date"`
				Stadium     int         `json:"stadium"`
				Channels    []int       `json:"channels"`
				Finished    bool        `json:"finished"`
				Matchday    int         `json:"matchday"`
			} `json:"matches"`
		} `json:"round_2"`
	} `json:"knockout"`
}
