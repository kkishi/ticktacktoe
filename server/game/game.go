package game

import (
	"fmt"
	"log"

	"github.com/kkishi/ticktacktoe/model/board"
	"github.com/kkishi/ticktacktoe/model/player"
	"github.com/kkishi/ticktacktoe/server/client"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Game struct {
	Board   board.Board
	Clients []*client.Client
	Names   []string
}

func New(a, b *client.Client) *Game {
	return &Game{
		Clients: []*client.Client{a, b},
	}
}

func (g *Game) Start() {
	log.Print("game started")
	defer g.Close()

	for i := 0; i < 2; i++ {
		if err := g.initPlayer(i); err != nil {
			g.Error(err)
			return
		}
	}

	for i := 0; ; i = 1 - i {
		if err := g.Clients[1-i].Info("Waiting for the opponent to make a move."); err != nil {
			g.Error(fmt.Errorf("error while sending an info to player %v; %v",
				player.FromIndex(1-i), err))
			return
		}
		if err := g.makeMove(i); err != nil {
			g.Error(err)
			return
		}
		if finished, winningPlayer := g.Board.Finished(); finished {
			if winningPlayer == player.Unknown {
				log.Print("draw")
			} else {
				log.Printf("player %v wins game", winningPlayer)
			}
			// Notify clients that the game has finished.
			for i, c := range g.Clients {
				var r tpb.Finish_Result
				if winningPlayer == player.Unknown {
					r = tpb.Finish_DRAW
				} else if player.FromIndex(i) == winningPlayer {
					r = tpb.Finish_WIN
				} else {
					r = tpb.Finish_LOSE
				}
				if err := c.Stream.Send(&tpb.Response{
					Event: &tpb.Response_Finish{
						Finish: &tpb.Finish{
							Result: r,
						},
					},
				}); err != nil {
					// Not much we can do here; just log the fact.
					log.Printf("error while sending a finish response to player %v; %v", player.FromIndex(i), err)
				}
			}
			return
		}
	}
}

func (g *Game) Close() {
	for _, c := range g.Clients {
		c.Cancel()
	}
}

func (g *Game) initPlayer(i int) error {
	p := player.FromIndex(i)

	r, err := g.Clients[i].Stream.Recv()
	if err != nil {
		return fmt.Errorf("error while waiting for a join request from player %v; %v", p, err)
	}
	j := r.GetJoin()
	if j == nil {
		return fmt.Errorf("expected a join request from player %v; got %v", p, r)
	}
	g.Names = append(g.Names, j.GetName())
	log.Printf("player %v joind: %v", p, r)

	if err := g.Clients[i].Info("Game is started."); err != nil {
		return fmt.Errorf("error while sending an info for a player %v: %v", p, err)
	}
	g.Board.AddObserver(&boardObserver{
		stream: g.Clients[i].Stream,
	})
	return nil
}

func (g *Game) makeMove(i int) error {
	p := player.FromIndex(i)
	if err := g.Clients[i].Stream.Send(&tpb.Response{
		Event: &tpb.Response_MakeMove{
			MakeMove: &tpb.MakeMove{},
		},
	}); err != nil {
		return fmt.Errorf("error while sending make move response to player %v; %v", p, err)
	}

	r, err := g.Clients[i].Stream.Recv()
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

func (g *Game) Error(err error) {
	for i := 0; i < 2; i++ {
		if err := g.Clients[i].Stream.Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Result: tpb.Finish_ERROR,
				},
			},
		}); err != nil {
			// Not much we can do here; just log the fact.
			log.Printf("error while sending an error finish response to player %v; %v", player.FromIndex(i), err)
		}
	}
}

type boardObserver struct {
	stream tpb.TickTackToe_GameServer
}

func (o *boardObserver) NotifyUpdate(row, col int, player player.Player) error {
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
