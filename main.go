package main

import (
	"fmt"

	"github.com/fioncat/tergame-snake/game"
)

func main() {
	g := game.Create(20, 40)
	err := g.Start()
	if err != nil {
		fmt.Printf("start game failed: %v\n", err)
	}
}
