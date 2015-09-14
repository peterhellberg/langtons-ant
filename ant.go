package main

import (
	"image"
	"image/color"
	"sync"
)

const (
	Up    = '▲'
	Right = '▶'
	Down  = '▼'
	Left  = '◀'
)

type Ant struct {
	sync.RWMutex
	X int
	Y int
	W int
	H int
	D rune
	B *image.Gray16
}

func (a *Ant) OnWhiteSquare() bool {
	a.RLock()
	defer a.RUnlock()

	return a.B.At(a.X, a.Y) == color.White
}

func (a *Ant) Turn() {
	a.RLock()
	d := a.D
	a.RUnlock()

	if a.OnWhiteSquare() {
		a.Lock()
		switch d {
		case Up:
			a.D = Right
		case Right:
			a.D = Down
		case Down:
			a.D = Left
		case Left:
			a.D = Up
		}
		a.Unlock()
	} else {
		a.Lock()
		switch d {
		case Up:
			a.D = Left
		case Left:
			a.D = Down
		case Down:
			a.D = Right
		case Right:
			a.D = Up
		}
		a.Unlock()
	}
}

func (a *Ant) FlipColor() {
	if a.OnWhiteSquare() {
		a.Lock()
		a.B.Set(a.X, a.Y, color.Black)
		a.Unlock()
	} else {
		a.Lock()
		a.B.Set(a.X, a.Y, color.White)
		a.Unlock()
	}
}

func (a *Ant) Move() {
	a.Lock()
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
	a.Unlock()
}
