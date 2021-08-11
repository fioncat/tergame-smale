package game

import (
	"fmt"
	"testing"
)

func TestDraw(t *testing.T) {
	foods := []pos{
		{x: 5, y: 5},
		{x: 19, y: 18},
		{x: 23, y: 13},
	}

	s := new(snake)
	s.head = &node{
		pos: pos{x: 15, y: 15},
	}
	s.direct = directDown

	fmt.Println(draw(s, foods, 20, 40))
	s.grow()
	fmt.Println(draw(s, foods, 20, 40))
	s.move()
	fmt.Println(draw(s, foods, 20, 40))
	s.move()
	fmt.Println(draw(s, foods, 20, 40))
	s.changeDirect(directRight)
	s.grow()
	fmt.Println(draw(s, foods, 20, 40))
	s.move()
	fmt.Println(draw(s, foods, 20, 40))
	s.move()
	fmt.Println(draw(s, foods, 20, 40))
	s.move()
	fmt.Println(draw(s, foods, 20, 40))
	s.grow()
	fmt.Println(draw(s, foods, 20, 40))

}
