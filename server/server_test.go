package main

import tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"

var _ tpb.TickTackToeServer = (*Server)(nil)
