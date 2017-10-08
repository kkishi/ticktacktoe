package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/kkishi/ticktacktoe/model/board"
	"github.com/kkishi/ticktacktoe/model/player"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Game struct {
	Board  board.Board
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
	for {
		r, err := g.Stream.Recv()
		if err != nil {
			g.Close()
			return fmt.Errorf("error while waiting for a response; %v", err)
		}
		if i := r.GetInfo(); i != nil {
			log.Printf("info: %q", i.GetText())
			continue
		}
		if f := r.GetFinish(); f != nil {
			g.Close()
			log.Printf("game has finished; %v", f)
			if f.GetResult() == tpb.Finish_ERROR {
				return fmt.Errorf("game finished with an error")
			}
			return ErrGameIsFinished
		}
		if u := r.GetUpdate(); u != nil {
			if err := g.Board.Take(int(u.GetRow()), int(u.GetCol()),
				player.Player(u.GetPlayer())); err != nil {
				return fmt.Errorf("invalid move (%d, %d) returned from server; %v",
					u.GetRow(), u.GetCol(), err)
			}
			continue
		}
		mm := r.GetMakeMove()
		if mm == nil {
			return fmt.Errorf("expected make move response; got %v", r)
		}
		return nil
	}
}

func (g *Game) MakeMove(r, c int) error {
	if !g.Board.CanTake(r, c) {
		return fmt.Errorf("invalid move (%d, %d) requested; %v", r, c)
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
	log.Printf("made move (%d, %d)", r, c)
	return nil
}

func (g *Game) Close() {
	if err := g.Stream.CloseSend(); err != nil {
		// Not much we can do here; just log the error.
		log.Printf("closing client stream failed; %v", err)
	}
}
