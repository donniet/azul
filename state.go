package main

import "fmt"

type Tile int
type Pattern []Tile // cannot contain FirstPlayer
type Window []Tile  // cannot contain FirstPlayer
type Factory []Tile // cannot contain Empty or FirstPlayer
type Bag []Tile     // cannot contain Empty or FirstPlayer
type Pot []Tile     // cannot contain Empty
type BrokenTiles []Tile

type PlayerBoard struct {
	Patterns []Pattern
	Window
	BrokenTiles
	Score int
}

type State struct {
	PlayerBoards []PlayerBoard
	Factories    []Factory
	Pot
	Bag
}

const (
	Empty Tile = iota
	FirstPlayer
	Blue
	Orange
	Red
	Black
	White
)

func (t Tile) String() string {
	switch t {
	case Blue:
		return "blue"
	case Orange:
		return "orange"
	case Red:
		return "red"
	case Black:
		return "black"
	case White:
		return "white"
	case FirstPlayer:
		return "first-player"
	default:
		return ""
	}
}

func (w Window) At(col, row int) Tile {
	return w[5*row+col]
}

func (w Window) Set(col, row int, tile Tile) {
	w[5*row+col] = tile
}

func NewWindow() Window {
	return make([]Tile, 5*5)
}

func (f Factory) Count(tile Tile) int {
	count := 0

	for i := 0; i < len(f); i++ {
		if f[i] == tile {
			count++
		}
	}

	return count
}
func (f *Factory) Take(tile Tile) bool {
	if f.Count(tile) == 0 {
		return false
	}

	for i := 0; i < len(*f); i++ {
		if (*f)[i] != tile {
			continue
		}

		*f = append((*f)[:i], (*f)[i+1:]...)
	}

	return true
}
func (f Factory) Put(tile []Tile) {}
func (f Factory) At(index int) Tile {
	if index < 0 || index >= len(f) {
		return Empty
	}
	return f[index]
}

func main() {
	fmt.Printf("The Blue tile is: %s\n", Blue)
	fmt.Printf("The Red tile is: %s\n", Red)
	fmt.Printf("The FirstPlayer tile is: %s\n", FirstPlayer)

	window := NewWindow()
	window.Set(2, 2, Blue)

	fmt.Printf("The tile at (2,2) is: %s\n", window.At(2, 2))
}
