package main

import "fmt"

type Player struct {
	id        int
	position  *Position
	history   []*Position
	direction Direction
}

func (p *Player) setPosition(x, y int) {
	p.history = append(p.history, &Position{p.position.x, p.position.y})
	p.position.set(x, y)
}

func (p *Player) forward() Direction {
	return p.direction
}

func (p *Player) left() Direction {
	switch p.direction {
	case up:
		return left
	case down:
		return right
	case left:
		return down
	case right:
		return up
	default:
		panic(fmt.Sprintf("invalid direction: %s", p.direction))
	}
}

func (p *Player) right() Direction {
	switch p.direction {
	case up:
		return right
	case down:
		return left
	case left:
		return up
	case right:
		return down
	default:
		panic(fmt.Sprintf("invalid direction: %s", p.direction))
	}
}
