package server

import tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"

const Version = "0.1"

type Impl struct{}

func (i *Impl) Game(stream tpb.TickTackToe_GameServer) error {
	return nil
}
