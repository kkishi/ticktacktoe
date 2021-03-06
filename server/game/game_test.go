package game

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kkishi/ticktacktoe/model/player"
	"github.com/kkishi/ticktacktoe/proto/mock_ticktacktoe_proto"
	"github.com/kkishi/ticktacktoe/server/client"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

type infoMatcher struct {
}

func (m *infoMatcher) Matches(x interface{}) bool {
	r, ok := x.(*tpb.Response)
	return ok && r.GetInfo() != nil
}

func (m *infoMatcher) String() string {
	return "InfoMatcher"
}

var _ gomock.Matcher = (*infoMatcher)(nil)

func TestJoin(t *testing.T) {
	const (
		nameA = "Test Player A"
		nameB = "Test Player B"
	)

	ca := gomock.NewController(t)
	defer ca.Finish()
	sa := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(ca)
	sa.EXPECT().Send(&infoMatcher{}).AnyTimes()
	sa.EXPECT().Context().AnyTimes().Return(context.Background())
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
				MakeMove: &tpb.MakeMove{},
			},
		}).Return((error)(nil)),
		sa.EXPECT().Recv().Return(&tpb.Request{}, errors.New("invalid request")),
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Result: tpb.Finish_ERROR,
				},
			},
		}).Return((error)(nil)),
	)

	cb := gomock.NewController(t)
	defer cb.Finish()
	sb := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(cb)
	sb.EXPECT().Send(&infoMatcher{}).AnyTimes()
	sb.EXPECT().Context().AnyTimes().Return(context.Background())
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
					Result: tpb.Finish_ERROR,
				},
			},
		}).Return((error)(nil)),
	)

	g := New(client.New(sa), client.New(sb))
	g.Start()

	if got := g.Names[player.A.ToIndex()]; got != nameA {
		t.Errorf("Got %q for the name of Player A; want %q", got, nameA)
	}
	if got := g.Names[player.B.ToIndex()]; got != nameB {
		t.Errorf("Got %q for the name of Player B; want %q", got, nameB)
	}
}

func TestMove(t *testing.T) {
	ca := gomock.NewController(t)
	defer ca.Finish()
	sa := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(ca)
	sa.EXPECT().Send(&infoMatcher{}).AnyTimes()
	sa.EXPECT().Context().AnyTimes().Return(context.Background())
	gomock.InOrder(
		// Expect a Join request.
		sa.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: "Test Player A",
				},
			},
		}, (error)(nil)),
		// Let Player A make the first move.
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{},
			},
		}).Return((error)(nil)),
		// Player A takes (0,0).
		sa.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Move{
				Move: &tpb.Move{
					Row: 0,
					Col: 0,
				},
			},
		}, (error)(nil)),
		// Player A gets notification for their own move.
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Update{
				Update: &tpb.Update{
					Row:    0,
					Col:    0,
					Player: tpb.Player_A,
				},
			},
		}).Return((error)(nil)),
		// Player A gets notification for Player B's move.
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Update{
				Update: &tpb.Update{
					Row:    0,
					Col:    1,
					Player: tpb.Player_B,
				},
			},
		}).Return((error)(nil)),
		// Let Player A make the second move, then it returns an error, which leads
		// to the game to finish.
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{},
			},
		}).Return(errors.New("invalid")),
		// Game finishes with an error.
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Result: tpb.Finish_ERROR,
				},
			},
		}).Return((error)(nil)),
	)

	cb := gomock.NewController(t)
	defer cb.Finish()
	sb := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(cb)
	sb.EXPECT().Send(&infoMatcher{}).AnyTimes()
	sb.EXPECT().Context().AnyTimes().Return(context.Background())
	gomock.InOrder(
		// Expect a Join request.
		sb.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: "Test Player B",
				},
			},
		}, (error)(nil)),
		// Player B gets notification for Player A's move.
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Update{
				Update: &tpb.Update{
					Row:    0,
					Col:    0,
					Player: tpb.Player_A,
				},
			},
		}).Return((error)(nil)),
		// Let Player B make the first move.
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{},
			},
		}).Return((error)(nil)),
		// Player B takes (0,1).
		sb.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Move{
				Move: &tpb.Move{
					Row: 0,
					Col: 1,
				},
			},
		}, (error)(nil)),
		// Player B gets notification for their own move.
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Update{
				Update: &tpb.Update{
					Row:    0,
					Col:    1,
					Player: tpb.Player_B,
				},
			},
		}).Return((error)(nil)),
		// Game finishes with an error.
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Finish{
				Finish: &tpb.Finish{
					Result: tpb.Finish_ERROR,
				},
			},
		}).Return((error)(nil)),
	)

	g := New(client.New(sa), client.New(sb))
	g.Start()

	if got := g.Board.Grid[0][0]; got != player.A {
		t.Errorf("g.Board.Grid[0][0] is takey by %v; want Player A", got)
	}
	if got := g.Board.Grid[0][1]; got != player.B {
		t.Errorf("g.Board.Grid[0][1] is takey by %v; want Player B", got)
	}
}

