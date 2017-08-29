package game

import "testing"

func newTestGame() *Game {
	return &Game{
		Board: [3][3]Player{
			{UnknownPlayer, UnknownPlayer, UnknownPlayer},
			{UnknownPlayer, PlayerA, UnknownPlayer},
			{UnknownPlayer, UnknownPlayer, PlayerB},
		},
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		row       int
		col       int
		wantError bool
	}{
		{0, 0, false},
		{1, 1, true},
		{2, 2, true},
	}
	for _, test := range tests {
		g := newTestGame()
		err := g.Take(test.row, test.col, PlayerA)
		if test.wantError && err == nil {
			t.Errorf("Take(%d,%d) returned no error; want an error",
				test.row, test.col)
		}
		if !test.wantError && err != nil {
			t.Errorf("Take(%d,%d) returned an error %v; want no error",
				test.row, test.col, err)
		}
		if !test.wantError && g.Board[test.row][test.col] != PlayerA {
			t.Errorf("Board[%d][%d] is occupied by player %v; want %v",
				g.Board[test.row][test.col], PlayerA)
		}
	}
}

func TestWinningPlayer(t *testing.T) {
	tests := []struct {
		game Game
		want Player
	}{
		{
			game: Game{},
			want: UnknownPlayer,
		},
		{
			// Horizontal
			game: Game{
				Board: [3][3]Player{
					{UnknownPlayer, PlayerB, PlayerB},
					{PlayerA, PlayerA, PlayerA},
					{UnknownPlayer, PlayerB, UnknownPlayer},
				},
			},
			want: PlayerA,
		},
		{
			// Vertical
			game: Game{
				Board: [3][3]Player{
					{UnknownPlayer, PlayerB, PlayerA},
					{PlayerA, PlayerB, PlayerA},
					{UnknownPlayer, PlayerB, UnknownPlayer},
				},
			},
			want: PlayerB,
		},
		{
			// Horizontal
			game: Game{
				Board: [3][3]Player{
					{UnknownPlayer, PlayerB, PlayerB},
					{PlayerA, PlayerA, PlayerA},
					{UnknownPlayer, PlayerB, UnknownPlayer},
				},
			},
			want: PlayerA,
		},
		{
			// Diagonal
			game: Game{
				Board: [3][3]Player{
					{PlayerB, PlayerB, PlayerA},
					{PlayerB, PlayerA, UnknownPlayer},
					{PlayerA, UnknownPlayer, UnknownPlayer},
				},
			},
			want: PlayerA,
		},
		{
			// Diagonal
			game: Game{
				Board: [3][3]Player{
					{PlayerB, PlayerA, PlayerA},
					{PlayerA, PlayerB, UnknownPlayer},
					{UnknownPlayer, UnknownPlayer, PlayerB},
				},
			},
			want: PlayerB,
		},
	}
	for _, test := range tests {
		if got := test.game.WinningPlayer(); got != test.want {
			t.Errorf("WinningPlayer() returned %v; want %v", got, test.want)
		}
	}
}
