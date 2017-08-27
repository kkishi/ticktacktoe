package server

import (
	"io"
	"log"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Impl struct{}

func (i *Impl) Game(stream tpb.TickTackToe_GameServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("in: %v\n", in)
		res := &tpb.Response{}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
