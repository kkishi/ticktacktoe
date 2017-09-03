package client

import (
	"errors"
	"fmt"
	"log"

	"github.com/kkishi/ticktacktoe/game"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

const (
	Self     game.Player = game.PlayerA
	Opponent             = game.PlayerB
)

type Game struct {
	Board  game.Board
	Stream tpb.TickTackToe_GameClient
}

var ErrGameIsFinished = errors.New("game is finished")

func NewGame(stream tpb.TickTackToe_GameClient) *Game {
	return &Game{
		Stream: stream,
	}
}

func (g *Game) Join(name string) error {
	if err := g.Stream.Send(&tpb.Request{
		Event: &tpb.Request_Join{
			Join: &tpb.Join{
				Name: name,
			},
		},
	}); err != nil {
		g.Close()
		return fmt.Errorf("error while sending a join request; %v", err)
	}
	return nil
}

func (g *Game) Wait() error {
	r, err := g.Stream.Recv()
	if err != nil {
		g.Close()
		return fmt.Errorf("error while waiting for a response; %v", err)
	}
	if f := r.GetFinish(); f != nil {
		g.Close()
		if f.GetError() {
			return fmt.Errorf("game finished with an error")
		}
		m := f.GetOpponent()
		if err := g.Board.Take(int(m.GetRow()), int(m.GetCol()), Opponent); err != nil {
			return fmt.Errorf("invalid move (%d, %d) returned from server; %v",
				m.GetRow(), m.GetCol(), err)
		}
		return ErrGameIsFinished
	}
	mm := r.GetMakeMove()
	if mm == nil {
		return fmt.Errorf("expected make move response; got %v", r)
	}
	if !mm.GetInitial() {
		m := mm.GetOpponent()
		if err := g.Board.Take(int(m.GetRow()), int(m.GetCol()), Opponent); err != nil {
			return fmt.Errorf("invalid move (%d, %d) returned from server; %v",
				m.GetRow(), m.GetCol(), err)
		}
	}
	return nil
}

func (g *Game) MakeMove(r, c int) error {
	if err := g.Board.Take(r, c, Self); err != nil {
		return fmt.Errorf("invalid move (%d, %d) requested; %v", r, c, err)
	}
	if err := g.Stream.Send(&tpb.Request{
		Event: &tpb.Request_Move{
			Move: &tpb.Move{
				Row: int32(r),
				Col: int32(c),
			},
		},
	}); err != nil {
		return fmt.Errorf("error while sending a move request; %v", err)
	}
	return nil
}

func (g *Game) Close() {
	if err := g.Stream.CloseSend(); err != nil {
		// Not much we can do here; just log the error.
		log.Printf("closing client stream failed; %v", err)
	}
}
