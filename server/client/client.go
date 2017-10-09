package client

import (
	"context"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

// Client is a wrapper around a stream, having a Context (derived from
// Stream.Context()) and its CancelFunc. Cancel can be called in the Game
// execution logic. This way the server's Game gRPC request handler can wait
// on Context.Done() for both client connection closures and Game closures.
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
