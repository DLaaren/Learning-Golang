package main

import "dlaaren/game_of_life"

func main() {
	game := game_of_life.InitGame(&game_of_life.InitStateRandom)
	gui := game_of_life.InitGUI(game)

	gui.ShowGame()
}
