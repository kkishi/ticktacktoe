package game

import "testing"

func newTestBoard() *Board {
	return &Board{
		{UnknownPlayer, UnknownPlayer, UnknownPlayer},
		{UnknownPlayer, PlayerA, UnknownPlayer},
		{UnknownPlayer, UnknownPlayer, PlayerB},
	}
}

func TestBoardTake(t *testing.T) {
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
		b := newTestBoard()
		err := b.Take(test.row, test.col, PlayerA)
		if test.wantError && err == nil {
			t.Errorf("Take(%d,%d) returned no error; want an error",
				test.row, test.col)
		}
		if !test.wantError && err != nil {
			t.Errorf("Take(%d,%d) returned an error %v; want no error",
				test.row, test.col, err)
		}
		if !test.wantError && b[test.row][test.col] != PlayerA {
			t.Errorf("Board[%d][%d] is occupied by player %v; want %v",
				test.row, test.col, b[test.row][test.col], PlayerA)
		}
	}
}

func TestBoardWinningPlayer(t *testing.T) {
	tests := []struct {
		board *Board
		want  Player
	}{
		{
			board: &Board{},
			want:  UnknownPlayer,
		},
		{
			// Horizontal
			board: &Board{
				{UnknownPlayer, PlayerB, PlayerB},
				{PlayerA, PlayerA, PlayerA},
				{UnknownPlayer, PlayerB, UnknownPlayer},
			},
			want: PlayerA,
		},
		{
			// Vertical
			board: &Board{
				{UnknownPlayer, PlayerB, PlayerA},
				{PlayerA, PlayerB, PlayerA},
				{UnknownPlayer, PlayerB, UnknownPlayer},
			},
			want: PlayerB,
		},
		{
			// Horizontal
			board: &Board{
				{UnknownPlayer, PlayerB, PlayerB},
				{PlayerA, PlayerA, PlayerA},
				{UnknownPlayer, PlayerB, UnknownPlayer},
			},
			want: PlayerA,
		},
		{
			// Diagonal
			board: &Board{
				{PlayerB, PlayerB, PlayerA},
				{PlayerB, PlayerA, UnknownPlayer},
				{PlayerA, UnknownPlayer, UnknownPlayer},
			},
			want: PlayerA,
		},
		{
			// Diagonal
			board: &Board{
				{PlayerB, PlayerA, PlayerA},
				{PlayerA, PlayerB, UnknownPlayer},
				{UnknownPlayer, UnknownPlayer, PlayerB},
			},
			want: PlayerB,
		},
	}
	for _, test := range tests {
		if got := test.board.WinningPlayer(); got != test.want {
			t.Errorf("WinningPlayer() returned %v; want %v", got, test.want)
		}
	}
}

func TestBoardString(t *testing.T) {
	got := (&Board{
		{PlayerB, PlayerA, PlayerA},
		{PlayerA, PlayerB, UnknownPlayer},
		{UnknownPlayer, UnknownPlayer, PlayerB},
	}).String()
	want := "BAA\nAB.\n..B"
	if got != want {
		t.Errorf("Board.String returned %q; want %q", got, want)
	}
}
