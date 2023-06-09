package main

type Position struct {
	x int
	y int
}

func (r *Position) set(x, y int) {
	r.x = x
	r.y = y
}

func (r *Position) clone() *Position {
	return &Position{r.x, r.y}
}
