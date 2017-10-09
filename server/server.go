package main

import (
	"fmt"
	"log"

	"github.com/kkishi/ticktacktoe/server/client"
	"github.com/kkishi/ticktacktoe/server/game"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Server struct {
	ch chan *client.Client
}

func NewServer() *Server {
	s := &Server{
		ch: make(chan *client.Client),
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
			case <-a.Context.Done():
				a = <-s.ch
			case <-b.Context.Done():
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

	c := client.New(stream)

	if err := c.Info("Waiting for a game to start."); err != nil {
		return fmt.Errorf("error while sending an info: %v", err)
	}

	// Add the client to the waiting queue.
	s.ch <- c

	// Wait on the client to be done. See the comment on client.Client for details.
	<-c.Context.Done()

	log.Print("a Game connection closed")
	return nil
}
