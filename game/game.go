package game

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Game struct {
	snake *snake

	rows, cols int

	foods map[string]pos

	points int

	flushInterval time.Duration
}

func Create(rows, cols int) *Game {
	return &Game{
		rows: rows,
		cols: cols,
	}
}

func (g *Game) Start() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()
	g.init()
	go g.flush()

loop:
	for {
		ev := termbox.PollEvent()
		switch ev.Key {
		case termbox.KeyEsc:
			break loop

		case termbox.KeyArrowUp:
			g.snake.changeDirect(directUp)

		case termbox.KeyArrowDown:
			g.snake.changeDirect(directDown)

		case termbox.KeyArrowLeft:
			g.snake.changeDirect(directLeft)

		case termbox.KeyArrowRight:
			g.snake.changeDirect(directRight)

		default:
			switch ev.Ch {
			case 'h':
				g.snake.changeDirect(directLeft)

			case 'j':
				g.snake.changeDirect(directDown)

			case 'k':
				g.snake.changeDirect(directUp)

			case 'l':
				g.snake.changeDirect(directRight)
			}
		}
	}
	return nil
}

func (g *Game) init() {
	g.snake = new(snake)
	initX := g.cols / 2
	initY := g.rows / 2
	g.snake.head = &node{
		pos: pos{x: initX, y: initY},
	}
	g.snake.direct = directDown
	g.snake.length = 1
	for i := 0; i < 5; i++ {
		g.snake.grow()
	}
	g.foods = make(map[string]pos, 1)
	g.flushInterval = time.Millisecond * 400
	g.createFood()
}

func (g *Game) flush() {
	g.step()
	for range time.Tick(g.flushInterval) {
		g.step()
		if g.snake.outBorder(g.rows, g.cols) {
			fmt.Println("You lose!")
			break
		}
	}
}

func (g *Game) step() {
	posKey := g.snake.head.pos.String()
	_, eat := g.foods[posKey]
	if eat {
		delete(g.foods, posKey)
		g.snake.grow()
		g.points++
	} else {
		g.snake.move()
	}
	if len(g.foods) == 0 {
		g.createFood()
	}

	foods := make([]pos, 0, len(g.foods))
	for _, fd := range g.foods {
		foods = append(foods, fd)
	}

	fmt.Print("\033[H\033[2J")
	m := draw(g.snake, foods, g.rows, g.cols)
	fmt.Println(m)
	fmt.Println("Use H, J, K, L to move your snake!")
	fmt.Printf("Points: %d\n", g.points)
}

func (g *Game) createFood() {
	snakePosMap := make(map[string]struct{}, g.snake.length)
	cur := g.snake.head
	for cur != nil {
		snakePosMap[cur.pos.String()] = struct{}{}
		cur = cur.pre
	}

	var choices []pos
	for i := 1; i < g.rows-1; i++ {
		for j := 1; j < g.cols-1; j++ {
			p := pos{x: j, y: i}
			_, exists := snakePosMap[p.String()]
			if exists {
				continue
			}
			_, exists = g.foods[p.String()]
			if exists {
				continue
			}
			choices = append(choices, p)
		}
	}
	if len(choices) == 0 {
		os.Exit(0)
		return
	}

	idx := rand.Intn(len(choices))
	food := choices[idx]
	g.foods[food.String()] = food
}
