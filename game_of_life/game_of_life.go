package game_of_life

const (
	dead  = 0
	alive = 1
)

type Cell struct {
	x, y           int
	liveliness     int
	new_liveliness int
}

func CheckForOverflow(coord int, len int) int {
	if coord < 0 {
		coord += len
	}
	if coord >= len {
		coord -= len
	}
	return coord
}

func (game *Game) CalculateLivelinessOfCell(cell *Cell) {
	num_alive := 0
	neighbours := make([]Cell, 8)
	off := 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {

			if i == 0 && j == 0 {
				continue
			}

			neighbour_x, neighbour_y := cell.x+j, cell.y+i
			neighbour_x = CheckForOverflow(neighbour_x, game.board_x_len)
			neighbour_y = CheckForOverflow(neighbour_y, game.board_y_len)

			neighbours[off] = game.board_state[neighbour_y][neighbour_x]

			if neighbours[off].liveliness == alive {
				num_alive += 1
			}
			off += 1
		}
	}

	if cell.liveliness == alive {
		// underpopulation
		if num_alive == 0 || num_alive == 1 {
			cell.new_liveliness = dead
			// fine
		} else if num_alive == 2 || num_alive == 3 {
			cell.new_liveliness = alive
			// overpopulation
		} else if num_alive > 3 {
			cell.new_liveliness = dead
		}
	} else if cell.liveliness == dead {
		// reproduction
		if num_alive == 3 {
			cell.new_liveliness = alive
		}
	}
}

type InitialState [][]int

type Game struct {
	board_x_len, board_y_len int
	board_state              [][]Cell
}

func InitGame(init *InitialState) *Game {
	//parse init state from int to cell
	y_len, x_len := len(*init), len((*init)[0])
	cells := make([][]Cell, y_len)
	for i := 0; i < y_len; i++ {
		cells[i] = make([]Cell, x_len)
	}

	for i := 0; i < y_len; i++ {
		for j := 0; j < x_len; j++ {
			cells[i][j].y = i
			cells[i][j].x = j
			cells[i][j].liveliness = (*init)[i][j]
		}
	}

	game := Game{board_state: cells, board_x_len: x_len, board_y_len: y_len}

	return &game
}

func (game *Game) NextLifeCycle() {
	for i := 0; i < game.board_y_len; i++ {
		for j := 0; j < game.board_x_len; j++ {
			game.CalculateLivelinessOfCell(&game.board_state[i][j])
		}
	}

	for i := 0; i < game.board_y_len; i++ {
		for j := 0; j < game.board_x_len; j++ {
			game.board_state[i][j].liveliness = game.board_state[i][j].new_liveliness

		}
	}
}
