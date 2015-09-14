// Langton's ant is a two-dimensional Turing machine with a very simple set of rules but complex emergent behavior.
package main

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	Up    = '▲'
	Right = '▶'
	Down  = '▼'
	Left  = '◀'
)

var (
	ant   *Ant
	board *image.Gray16
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

func render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y := 0; y < ant.H; y++ {
		for x := 0; x < ant.W; x++ {
			fg := termbox.ColorDefault
			ch := ' '

			if ant.X == x && ant.Y == y {
				fg = termbox.ColorRed
				ch = ant.D
			}

			if board.At(x, y) == color.Black {
				termbox.SetCell(x, y, ch, fg, termbox.ColorBlack)
			} else {
				termbox.SetCell(x, y, ch, fg, termbox.ColorDefault)
			}
		}
	}

	termbox.Flush()
}

func setup() (*Ant, *image.Gray16) {
	w, h := termbox.Size()
	ant := &Ant{X: w / 2, Y: h / 2, W: w, H: h, D: Up}

	board = image.NewGray16(image.Rect(0, 0, ant.W, ant.H))
	draw.Draw(board, board.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	return ant, board
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	ant, board = setup()

	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	go func() {
		for {
			ant.Turn()
			ant.FlipColor()
			ant.Move()
			time.Sleep(33 * time.Millisecond)
		}
	}()

	render()
loop:
	for {
		select {
		case ev := <-events:
			if ev.Type == termbox.EventKey {
				break loop
			}

			if ev.Type == termbox.EventResize {
				ant, board = setup()
			}
		default:
			render()
			time.Sleep(33 * time.Millisecond)
		}
	}
}
