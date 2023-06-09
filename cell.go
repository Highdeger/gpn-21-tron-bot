package main

type CellState int

const (
	Empty CellState = iota
	Blocked
	PossibleBlockage
)
