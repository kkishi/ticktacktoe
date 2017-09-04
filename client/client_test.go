package client

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkishi/ticktacktoe/proto/mock_ticktacktoe_proto"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

func TestJoin(t *testing.T) {
	const name = "Test Player"

	c := gomock.NewController(t)
	defer c.Finish()
	s := mock_ticktacktoe_proto.NewMockTickTackToe_GameClient(c)

	g := NewGame(s)

	gomock.InOrder(
		s.EXPECT().Send(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: name,
				},
			},
		}).Return(errors.New("fail")),
		s.EXPECT().CloseSend().Return((error)(nil)),
	)
	if err := g.Join(name); err == nil {
		t.Error("expected non-nil error from Join")
	}

	s.EXPECT().Send(&tpb.Request{
		Event: &tpb.Request_Join{
			Join: &tpb.Join{
				Name: name,
			},
		},
	}).Return((error)(nil))
	if err := g.Join(name); err != nil {
		t.Errorf("expected non error from Join; got %v", err)
	}
}

func TestWaitFinishWithError(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	s := mock_ticktacktoe_proto.NewMockTickTackToe_GameClient(c)

	g := NewGame(s)

	gomock.InOrder(
		s.EXPECT().Recv().Return(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Result: tpb.Finish_ERROR,
				},
			},
		}, (error)(nil)),
		s.EXPECT().CloseSend().Return((error)(nil)),
	)
	if err := g.Wait(); err == nil {
		t.Error("expected non-nil error from Wait")
	} else if err == ErrGameIsFinished {
		t.Error("expected an error from Wait which is not ErrGameIsFinished")
	}
}

func TestWaitFinish(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	s := mock_ticktacktoe_proto.NewMockTickTackToe_GameClient(c)

	g := NewGame(s)

	gomock.InOrder(
		s.EXPECT().Recv().Return(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Opponent: &tpb.Move{
						Row: 1,
						Col: 2,
					},
				},
			},
		}, (error)(nil)),
		s.EXPECT().CloseSend().Return((error)(nil)),
	)
	if err := g.Wait(); err != ErrGameIsFinished {
		t.Errorf("got an error %v from Wait; want ErrGameIsFinished", err)
	}
	if got := g.Board[1][2]; got != Opponent {
		t.Errorf("Board[1][2] is occupied by %v; want %v", got, Opponent)
	}
}

func TestWaitNotFinished(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	s := mock_ticktacktoe_proto.NewMockTickTackToe_GameClient(c)

	g := NewGame(s)

	s.EXPECT().Recv().Return(&tpb.Response{
		Event: &tpb.Response_MakeMove{
			MakeMove: &tpb.MakeMove{
				Opponent: &tpb.Move{
					Row: 1,
					Col: 2,
				},
			},
		},
	}, (error)(nil))

	if err := g.Wait(); err != nil {
		t.Errorf("got an error %v from Wait; want no error", err)
	}
	if got := g.Board[1][2]; got != Opponent {
		t.Errorf("Board[1][2] is occupied by %v; want %v", got, Opponent)
	}
}

func TestMakeMove(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	s := mock_ticktacktoe_proto.NewMockTickTackToe_GameClient(c)

	g := NewGame(s)

	if err := g.MakeMove(-1, -1); err == nil {
		t.Error("expected an error from MakeMove(-1, -1)")
	}

	s.EXPECT().Send(&tpb.Request{
		Event: &tpb.Request_Move{
			Move: &tpb.Move{
				Row: 2,
				Col: 1,
			},
		},
	}).Return((error)(nil))

	if err := g.MakeMove(2, 1); err != nil {
		t.Errorf("got an error %v from MakeMove; want no error", err)
	}
	if got := g.Board[2][1]; got != Self {
		t.Errorf("Board[2][1] is occupied by %v; want %v", got, Self)
	}
}
