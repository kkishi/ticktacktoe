package server

import (
	"log"
	"sync"

	"github.com/kkishi/ticktacktoe/game"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Impl struct {
	games      []*game.Game
	gamesMutex sync.Mutex
}

func (i *Impl) Join(stream tpb.TickTackToe_GameServer) *game.Game {
	i.gamesMutex.Lock()
	defer i.gamesMutex.Unlock()
	var g *game.Game
	for _, ga := range i.games {
		if ga.Waiting() {
			g = ga
			break
		}
	}
	if g == nil {
		g = game.New()
		i.games = append(i.games, g)
	}
	g.Join(stream)
	return g
}

func (i *Impl) Game(stream tpb.TickTackToe_GameServer) error {
	g := i.Join(stream)
	log.Print("a player join to a game")
	if !g.Waiting() {
		go g.Start()
	}
	g.Cond.L.Lock()
	defer g.Cond.L.Unlock()
	g.Cond.Wait()
	return nil
}
