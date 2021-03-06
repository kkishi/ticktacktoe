package board

import (
	"fmt"

	"github.com/kkishi/ticktacktoe/model/player"
)

type BoardObserver interface {
	NotifyUpdate(row, col int, player player.Player) error
}

type Board struct {
	Grid      [3][3]player.Player
	observers []BoardObserver
}

func (b *Board) CanTake(row, col int) bool {
	return 0 <= row && row <= 2 && 0 <= col && col <= 2 &&
		b.Grid[row][col] == player.Unknown
}

func (b *Board) Take(row, col int, player player.Player) error {
	if !b.CanTake(row, col) {
		return fmt.Errorf("invalid move: %d,%d", row, col)
	}
	b.Grid[row][col] = player
	for _, o := range b.observers {
		if err := o.NotifyUpdate(row, col, player); err != nil {
			return fmt.Errorf("error while notifying an update: %v", err)
		}
	}
	return nil
}

// Finished returns if the game is finished. The second return value indicates
// the winning player or Unknown if it's draw.
func (b *Board) Finished() (bool, player.Player) {
	// Horizontal
	for r := 0; r < 3; r++ {
		p := b.Grid[r][0]
		if p != player.Unknown && b.Grid[r][1] == p && b.Grid[r][2] == p {
			return true, p
		}
	}
	// Vertical
	for c := 0; c < 3; c++ {
		p := b.Grid[0][c]
		if p != player.Unknown && b.Grid[1][c] == p && b.Grid[2][c] == p {
			return true, p
		}
	}
	// Diagonal
	if p := b.Grid[0][0]; p != player.Unknown && b.Grid[1][1] == p && b.Grid[2][2] == p {
		return true, p
	}
	if p := b.Grid[0][2]; p != player.Unknown && b.Grid[1][1] == p && b.Grid[2][0] == p {
		return true, p
	}
	// Check if there's still an unoccupied cell.
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if b.Grid[r][c] == player.Unknown {
				return false, player.Unknown
			}
		}
	}
	return true, player.Unknown
}

func (b *Board) String() string {
	var ret string
	for i := 0; i < 3; i++ {
		if i > 0 {
			ret += "\n"
		}
		for j := 0; j < 3; j++ {
			switch b.Grid[i][j] {
			case player.A:
				ret += "A"
			case player.B:
				ret += "B"
			case player.Unknown:
				ret += "."
			}
		}
	}
	return ret
}

func (b *Board) AddObserver(o BoardObserver) {
	b.observers = append(b.observers, o)
}