func TestFinish(t *testing.T) {
	tests := []struct {
		moves [][]*tpb.Move
		want  player.Player
	}{
		{
			moves: [][]*tpb.Move{
				{{0, 0}, {1, 1}, {2, 2}},
				{{0, 1}, {0, 2}},
			},
			want: player.A,
		},
		{
			moves: [][]*tpb.Move{
				{{0, 0}, {1, 1}, {0, 1}},
				{{2, 0}, {2, 1}, {2, 2}},
			},
			want: player.B,
		},
		{
			moves: [][]*tpb.Move{
				{{0, 0}, {0, 2}, {1, 1}, {1, 2}, {2, 1}},
				{{0, 1}, {1, 0}, {2, 0}, {2, 2}},
			},
			want: player.Unknown,
		},
	}

	for _, test := range tests {
		var ctrls []*gomock.Controller
		var servers []*mock_ticktacktoe_proto.MockTickTackToe_GameServer
		var calls [][]*gomock.Call

		for i := 0; i < 2; i++ {
			ctrls = append(ctrls, gomock.NewController(t))
			servers = append(servers,
				mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(ctrls[i]))
			servers[i].EXPECT().Send(&infoMatcher{}).AnyTimes()
			servers[i].EXPECT().Context().AnyTimes().Return(context.Background())
			calls = append(calls, []*gomock.Call{
				servers[i].EXPECT().Recv().Return(&tpb.Request{
					Event: &tpb.Request_Join{
						Join: &tpb.Join{
							Name: fmt.Sprintf("Test %v", player.FromIndex(i)),
						},
					},
				}, (error)(nil)),
			})
		}

		p := player.A
		turn := 0
		for {
			pi := p.ToIndex()
			moves := test.moves[pi]
			if len(moves) <= turn {
				break
			}
			move := moves[turn]

			// Let the player make a move.
			calls[pi] = append(calls[pi], servers[pi].EXPECT().Send(&tpb.Response{
				Event: &tpb.Response_MakeMove{
					MakeMove: &tpb.MakeMove{},
				},
			}).Return((error)(nil)))
			calls[pi] = append(calls[pi], servers[pi].EXPECT().Recv().Return(
				&tpb.Request{
					Event: &tpb.Request_Move{
						Move: move,
					},
				}, (error)(nil)))

			// The move is notified.
			for i := 0; i < 2; i++ {
				calls[i] = append(calls[i], servers[i].EXPECT().Send(&tpb.Response{
					Event: &tpb.Response_Update{
						Update: &tpb.Update{
							Row:    move.Row,
							Col:    move.Col,
							Player: tpb.Player(p),
						},
					},
				}).Return((error)(nil)))
			}

			if p == player.A {
				p = player.B
			} else {
				p = player.A
				turn++
			}
		}

		for i := 0; i < 2; i++ {
			var r tpb.Finish_Result
			if test.want == player.Unknown {
				r = tpb.Finish_DRAW
			} else if player.FromIndex(i) == test.want {
				r = tpb.Finish_WIN
			} else {
				r = tpb.Finish_LOSE
			}
			calls[i] = append(calls[i], servers[i].EXPECT().Send(&tpb.Response{
				Event: &tpb.Response_Finish{
					Finish: &tpb.Finish{
						Result: r,
					},
				},
			}).Return((error)(nil)))
			gomock.InOrder(calls[i]...)
		}

		g := New(client.New(servers[0]), client.New(servers[1]))
		g.Start()

		if finished, got := g.Board.Finished(); !finished {
			t.Error("Finished returned that the game is not finished; expected finished")
		} else if got != test.want {
			t.Errorf("Finished returned %v as the winning player; want %v", got, test.want)
		}

		for i := 0; i < 2; i++ {
			ctrls[i].Finish()
		}
	}
}
