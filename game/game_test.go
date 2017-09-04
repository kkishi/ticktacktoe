package game

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkishi/ticktacktoe/proto/mock_ticktacktoe_proto"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

func TestJoin(t *testing.T) {
	const (
		nameA = "Test Player A"
		nameB = "Test Player B"
	)
	g := New()

	ca := gomock.NewController(t)
	defer ca.Finish()
	sa := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(ca)
	gomock.InOrder(
		sa.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: nameA,
				},
			},
		}, (error)(nil)),
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{
					Initial: true,
				},
			},
		}).Return((error)(nil)),
		sa.EXPECT().Recv().Return(&tpb.Request{}, errors.New("invalid request")),
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Error: true,
				},
			},
		}).Return((error)(nil)),
	)
	g.Join(sa)

	cb := gomock.NewController(t)
	defer cb.Finish()
	sb := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(cb)
	gomock.InOrder(
		sb.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: nameB,
				},
			},
		}, (error)(nil)),
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Error: true,
				},
			},
		}).Return((error)(nil)),
	)
	g.Join(sb)

	g.Start()

	if got := g.Names[PlayerA.ToIndex()]; got != nameA {
		t.Errorf("Got %q for the name of PlayerA; want %q", got, nameA)
	}
	if got := g.Names[PlayerB.ToIndex()]; got != nameB {
		t.Errorf("Got %q for the name of PlayerB; want %q", got, nameB)
	}
}

func TestWaiting(t *testing.T) {
	g := New()
	if !g.Waiting() {
		t.Error("g.Waiting() = false; expected true")
	}
	c := gomock.NewController(t)
	defer c.Finish()
	s := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(c)
	s.EXPECT().Recv().AnyTimes()
	g.Join(s)
	if !g.Waiting() {
		t.Error("g.Waiting() = false; expected true")
	}
	g.Join(s)
	if g.Waiting() {
		t.Error("g.Waiting() = true; expected false")
	}
}

func TestMove(t *testing.T) {
	g := New()

	ca := gomock.NewController(t)
	defer ca.Finish()
	sa := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(ca)
	gomock.InOrder(
		sa.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: "Test Player A",
				},
			},
		}, (error)(nil)),
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{
					Initial: true,
				},
			},
		}).Return((error)(nil)),
		sa.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Move{
				Move: &tpb.Move{
					Row: 0,
					Col: 0,
				},
			},
		}, (error)(nil)),
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{
					Opponent: &tpb.Move{
						Row: 0,
						Col: 1,
					},
				},
			},
		}).Return(errors.New("invalid request")),
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Error: true,
				},
			},
		}).Return((error)(nil)),
	)
	g.Join(sa)

	cb := gomock.NewController(t)
	defer cb.Finish()
	sb := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(cb)
	gomock.InOrder(
		sb.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: "Test Player B",
				},
			},
		}, (error)(nil)),
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{
					Opponent: &tpb.Move{
						Row: 0,
						Col: 0,
					},
				},
			},
		}).Return((error)(nil)),
		sb.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Move{
				Move: &tpb.Move{
					Row: 0,
					Col: 1,
				},
			},
		}, (error)(nil)),
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Error: true,
				},
			},
		}).Return((error)(nil)),
	)
	g.Join(sb)

	g.Start()

	if got := g.Board[0][0]; got != PlayerA {
		t.Errorf("g.Board[0][0] is takey by %v; want PlayerA", got)
	}
	if got := g.Board[0][1]; got != PlayerB {
		t.Errorf("g.Board[0][1] is takey by %v; want PlayerB", got)
	}
}

func TestFinish(t *testing.T) {
	tests := []struct {
		moves [][]*tpb.Move
		want  Player
	}{
		{
			moves: [][]*tpb.Move{
				{{0, 0}, {1, 1}, {2, 2}},
				{{0, 1}, {0, 2}},
			},
			want: PlayerA,
		},
		{
			moves: [][]*tpb.Move{
				{{0, 0}, {1, 1}, {0, 1}},
				{{2, 0}, {2, 1}, {2, 2}},
			},
			want: PlayerB,
		},
	}

	for _, test := range tests {
		g := New()

		var ctrls []*gomock.Controller
		var servers []*mock_ticktacktoe_proto.MockTickTackToe_GameServer
		var calls [][]*gomock.Call

		for i := 0; i < 2; i++ {
			ctrls = append(ctrls, gomock.NewController(t))
			servers = append(servers,
				mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(ctrls[i]))
			calls = append(calls, []*gomock.Call{
				servers[i].EXPECT().Recv().Return(&tpb.Request{
					Event: &tpb.Request_Join{
						Join: &tpb.Join{
							Name: fmt.Sprintf("Test %v", PlayerFromIndex(i)),
						},
					},
				}, (error)(nil)),
			})
		}

		var lastMove *tpb.Move
		player := PlayerA
		turn := 0
		for {
			pi := player.ToIndex()
			moves := test.moves[pi]
			if len(moves) <= turn {
				break
			}
			move := moves[turn]

			calls[pi] = append(calls[pi], servers[pi].EXPECT().Send(&tpb.Response{
				Event: &tpb.Response_MakeMove{
					MakeMove: &tpb.MakeMove{
						Initial:  lastMove == nil,
						Opponent: lastMove,
					},
				},
			}).Return((error)(nil)))

			calls[pi] = append(calls[pi], servers[pi].EXPECT().Recv().Return(
				&tpb.Request{
					Event: &tpb.Request_Move{
						Move: move,
					},
				}, (error)(nil)))

			lastMove = move
			if player == PlayerA {
				player = PlayerB
			} else {
				player = PlayerA
				turn++
			}
		}

		for i := 0; i < 2; i++ {
			f := &tpb.Finish{}
			if test.want == PlayerFromIndex(i) {
				f.Win = true
			} else {
				f.Opponent = lastMove
			}
			calls[i] = append(calls[i], servers[i].EXPECT().Send(&tpb.Response{
				Event: &tpb.Response_Finish{
					Finish: f,
				},
			}).Return((error)(nil)))
			gomock.InOrder(calls[i]...)
		}

		for i := 0; i < 2; i++ {
			g.Join(servers[i])
		}
		g.Start()

		if got := g.Board.WinningPlayer(); got != test.want {
			t.Errorf("WinningPlayer returned %v; want %v", got, test.want)
		}

		for i := 0; i < 2; i++ {
			ctrls[i].Finish()
		}
	}
}
