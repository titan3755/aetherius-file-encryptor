package utils

import (
	"github.com/inancgumus/screen"
)

func ResetTerminal() {
	screen.Clear()
	screen.MoveTopLeft()
}