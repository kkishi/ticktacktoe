package game

import "fmt"

type Player int

const (
	UnknownPlayer Player = iota
	PlayerA
	PlayerB
)

func PlayerFromIndex(i int) Player {
	if i == 0 {
		return PlayerA
	}
	return PlayerB
}

func (p Player) ToIndex() int {
	if p == PlayerA {
		return 0
	}
	return 1
}

func (p Player) String() string {
	switch p {
	case UnknownPlayer:
		return "UnknownPlayer"
	case PlayerA:
		return "PlayerA"
	case PlayerB:
		return "PlayerB"
	}
	return fmt.Sprintf("Player(%d)", p)
}

func (p Player) Next() Player {
	if p == PlayerA {
		return PlayerB
	}
	return PlayerA
}
