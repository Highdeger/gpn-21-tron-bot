package main

import (
	"errors"
	"fmt"
	"io"
	"os"
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
		if err != nil && !errors.Is(err, io.EOF) {
			panic(fmt.Sprintf("read error: %s\n", err))
		}

		if n > 0 {
			s := string(b)
			s = strings.TrimSpace(s)
			s = strings.Trim(s, "\n\x00")

			// fmt.Printf("READ (%d) (contains newline: %t) => %s\n", len(s), strings.Contains(s, "\n"), s)

			splited := strings.Split(s, "|")
			action := splited[0]
			args := strings.Join(splited[1:], "|")

			handle(Action(action), args)

			// for _, line := range strings.Split(s, "\n") {
			// 	line = strings.TrimSpace(line)

			// 	// fmt.Printf("LINE => %s\n", string(b))

			// 	parts := strings.Split(line, "|")
			// 	action := Action(parts[0])
			// 	args := parts[1:]
			// 	handle(action, args...)
			// }
		}
	}

	wg.Done()
}

func handle(action Action, rawArgs string) {
	splited := strings.Split(rawArgs, "|")

	switch action {
	case actionMotd:
		handleMotd(splited[0])
	case actionError:
		handleError(splited[0])
	case actionGame:
		handleGame(getInt(splited[0]), getInt(splited[1]), getInt(splited[2]))
	case actionPos:
		lines := strings.Split(rawArgs, "\n")
		for _, line := range lines {
			fmt.Printf("line: %s\n", line)
			splited := strings.Split(line, "|")
			handlePos(getInt(splited[1]), getInt(splited[2]), getInt(splited[3]))
		}
	case actionTick:
		handleTick()
	case actionDie:
		ids := make([]int, 0)
		for _, s := range splited {
			ids = append(ids, getInt(s))
		}
		handleDie(ids...)
	case actionMessage:
		handleMessage(getInt(splited[0]), splited[1])
	case actionWin:
		handleWin(getInt(splited[0]), getInt(splited[1]))
	case actionLose:
		handleLose(getInt(splited[0]), getInt(splited[1]))
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
	switch os.Args[1] {
	case "zigzag":
		zigzag()
	default:
		move(up)
	}

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
