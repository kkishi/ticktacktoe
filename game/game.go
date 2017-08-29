package game

import "fmt"

type Player int

const (
	UnknownPlayer Player = iota
	PlayerA
	PlayerB
)

type GameState int

const (
	WaitingForPlayer GameState = iota
	WaitingForMoveA
	WaitingForMoveB
)

type Game struct {
	Board [3][3]Player
	State GameState
}

func New() *Game {
	return new(Game)
}

func (g *Game) Take(row, col int, player Player) error {
	if row < 0 || 2 < row ||
		col < 0 || 2 < col {
		return fmt.Errorf("invalid cell to take: %d,%d", row, col)
	}
	if g.Board[row][col] != UnknownPlayer {
		return fmt.Errorf("cell %d,%d is already taken", row, col)
	}
	g.Board[row][col] = player
	return nil
}

func (g *Game) WinningPlayer() Player {
	// Horizontal
	for r := 0; r < 3; r++ {
		p := g.Board[r][0]
		if p != UnknownPlayer && g.Board[r][1] == p && g.Board[r][2] == p {
			return p
		}
	}
	// Vertical
	for c := 0; c < 3; c++ {
		p := g.Board[0][c]
		if p != UnknownPlayer && g.Board[1][c] == p && g.Board[2][c] == p {
			return p
		}
	}
	// Diagonal
	if p := g.Board[0][0]; p != UnknownPlayer && g.Board[1][1] == p && g.Board[2][2] == p {
		return p
	}
	if p := g.Board[0][2]; p != UnknownPlayer && g.Board[1][1] == p && g.Board[2][0] == p {
		return p
	}
	return UnknownPlayer
}
