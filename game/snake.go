package game

import (
	"fmt"
	"sync"
)

const (
	directUp = iota
	directDown
	directLeft
	directRight
)

type pos struct {
	x, y int
}

func (p pos) String() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

type node struct {
	pos
	next *node
	pre  *node
}

type snake struct {
	head   *node
	direct int
	mu     sync.Mutex
	length int
}

func (s *snake) grow() {
	newNode := new(node)
	newNode.pos = s.head.pos
	s.moveNode(newNode)
	s.head.next = newNode
	newNode.pre = s.head
	s.head = newNode
	s.length++
}

func (s *snake) move() {
	tail := s.head
	for tail.pre != nil {
		tail = tail.pre
	}
	cur := tail
	for cur != nil {
		cn := cur.next
		if cn != nil {
			cur.pos = cn.pos
		}
		cur = cn
	}
	s.moveNode(s.head)
}

func (s *snake) outBorder(rows, cols int) bool {
	if s.head.x <= 0 || s.head.y <= 0 {
		return true
	}
	if s.head.x >= cols-1 || s.head.y >= rows-1 {
		return true
	}
	headp := s.head.pos
	cur := s.head.pre
	for cur != nil {
		if cur.pos.x == headp.x && cur.pos.y == headp.y {
			return true
		}
		cur = cur.pre
	}
	return false
}

func (s *snake) moveNode(n *node) {
	s.mu.Lock()
	defer s.mu.Unlock()
	switch s.direct {
	case directUp:
		n.pos.y--
	case directDown:
		n.pos.y++
	case directLeft:
		n.pos.x--
	case directRight:
		n.pos.x++
	}
}

func (s *snake) changeDirect(d int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.direct == directUp && d == directDown {
		return
	}
	if s.direct == directDown && d == directUp {
		return
	}
	if s.direct == directRight && d == directLeft {
		return
	}
	if s.direct == directLeft && d == directRight {
		return
	}
	s.direct = d
}
