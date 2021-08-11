package game

import "strings"

const (
	headCh = '@'
	bodyCh = '#'
	foodCh = '*'

	hbCh = '-'
	vbCh = '|'
)

func draw(s *snake, foods []pos, rows, cols int) string {
	area := make([][]byte, rows)

	// draw border.
	for i := 0; i < rows; i++ {
		area[i] = make([]byte, cols)
		if i == 0 || i == rows-1 {
			for j := 0; j < cols; j++ {
				area[i][j] = hbCh
			}
		} else {
			for j := 0; j < cols; j++ {
				if j == 0 || j == cols-1 {
					area[i][j] = vbCh
				} else {
					area[i][j] = ' '
				}
			}
		}
	}

	// draw foods.
	for _, fd := range foods {
		if posValid(fd, rows, cols) {
			area[fd.y][fd.x] = foodCh
		}
	}

	// draw snake.
	cur := s.head
	isHead := true
	for cur != nil {
		var ch byte
		ch = bodyCh
		if isHead {
			ch = headCh
			isHead = false
		}
		if posValid(cur.pos, rows, cols) {
			area[cur.pos.y][cur.pos.x] = ch
		}
		cur = cur.pre
	}

	lines := make([]string, rows)
	for i := 0; i < rows; i++ {
		lines[i] = string(area[i])
	}
	return strings.Join(lines, "\n")
}

func posValid(p pos, rows, cols int) bool {
	if p.x <= 0 || p.y <= 0 {
		return false
	}
	if p.x >= cols-1 || p.y >= rows-1 {
		return false
	}
	return true
}
