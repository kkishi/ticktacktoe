package main

import (
	"fmt"
	"log"

	"github.com/kkishi/ticktacktoe/server/game"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Server struct {
	ch chan tpb.TickTackToe_GameServer
}

func NewServer() *Server {
	s := &Server{
		ch: make(chan tpb.TickTackToe_GameServer),
	}
	go s.SpawnGames()
	return s
}

func (s *Server) SpawnGames() {
	for {
		a := <-s.ch
		b := <-s.ch
	loop:
		for {
			select {
			case <-a.Context().Done():
				a = <-s.ch
			case <-b.Context().Done():
				b = <-s.ch
			default:
				break loop
			}
		}
		g := game.New(a, b)
		go g.Start()
	}
}

func (s *Server) Game(stream tpb.TickTackToe_GameServer) error {
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
	s.ch <- stream
	<-stream.Context().Done()
	log.Print("a Game connection closed")
	return nil
}
