package game_of_life

import "math/rand"

var InitStateTetris = InitialState{
	{0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0},
	{0, 0, 1, 1, 1, 0},
	{0, 1, 1, 1, 0, 0},
	{0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0},
}

var InitStateBeacon = InitialState{
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 1, 1, 1, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
}

var InitStateBlinker = InitialState{
	{0, 0, 0, 0, 0, 0},
	{0, 1, 1, 0, 0, 0},
	{0, 1, 1, 0, 0, 0},
	{0, 0, 0, 1, 1, 0},
	{0, 0, 0, 1, 1, 0},
	{0, 0, 0, 0, 0, 0},
}

var InitStateRandom = InitStateGenerateRandom(10, 10)

func InitStateGenerateRandom(x int, y int) InitialState {
	m := make([][]int, y)
	for i := 0; i < y; i++ {
		m[i] = make([]int, x)
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			m[i][j] = rand.Int() % 2
		}
	}
	return m
}
