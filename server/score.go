package main

type nationScore struct {
	Team string
	Nation string
	Score []int32
}

var scores = []nationScore{
	{
		"F",
		"KOR",
		[]int32{0,0},
	},
	{
		"F",
		"MEX",
		[]int32{3,3},
	},	{
		"F",
		"SWE",
		[]int32{3,0},
	},	{
		"F",
		"GER",
		[]int32{0,3},
	},
}