# langtons-ant

[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/langtons-ant#license-mit)

[Langton's ant](https://en.wikipedia.org/wiki/Langton's_ant) written in Go using [termbox-go](https://github.com/nsf/termbox-go).

![Animation](http://assets.c7.se/langtons-ant/animation.gif)

**Attractor of Langton's ant (The Higway)**

![The Highway](http://assets.c7.se/langtons-ant/highway.png)

## Description

Langton's ant is a two-dimensional Turing machine with a very simple set of rules but complex emergent behavior.

## Rules

 - At a white square, turn 90° right, flip the color of the square, move forward one unit
 - At a black square, turn 90° left, flip the color of the square, move forward one unit

## Installation

		go get -u github.com/peterhellberg/langtons-ant

## License (MIT)

Copyright (c) 2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
