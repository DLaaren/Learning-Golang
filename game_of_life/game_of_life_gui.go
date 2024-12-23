package game_of_life

import "fmt"
import "time"

const (
	clear       = "\033[H\033[2J"
	alive_ascii = "#"
	dead_ascii  = "-"
)

type GUI struct {
	game *Game
}

func InitGUI(game *Game) *GUI {
	gui := GUI{game: game}
	return &gui
}

func (gui *GUI) ShowGame() {
	x_len, y_len := gui.game.board_x_len, gui.game.board_y_len

	board_ascii := make([][]string, y_len)
	for i := 0; i < y_len; i++ {
		board_ascii[i] = make([]string, x_len)
	}

	for {
		board_state := gui.game.board_state
		fmt.Printf("%s", clear)
		fmt.Printf("Game Of Life\n")
		for i := 0; i < y_len; i++ {
			for j := 0; j < x_len; j++ {
				if board_state[i][j].liveliness == alive {
					board_ascii[i][j] = alive_ascii
				} else {
					board_ascii[i][j] = dead_ascii
				}
			}
			fmt.Printf("%s\n", board_ascii[i])
		}
		gui.game.NextLifeCycle()
		time.Sleep(1 * time.Second)
	}
}
