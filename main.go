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
	ant   *Ant
	board *image.Gray16

	delay = flag.Duration("delay", 33*time.Millisecond, "The delay in the logic loop")
	right = flag.Bool("right", false, "Right direction at start")
	left  = flag.Bool("left", false, "Left direction at start")
	up    = flag.Bool("up", false, "Up direction at start")
	down  = flag.Bool("down", false, "Down direction at start")

	direction = Down
)

func setup() (*Ant, *image.Gray16) {
	w, h := termbox.Size()
	ant := &Ant{X: w / 2, Y: h / 2, W: w, H: h, D: direction}

	board = image.NewGray16(image.Rect(0, 0, ant.W, ant.H))
	draw.Draw(board, board.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	return ant, board
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

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	ant, board = setup()

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
				ant, board = setup()
			}
		default:
			render()
			time.Sleep(33 * time.Millisecond)
		}
	}
}
