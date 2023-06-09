package main

type Direction string

const (
	up    Direction = "up"
	down  Direction = "down"
	left  Direction = "left"
	right Direction = "right"
)

func (r Direction) String() string {
	return string(r)
}
