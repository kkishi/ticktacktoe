package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/kkishi/ticktacktoe/model/player"
	"google.golang.org/grpc"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

var name = flag.String("name", "", "Name of the player")

func main() {
	flag.Parse()
	if *name == "" {
		log.Fatal("--name is required")
	}

	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to the server: %v", err)
	}
	defer conn.Close()

	c := tpb.NewTickTackToeClient(conn)

	stream, err := c.Game(context.Background())
	if err != nil {
		log.Fatalf("failed to connect to the server: %v", err)
	}

	g := NewGame(stream)
	defer g.Close()

	if err := g.Join(*name); err != nil {
		log.Fatalf("failed to join a game; %v", err)
	}
	log.Print("joined to a game")

	for {
		if err := g.Wait(); err == ErrGameIsFinished {
			finished, p := g.Board.Finished()
			if !finished {
				log.Fatal("game is not finished locally")
			}
			if p == player.Unknown {
				fmt.Print("Draw.")
			} else {
				fmt.Printf("Player %v wins.", p)
			}
			fmt.Printf(" final board:\n%s\n", g.Board.String())
			return
		} else if err != nil {
			log.Fatalf("game finished with an error; %v", err)
		}
		fmt.Printf("board:\n%s\n", g.Board.String())
		var r, c int
		for {
			fmt.Print("> ")
			if n, err := fmt.Scanf("%d %d", &r, &c); err == io.EOF {
				return
			} else if err == nil && n == 2 {
				break
			}
		}
		err := g.MakeMove(r, c)
		if err != nil {
			log.Fatalf("error while making a move; %v", err)
		}
	}
}
