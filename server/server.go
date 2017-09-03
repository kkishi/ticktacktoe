package server

import (
	"sync"

	"github.com/kkishi/ticktacktoe/game"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Impl struct {
	games      []*game.Game
	gamesMutex sync.Mutex
}

func (i *Impl) Game(stream tpb.TickTackToe_GameServer) error {
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
	}
	g.Join(stream)
	return nil
}
