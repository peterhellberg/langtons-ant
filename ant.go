package main

import "image/color"

const (
	Up    = '▲'
	Right = '▶'
	Down  = '▼'
	Left  = '◀'
)

type Ant struct {
	X int
	Y int
	W int
	H int
	D rune
}

func (a *Ant) OnWhiteSquare() bool {
	return board.At(a.X, a.Y) == color.White
}

func (a *Ant) Turn() {
	if a.OnWhiteSquare() {
		a.turnRight()
	} else {
		a.turnLeft()
	}
}

func (a *Ant) turnRight() {
	switch a.D {
	case Up:
		a.D = Right
	case Right:
		a.D = Down
	case Down:
		a.D = Left
	case Left:
		a.D = Up
	}
}

func (a *Ant) turnLeft() {
	switch a.D {
	case Up:
		a.D = Left
	case Left:
		a.D = Down
	case Down:
		a.D = Right
	case Right:
		a.D = Up
	}
}

func (a *Ant) FlipColor() {
	if a.OnWhiteSquare() {
		board.Set(a.X, a.Y, color.Black)
	} else {
		board.Set(a.X, a.Y, color.White)
	}
}

func (a *Ant) Move() {
	switch a.D {
	case Up:
		a.Y--
	case Right:
		a.X++
	case Down:
		a.Y++
	case Left:
		a.X--
	}
}
