// Langton's ant is a two-dimensional Turing machine with a very simple set of rules but complex emergent behavior.
package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	ant *Ant

	delay = flag.Duration("delay", 33*time.Millisecond, "The delay in the logic loop")
	right = flag.Bool("right", false, "Right direction at start")
	left  = flag.Bool("left", false, "Left direction at start")
	up    = flag.Bool("up", false, "Up direction at start")
	down  = flag.Bool("down", false, "Down direction at start")

	direction = Down
)

func init() {
	flag.Parse()

	if *up {
		direction = Up
	}

	if *down {
		direction = Down
	}

	if *right {
		direction = Right
	}

	if *left {
		direction = Left
	}
}

func setup() {
	w, h := termbox.Size()

	board := image.NewGray16(image.Rect(0, 0, w, h))
	draw.Draw(board, board.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	ant = &Ant{X: w / 2, Y: h / 2, W: w, H: h, D: direction, B: board}
}

func logic() {
	for {
		ant.Turn()
		time.Sleep(*delay)
		ant.FlipColor()
		ant.Move()
	}
}

func render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	ant.RLock()
	antW := ant.W
	antH := ant.H
	ant.RUnlock()

	for y := 0; y < antH; y++ {
		for x := 0; x < antW; x++ {
			fg := termbox.ColorDefault
			ch := ' '

			ant.RLock()
			antX := ant.X
			antY := ant.Y
			antD := ant.D
			antColor := ant.B.At(x, y)
			ant.RUnlock()

			if antX == x && antY == y {
				fg = termbox.ColorRed
				ch = antD
			}

			if antColor == color.Black {
				termbox.SetCell(x, y, ch, fg, termbox.ColorBlack)
			} else {
				termbox.SetCell(x, y, ch, fg, termbox.ColorDefault)
			}
		}
	}

	termbox.Flush()
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	setup()

	go logic()

	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()
loop:
	for {
		select {
		case ev := <-events:
			if ev.Type == termbox.EventKey {
				break loop
			}

			if ev.Type == termbox.EventResize {
				setup()
			}
		default:
			render()
			time.Sleep(33 * time.Millisecond)
		}
	}
}
