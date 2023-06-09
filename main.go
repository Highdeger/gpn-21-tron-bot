package main

import (
	"net"
	"sync"
)

var (
	conn     net.Conn
	err      error
	state    State
	running  bool
	playerId int
	wg       *sync.WaitGroup
)

func main() {
	wg.Add(1)
	go read()

	go state.printStatus()

	join("tiwaz", "92bjaHWZxb66gbfeB!*69FNw3Ynm^f@B^4zhmCL%&u")

	wg.Wait()

	err = conn.Close()
	if err != nil {
		panic(err)
	}
}

func init() {
	conn, err = net.Dial("tcp", "gpn-tron.duckdns.org:4000")
	if err != nil {
		panic(err)
	}

	state = State{
		tick:       0,
		width:      0,
		height:     0,
		players:    make(map[int]Player),
		cellMatrix: make([][]CellState, 0, 0),
	}

	running = true
	playerId = -1
	wg = &sync.WaitGroup{}
}
