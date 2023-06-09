package main

type Action string

const (
	// server actions
	actionMotd    Action = "motd"
	actionError   Action = "error"
	actionGame    Action = "game"
	actionPos     Action = "pos"
	actionTick    Action = "tick"
	actionDie     Action = "die"
	actionMessage Action = "message"
	actionWin     Action = "win"
	actionLose    Action = "lose"

	// client actions
	actionJoin Action = "join"
	actionMove Action = "move"
	actionChat Action = "chat"
)

func (r Action) String() string {
	return string(r)
}
