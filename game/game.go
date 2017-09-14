package game

import (
	"errors"
	"fmt"
	"log"
	"sync"

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
				MakeMove: &tpb.MakeMove{},
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
		return fmt.Errorf("error while taking cell (%d, %d) for player %v; %v", m.GetRow(), m.GetCol(), s.player, err)
	}
	log.Printf("player %v made move (%d, %d)", s.player, m.GetRow(), m.GetCol())
	log.Printf("board: \n%s", s.game.Board.String())
	if finished, winningPlayer := s.game.Board.Finished(); finished {
		if winningPlayer == UnknownPlayer {
			log.Print("draw")
		} else {
			log.Printf("player %v wins game", winningPlayer)
		}
		// Notify clients that the game has finished.
		for i, stream := range s.game.Streams {
			var r tpb.Finish_Result
			if winningPlayer == UnknownPlayer {
				r = tpb.Finish_DRAW
			} else if PlayerFromIndex(i) == winningPlayer {
				r = tpb.Finish_WIN
			} else {
				r = tpb.Finish_LOSE
			}
			if err := stream.Send(&tpb.Response{
				Event: &tpb.Response_Finish{
					Finish: &tpb.Finish{
						Result: r,
					},
				},
			}); err != nil {
				// Not much we can do here; just log the fact.
				log.Printf("error while sending a finish response to player %v; %v", PlayerFromIndex(i), err)
			}
		}
		return ErrGameIsFinished
	}
	next := s.player.Next()
	if err := s.game.Streams[next.ToIndex()].Send(&tpb.Response{
		Event: &tpb.Response_MakeMove{
			MakeMove: &tpb.MakeMove{},
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

	// The server thrad for PlayerA should wait on this condition. For PlayerB
	// the thread should call Game.Start.
	Cond *sync.Cond
}

func New() *Game {
	g := &Game{
		Cond: sync.NewCond(new(sync.Mutex)),
	}
	g.State = &setupState{
		player: PlayerA,
		game:   g,
	}
	return g
}

func (g *Game) Start() {
	log.Print("game started")
	for {
		if err := g.State.Handle(); err == ErrGameIsFinished {
			break
		} else if err != nil {
			g.FinishWithError(err)
			break
		}
	}
	g.Cond.Broadcast()
}

func (g *Game) Waiting() bool {
	return len(g.Streams) < 2
}

type boardObserver struct {
	stream tpb.TickTackToe_GameServer
}

func (o *boardObserver) NotifyUpdate(row, col int, player Player) error {
	if err := o.stream.Send(&tpb.Response{
		Event: &tpb.Response_Update{
			&tpb.Update{
				Row:    int32(row),
				Col:    int32(col),
				Player: tpb.Player(player),
			},
		},
	}); err != nil {
		return fmt.Errorf("error while notifying an update: %v", err)
	}
	return nil
}

func (g *Game) Join(stream tpb.TickTackToe_GameServer) {
	g.Streams = append(g.Streams, stream)
	g.Board.AddObserver(&boardObserver{
		stream: stream,
	})
}

func (g *Game) FinishWithError(err error) {
	for i := 0; i < 2; i++ {
		if err := g.Streams[i].Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Result: tpb.Finish_ERROR,
				},
			},
		}); err != nil {
			// Not much we can do here; just log the fact.
			log.Printf("error while sending an error finish response to player %v; %v", PlayerFromIndex(i), err)
		}
	}
}
