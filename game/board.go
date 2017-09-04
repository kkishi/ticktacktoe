package game

import (
	"fmt"
)

type Board [3][3]Player

// Take takes a cell at (row, col) for a player. When error is not nil Board
// is unchanged.
func (b *Board) Take(row, col int, player Player) error {
	if row < 0 || 2 < row ||
		col < 0 || 2 < col {
		return fmt.Errorf("invalid cell to take: %d,%d", row, col)
	}
	if b[row][col] != UnknownPlayer {
		return fmt.Errorf("cell %d,%d is already taken", row, col)
	}
	b[row][col] = player
	return nil
}

// Finished returns if the game is finished. The second return value indicates
// the winning player or UnknownPlayer if it's draw.
func (b *Board) Finished() (bool, Player) {
	// Horizontal
	for r := 0; r < 3; r++ {
		p := b[r][0]
		if p != UnknownPlayer && b[r][1] == p && b[r][2] == p {
			return true, p
		}
	}
	// Vertical
	for c := 0; c < 3; c++ {
		p := b[0][c]
		if p != UnknownPlayer && b[1][c] == p && b[2][c] == p {
			return true, p
		}
	}
	// Diagonal
	if p := b[0][0]; p != UnknownPlayer && b[1][1] == p && b[2][2] == p {
		return true, p
	}
	if p := b[0][2]; p != UnknownPlayer && b[1][1] == p && b[2][0] == p {
		return true, p
	}
	// Check if there's still an unoccupied cell.
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if b[r][c] == UnknownPlayer {
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
			switch b[i][j] {
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
