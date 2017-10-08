package server

import (
	"fmt"
	"log"

	"github.com/kkishi/ticktacktoe/game"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Impl struct {
	ch chan tpb.TickTackToe_GameServer
}

func New() *Impl {
	i := &Impl{
		ch: make(chan tpb.TickTackToe_GameServer),
	}
	go i.SpawnGames()
	return i
}

func (i *Impl) SpawnGames() {
	for {
		a := <-i.ch
		b := <-i.ch
	loop:
		for {
			select {
			case <-a.Context().Done():
				a = <-i.ch
			case <-b.Context().Done():
				b = <-i.ch
			default:
				break loop
			}
		}
		g := game.New(a, b)
		go g.Start()
	}
}

func (i *Impl) Game(stream tpb.TickTackToe_GameServer) error {
	log.Print("new Game connection")
	if err := stream.Send(&tpb.Response{
		Event: &tpb.Response_Info{
			&tpb.Message{
				Text: "Waiting for a game to start.",
			},
		},
	}); err != nil {
		return fmt.Errorf("error while sending an info: %v", err)
	}
	i.ch <- stream
	<-stream.Context().Done()
	log.Print("a Game connection closed")
	return nil
}
