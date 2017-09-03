package game

import (
	"errors"
	"fmt"
	"log"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type GameState interface {
	Handle() error
}

type setupState struct {
	player Player
	game   *Game
}

var ErrGameIsFinished = errors.New("game is finished")

func (s *setupState) Handle() error {
	r, err := s.game.Streams[s.player.ToIndex()].Recv()
	if err != nil {
		return fmt.Errorf("error while waiting for a join request from player %v; %v", s.player, err)
	}
	j := r.GetJoin()
	if j == nil {
		return fmt.Errorf("expected a join request from player %v; got %v", s.player, r)
	}
	s.game.Names = append(s.game.Names, j.GetName())
	log.Printf("player %v joind: %v", s.player, r)
	if s.player == PlayerA {
		s.player = PlayerB
	} else {
		s.game.State = &runningState{
			player: PlayerA,
			game:   s.game,
		}
		if err := s.game.Streams[PlayerA.ToIndex()].Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{
					Initial: true,
				},
			},
		}); err != nil {
			return fmt.Errorf("error while sending make move response to player %v; %v", PlayerA, err)
		}
	}
	return nil
}

type runningState struct {
	player Player
	game   *Game
}

func (s *runningState) Handle() error {
	r, err := s.game.Streams[s.player.ToIndex()].Recv()
	if err != nil {
		return fmt.Errorf("error while waiting for move request from player %v: %v", s.player, err)
	}
	m := r.GetMove()
	if m == nil {
		return fmt.Errorf("expected a move request from player %v; got %v", s.player, r)
	}
	if err := s.game.Board.Take(int(m.GetRow()), int(m.GetCol()), s.player); err != nil {
		return fmt.Errorf("invalid move request (%d, %d) from player %v; %v", m.GetRow(), m.GetCol(), s.player, err)
	}
	log.Printf("player %v made move %v", s.player, r)
	log.Printf("Board: \n%s", s.game.Board.String())
	if s.game.Board.WinningPlayer() != UnknownPlayer {
		return ErrGameIsFinished
	}
	next := s.player.Next()
	if err := s.game.Streams[next.ToIndex()].Send(&tpb.Response{
		Event: &tpb.Response_MakeMove{
			MakeMove: &tpb.MakeMove{
				Opponent: m,
			},
		},
	}); err != nil {
		return fmt.Errorf("error while sending make move response to player %v; %v", next, err)
	}
	s.player = next
	return nil
}

type Game struct {
	Board   Board
	State   GameState
	Streams []tpb.TickTackToe_GameServer
	Names   []string
}

func New() *Game {
	g := &Game{}
	g.State = &setupState{
		player: PlayerA,
		game:   g,
	}
	return g
}

func (g *Game) Start() {
	for {
		if err := g.State.Handle(); err == ErrGameIsFinished {
			g.Finish()
			break
		} else if err != nil {
			g.FinishWithError(err)
			break
		}
	}
}

func (g *Game) Waiting() bool {
	return len(g.Streams) < 2
}

func (g *Game) Join(stream tpb.TickTackToe_GameServer) {
	g.Streams = append(g.Streams, stream)
}

func (g *Game) FinishWithError(err error) {
	for i := 0; i < 2; i++ {
		if err := g.Streams[i].Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Error: true,
				},
			},
		}); err != nil {
			// Not much we can do here; just log the fact.
			log.Printf("error while sending an error finish response to player %v; %v", PlayerFromIndex(i), err)
		}
	}
}

func (g *Game) Finish() {
	win := g.Board.WinningPlayer()
	for i := 0; i < 2; i++ {
		if err := g.Streams[i].Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Win: PlayerFromIndex(i) == win,
				},
			},
		}); err != nil {
			// Not much we can do here; just log the fact.
			log.Printf("error while sending an error finish response to player %v; %v", PlayerFromIndex(i), err)
		}
	}
}
