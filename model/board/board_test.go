package board

import (
	"testing"

	"github.com/kkishi/ticktacktoe/model/player"
)

func newTestBoard() *Board {
	return &Board{
		Grid: [3][3]player.Player{
			{player.Unknown, player.Unknown, player.Unknown},
			{player.Unknown, player.A, player.Unknown},
			{player.Unknown, player.Unknown, player.B},
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
		err := b.Take(test.row, test.col, player.A)
		if test.wantError && err == nil {
			t.Errorf("Take(%d,%d) returned no error; want an error",
				test.row, test.col)
		}
		if !test.wantError && err != nil {
			t.Errorf("Take(%d,%d) returned an error %v; want no error",
				test.row, test.col, err)
		}
		if !test.wantError && b.Grid[test.row][test.col] != player.A {
			t.Errorf("Board[%d][%d] is occupied by player %v; want %v",
				test.row, test.col, b.Grid[test.row][test.col], player.A)
		}
	}
}

func TestBoardFinished(t *testing.T) {
	tests := []struct {
		grid         [3][3]player.Player
		wantFinished bool
		wantPlayer   player.Player
	}{
		{
			// Not finished
			grid:         [3][3]player.Player{},
			wantFinished: false,
			wantPlayer:   player.Unknown,
		},
		{
			// Draw
			grid: [3][3]player.Player{
				{player.A, player.B, player.A},
				{player.B, player.A, player.A},
				{player.B, player.A, player.B},
			},
			wantFinished: true,
			wantPlayer:   player.Unknown,
		},
		{
			// Horizontal
			grid: [3][3]player.Player{
				{player.Unknown, player.B, player.B},
				{player.A, player.A, player.A},
				{player.Unknown, player.B, player.Unknown},
			},
			wantFinished: true,
			wantPlayer:   player.A,
		},
		{
			// Vertical
			grid: [3][3]player.Player{
				{player.Unknown, player.B, player.A},
				{player.A, player.B, player.A},
				{player.Unknown, player.B, player.Unknown},
			},
			wantFinished: true,
			wantPlayer:   player.B,
		},
		{
			// Horizontal
			grid: [3][3]player.Player{
				{player.Unknown, player.B, player.B},
				{player.A, player.A, player.A},
				{player.Unknown, player.B, player.Unknown},
			},
			wantFinished: true,
			wantPlayer:   player.A,
		},
		{
			// Diagonal
			grid: [3][3]player.Player{
				{player.B, player.B, player.A},
				{player.B, player.A, player.Unknown},
				{player.A, player.Unknown, player.Unknown},
			},
			wantFinished: true,
			wantPlayer:   player.A,
		},
		{
			// Diagonal
			grid: [3][3]player.Player{
				{player.B, player.A, player.A},
				{player.A, player.B, player.Unknown},
				{player.Unknown, player.Unknown, player.B},
			},
			wantFinished: true,
			wantPlayer:   player.B,
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
		Grid: [3][3]player.Player{
			{player.B, player.A, player.A},
			{player.A, player.B, player.Unknown},
			{player.Unknown, player.Unknown, player.B},
		},
	}).String()
	want := "BAA\nAB.\n..B"
	if got != want {
		t.Errorf("Board.String returned %q; want %q", got, want)
	}
}
