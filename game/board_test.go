package game

import "testing"

func newTestBoard() *Board {
	return &Board{
		Grid: [3][3]Player{
			{UnknownPlayer, UnknownPlayer, UnknownPlayer},
			{UnknownPlayer, PlayerA, UnknownPlayer},
			{UnknownPlayer, UnknownPlayer, PlayerB},
		},
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
		if !test.wantError && b.Grid[test.row][test.col] != PlayerA {
			t.Errorf("Board[%d][%d] is occupied by player %v; want %v",
				test.row, test.col, b.Grid[test.row][test.col], PlayerA)
		}
	}
}

func TestBoardFinished(t *testing.T) {
	tests := []struct {
		grid         [3][3]Player
		wantFinished bool
		wantPlayer   Player
	}{
		{
			// Not finished
			grid:         [3][3]Player{},
			wantFinished: false,
			wantPlayer:   UnknownPlayer,
		},
		{
			// Draw
			grid: [3][3]Player{
				{PlayerA, PlayerB, PlayerA},
				{PlayerB, PlayerA, PlayerA},
				{PlayerB, PlayerA, PlayerB},
			},
			wantFinished: true,
			wantPlayer:   UnknownPlayer,
		},
		{
			// Horizontal
			grid: [3][3]Player{
				{UnknownPlayer, PlayerB, PlayerB},
				{PlayerA, PlayerA, PlayerA},
				{UnknownPlayer, PlayerB, UnknownPlayer},
			},
			wantFinished: true,
			wantPlayer:   PlayerA,
		},
		{
			// Vertical
			grid: [3][3]Player{
				{UnknownPlayer, PlayerB, PlayerA},
				{PlayerA, PlayerB, PlayerA},
				{UnknownPlayer, PlayerB, UnknownPlayer},
			},
			wantFinished: true,
			wantPlayer:   PlayerB,
		},
		{
			// Horizontal
			grid: [3][3]Player{
				{UnknownPlayer, PlayerB, PlayerB},
				{PlayerA, PlayerA, PlayerA},
				{UnknownPlayer, PlayerB, UnknownPlayer},
			},
			wantFinished: true,
			wantPlayer:   PlayerA,
		},
		{
			// Diagonal
			grid: [3][3]Player{
				{PlayerB, PlayerB, PlayerA},
				{PlayerB, PlayerA, UnknownPlayer},
				{PlayerA, UnknownPlayer, UnknownPlayer},
			},
			wantFinished: true,
			wantPlayer:   PlayerA,
		},
		{
			// Diagonal
			grid: [3][3]Player{
				{PlayerB, PlayerA, PlayerA},
				{PlayerA, PlayerB, UnknownPlayer},
				{UnknownPlayer, UnknownPlayer, PlayerB},
			},
			wantFinished: true,
			wantPlayer:   PlayerB,
		},
	}
	for _, test := range tests {
		gotFinished, gotPlayer := (&Board{Grid: test.grid}).Finished()
		if gotFinished != test.wantFinished || gotPlayer != test.wantPlayer {
			t.Errorf("Finished returned (%t, %v); want (%t, %v)",
				gotFinished, gotPlayer, test.wantFinished, test.wantPlayer)
		}
	}
}

func TestBoardString(t *testing.T) {
	got := (&Board{
		Grid: [3][3]Player{
			{PlayerB, PlayerA, PlayerA},
			{PlayerA, PlayerB, UnknownPlayer},
			{UnknownPlayer, UnknownPlayer, PlayerB},
		},
	}).String()
	want := "BAA\nAB.\n..B"
	if got != want {
		t.Errorf("Board.String returned %q; want %q", got, want)
	}
}
