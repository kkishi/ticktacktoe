package game

import (
	"fmt"
)

type BoardObserver interface {
	NotifyUpdate(row, col int, player Player) error
}

type Board struct {
	Grid      [3][3]Player
	observers []BoardObserver
}

func (b *Board) CanTake(row, col int) bool {
	return 0 <= row && row <= 2 && 0 <= col && col <= 2 &&
		b.Grid[row][col] == UnknownPlayer
}

func (b *Board) Take(row, col int, player Player) error {
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
// the winning player or UnknownPlayer if it's draw.
func (b *Board) Finished() (bool, Player) {
	// Horizontal
	for r := 0; r < 3; r++ {
		p := b.Grid[r][0]
		if p != UnknownPlayer && b.Grid[r][1] == p && b.Grid[r][2] == p {
			return true, p
		}
	}
	// Vertical
	for c := 0; c < 3; c++ {
		p := b.Grid[0][c]
		if p != UnknownPlayer && b.Grid[1][c] == p && b.Grid[2][c] == p {
			return true, p
		}
	}
	// Diagonal
	if p := b.Grid[0][0]; p != UnknownPlayer && b.Grid[1][1] == p && b.Grid[2][2] == p {
		return true, p
	}
	if p := b.Grid[0][2]; p != UnknownPlayer && b.Grid[1][1] == p && b.Grid[2][0] == p {
		return true, p
	}
	// Check if there's still an unoccupied cell.
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if b.Grid[r][c] == UnknownPlayer {
				return false, UnknownPlayer
			}
		}
	}
	return true, UnknownPlayer
}

func (b *Board) String() string {
	var ret string
	for i := 0; i < 3; i++ {
		if i > 0 {
			ret += "\n"
		}
		for j := 0; j < 3; j++ {
			switch b.Grid[i][j] {
			case PlayerA:
				ret += "A"
			case PlayerB:
				ret += "B"
			case UnknownPlayer:
				ret += "."
			}
		}
	}
	return ret
}

func (b *Board) AddObserver(o BoardObserver) {
	b.observers = append(b.observers, o)
}
