package main

func zigzag() {
	switch state.tick % 4 {
	case 0:
		move(up)
	case 1:
		move(left)
	case 2:
		move(up)
	case 3:
		move(right)
	}
}
