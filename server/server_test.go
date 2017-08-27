package server

import tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"

var _ tpb.TickTackToeServer = (*Impl)(nil)
