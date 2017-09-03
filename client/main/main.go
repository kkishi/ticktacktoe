package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	"github.com/kkishi/ticktacktoe/client"
	"github.com/kkishi/ticktacktoe/game"
	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

var name = flag.String("name", "", "Name of the player")

func main() {
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

	g := client.NewGame(stream)
	defer g.Close()

	if err := g.Join(*name); err != nil {
		log.Fatalf("failed to join a game; %v", err)
	}

	for {
		if err := g.Wait(); err == client.ErrGameIsFinished {
			switch g.Board.WinningPlayer() {
			case client.Self:
				fmt.Println("you win the game")
			case client.Opponent:
				fmt.Println("you lost the game")
			case game.UnknownPlayer:
				fmt.Println("game finished a draw")
			}
			return
		} else if err != nil {
			log.Fatal("game finished with an error")
		}
		fmt.Printf("board:\n%s\n", g.Board.String())
		var r, c int
		for {
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
