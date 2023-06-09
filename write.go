package main

import "fmt"

func write(action Action, params ...string) {
	packet := make([]byte, 0)
	packet = append(packet, []byte(fmt.Sprintf("%s", action.String()))...)

	for _, param := range params {
		packet = append(packet, []byte(fmt.Sprintf("|%s", param))...)
	}

	packet = append(packet, []byte("\n")...)

	_, err = conn.Write(packet)
	if err != nil {
		panic(fmt.Sprintf("write error: %s\n", err))
	}
}

func join(username, password string) {
	write(actionJoin, username, password)
}

func move(direction Direction) {
	write(actionMove, direction.String())
}

func chat(message string) {
	write(actionChat, message)
}
