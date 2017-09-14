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
					Result: tpb.Finish_ERROR,
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
		// Expect a Join request.
		sa.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: "Test Player A",
				},
			},
		}, (error)(nil)),
		// Let PlayerA make the first move.
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{},
			},
		}).Return((error)(nil)),
		// PlayerA takes (0,0).
		sa.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Move{
				Move: &tpb.Move{
					Row: 0,
					Col: 0,
				},
			},
		}, (error)(nil)),
		// PlayerA gets notification for their own move.
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Update{
				Update: &tpb.Update{
					Row:    0,
					Col:    0,
					Player: tpb.Player_A,
				},
			},
		}).Return((error)(nil)),
		// PlayerA gets notification for PlayerB's move.
		sa.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Update{
				Update: &tpb.Update{
					Row:    0,
					Col:    1,
					Player: tpb.Player_B,
				},
			},
		}).Return((error)(nil)),
		// Let PlayerA make the second move, then it returns an error, which leads
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
	g.Join(sa)

	cb := gomock.NewController(t)
	defer cb.Finish()
	sb := mock_ticktacktoe_proto.NewMockTickTackToe_GameServer(cb)
	gomock.InOrder(
		// Expect a Join request.
		sb.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Join{
				Join: &tpb.Join{
					Name: "Test Player B",
				},
			},
		}, (error)(nil)),
		// PlayerB gets notification for PlayerA's move.
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_Update{
				Update: &tpb.Update{
					Row:    0,
					Col:    0,
					Player: tpb.Player_A,
				},
			},
		}).Return((error)(nil)),
		// Let PlayerB make the first move.
		sb.EXPECT().Send(&tpb.Response{
			Event: &tpb.Response_MakeMove{
				MakeMove: &tpb.MakeMove{},
			},
		}).Return((error)(nil)),
		// PlayerB takes (0,1).
		sb.EXPECT().Recv().Return(&tpb.Request{
			Event: &tpb.Request_Move{
				Move: &tpb.Move{
					Row: 0,
					Col: 1,
				},
			},
		}, (error)(nil)),
		// PlayerB gets notification for their own move.
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
	g.Join(sb)

	g.Start()

	if got := g.Board.Grid[0][0]; got != PlayerA {
		t.Errorf("g.Board.Grid[0][0] is takey by %v; want PlayerA", got)
	}
	if got := g.Board.Grid[0][1]; got != PlayerB {
		t.Errorf("g.Board.Grid[0][1] is takey by %v; want PlayerB", got)
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
		{
			moves: [][]*tpb.Move{
				{{0, 0}, {0, 2}, {1, 1}, {1, 2}, {2, 1}},
				{{0, 1}, {1, 0}, {2, 0}, {2, 2}},
			},
			want: UnknownPlayer,
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

		player := PlayerA
		turn := 0
		for {
			pi := player.ToIndex()
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
							Player: tpb.Player(player),
						},
					},
				}).Return((error)(nil)))
			}

			if player == PlayerA {
				player = PlayerB
			} else {
				player = PlayerA
				turn++
			}
		}

		for i := 0; i < 2; i++ {
			var r tpb.Finish_Result
			if test.want == UnknownPlayer {
				r = tpb.Finish_DRAW
			} else if PlayerFromIndex(i) == test.want {
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

		for i := 0; i < 2; i++ {
			g.Join(servers[i])
		}
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
