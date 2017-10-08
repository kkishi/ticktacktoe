package client

import (
	"context"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type Client struct {
	Context context.Context
	Cancel  context.CancelFunc
	Stream  tpb.TickTackToe_GameServer
}

func New(stream tpb.TickTackToe_GameServer) *Client {
	ctx, cancel := context.WithCancel(stream.Context())
	return &Client{
		Context: ctx,
		Cancel:  cancel,
		Stream:  stream,
	}
}

func (c *Client) Info(text string) error {
	return c.Stream.Send(&tpb.Response{
		Event: &tpb.Response_Info{
			&tpb.Message{
				Text: text,
			},
		},
	})
}
