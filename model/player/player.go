package player

import "fmt"

type Player int

const (
	Unknown Player = iota
	A
	B
)

func FromIndex(i int) Player {
	if i == 0 {
		return A
	}
	return B
}

func (p Player) ToIndex() int {
	if p == A {
		return 0
	}
	return 1
}

func (p Player) String() string {
	switch p {
	case Unknown:
		return "Unknown player"
	case A:
		return "Player A"
	case B:
		return "Player B"
	}
	return fmt.Sprintf("Player(%d)", p)
}

func (p Player) Next() Player {
	if p == A {
		return B
	}
	return A
}
