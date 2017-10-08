package game

import (
	"fmt"
	"log"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Game struct {
	Board   Board
	Streams []tpb.TickTackToe_GameServer
	Names   []string
}

func New(a, b tpb.TickTackToe_GameServer) *Game {
	return &Game{
		Streams: []tpb.TickTackToe_GameServer{a, b},
	}
}

func (g *Game) Start() {
	log.Print("game started")

	for i := 0; i < 2; i++ {
		if err := g.initPlayer(i); err != nil {
			g.Finish(err)
			return
		}
	}

	for i := 0; ; i = 1 - i {
		if err := g.Streams[1-i].Send(&tpb.Response{
			Event: &tpb.Response_Info{
				&tpb.Message{
					Text: "Waiting for the opponent to make a move.",
				},
			},
		}); err != nil {
			g.Finish(fmt.Errorf("error while sending an info to player %v; %v",
				PlayerFromIndex(1-i), err))
			return
		}
		if err := g.makeMove(i); err != nil {
			g.Finish(err)
			return
		}
		if finished, winningPlayer := g.Board.Finished(); finished {
			if winningPlayer == UnknownPlayer {
				log.Print("draw")
			} else {
				log.Printf("player %v wins game", winningPlayer)
			}
			// Notify clients that the game has finished.
			for i, stream := range g.Streams {
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
			return
		}
	}
}

func (g *Game) initPlayer(i int) error {
	p := PlayerFromIndex(i)

	r, err := g.Streams[i].Recv()
	if err != nil {
		return fmt.Errorf("error while waiting for a join request from player %v; %v", p, err)
	}
	j := r.GetJoin()
	if j == nil {
		return fmt.Errorf("expected a join request from player %v; got %v", p, r)
	}
	g.Names = append(g.Names, j.GetName())
	log.Printf("player %v joind: %v", p, r)

	if err := g.Streams[i].Send(&tpb.Response{
		Event: &tpb.Response_Info{
			&tpb.Message{
				Text: "Game is started.",
			},
		},
	}); err != nil {
		return fmt.Errorf("error while sending an info for a player %v: %v", p, err)
	}
	g.Board.AddObserver(&boardObserver{
		stream: g.Streams[i],
	})
	return nil
}

func (g *Game) makeMove(i int) error {
	p := PlayerFromIndex(i)
	if err := g.Streams[i].Send(&tpb.Response{
		Event: &tpb.Response_MakeMove{
			MakeMove: &tpb.MakeMove{},
		},
	}); err != nil {
		return fmt.Errorf("error while sending make move response to player %v; %v", p, err)
	}

	r, err := g.Streams[i].Recv()
	if err != nil {
		return fmt.Errorf("error while waiting for move request from player %v: %v", p, err)
	}
	m := r.GetMove()
	if m == nil {
		return fmt.Errorf("expected a move request from player %v; got %v", p, r)
	}
	if err := g.Board.Take(int(m.GetRow()), int(m.GetCol()), p); err != nil {
		return fmt.Errorf("error while taking cell (%d, %d) for player %v; %v", m.GetRow(), m.GetCol(), p, err)
	}
	log.Printf("player %v made move (%d, %d)", p, m.GetRow(), m.GetCol())
	log.Printf("board: \n%s", g.Board.String())
	return nil
}

func (g *Game) Finish(err error) {
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
