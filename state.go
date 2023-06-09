package main

import (
	"fmt"
	"strings"
	"time"
)

type State struct {
	tick       int
	width      int
	height     int
	players    map[int]*Player
	cellMatrix [][]CellState
}

func (r *State) addPlayer(id, x, y int) {
	r.players[id] = &Player{
		id: id,
		position: &Position{
			x: x,
			y: y,
		},
		history:   make([]*Position, 0),
		direction: up,
	}
}

func (r *State) addTick() {
	r.tick++
}

func (r *State) printStatus() {
	for {
		playersInfoBuilder := strings.Builder{}

		for i, player := range r.players {
			if i != 0 {
				playersInfoBuilder.WriteString("|")
			}

			playersInfoBuilder.WriteString(fmt.Sprintf("%d>>%d:%d", player.id, player.position.x, player.position.y))
		}

		playersInfoBuilder.WriteString("\n")

		fmt.Printf("ticks: %d, players: %s\n", r.tick, playersInfoBuilder.String())

		time.Sleep(1 * time.Second)
	}
}
