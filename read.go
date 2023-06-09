package main

import (
	"fmt"
	"strings"
)

func read() {
	var (
		b   []byte
		n   int
		err error
	)

	for running {
		b = make([]byte, 1024)

		n, err = conn.Read(b)
		if err != nil {
			panic(fmt.Sprintf("read error: %s\n", err))
		}

		if n > 0 {
			s := strings.Split(string(b), "\n")[0]
			parts := strings.Split(s, "|")
			action := Action(parts[0])
			args := parts[1:]
			handle(action, args...)
		}
	}

	wg.Done()
}

func handle(action Action, args ...string) {
	switch action {
	case actionMotd:
		handleMotd(args[0])
	case actionError:
		handleError(args[0])
	case actionGame:
		handleGame(getInt(args[0]), getInt(args[1]), getInt(args[2]))
	case actionPos:
		handlePos(getInt(args[0]), getInt(args[1]), getInt(args[2]))
	case actionTick:
		handleTick()
	case actionDie:
		ids := make([]int, 0)
		for _, s := range args {
			ids = append(ids, getInt(s))
		}
		handleDie(ids...)
	case actionMessage:
		handleMessage(getInt(args[0]), args[1])
	case actionWin:
		handleWin(getInt(args[0]), getInt(args[1]))
	case actionLose:
		handleLose(getInt(args[0]), getInt(args[1]))
	default:
		fmt.Printf("unknown action: %s\n", action)
	}
}

func handleMotd(message string) {
	fmt.Printf("message of the day: %s\n", message)
}

func handleError(errorMessage string) {
	fmt.Printf("error: %s\n", errorMessage)
}

func handleGame(width, height, id int) {
	state.tick = 0
	state.width = width
	state.height = height
	state.players = make(map[int]Player)
	state.cellMatrix = make([][]CellState, width, height)

	playerId = id

	state.players[id] = Player{
		id: id,
		position: Position{
			x: 0,
			y: 0,
		},
		history:   make([]Position, 0),
		direction: up,
	}
}

func handlePos(id, x, y int) {
	p, found := state.players[id]
	if !found {
		p.setPosition(x, y)
	} else {
		state.addPlayer(id, x, y)
	}
}

func handleTick() {
	move(up)
	// switch state.tick % 4 {
	// case 0:
	// 	move(up)
	// case 1:
	// 	move(left)
	// case 2:
	// 	move(up)
	// case 3:
	// 	move(right)
	// }

	state.addTick()
}

func handleDie(ids ...int) {
	deadPlayersBuilder := strings.Builder{}

	for i, id := range ids {
		if i != 0 {
			deadPlayersBuilder.WriteString(", ")
		}

		deadPlayersBuilder.WriteString(fmt.Sprintf("%d", id))
	}

	fmt.Printf("players died: %s\n", deadPlayersBuilder.String())
}

func handleMessage(id int, message string) {
	fmt.Printf("message from the player '%d': %s\n", id, message)
}

func handleWin(wins, loses int) {
	fmt.Printf("player '%d' wins: %d, loses: %d, ratio: %.3f\n", playerId, wins, loses, float64(wins)/float64(wins+loses))
}

func handleLose(wins, loses int) {
	fmt.Printf("player '%d' wins: %d, loses: %d, ratio: %.3f\n", playerId, wins, loses, float64(wins)/float64(wins+loses))
	running = false
}
