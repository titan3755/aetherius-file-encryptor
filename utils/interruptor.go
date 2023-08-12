package utils

import (
	"fmt"
	"strings"
	"github.com/pterm/pterm"
)

func Interruptor() bool {
	var userAnswer string
	var runState bool
	pterm.DefaultBasicText.WithStyle(
		pterm.NewStyle(pterm.FgRed),
	).Print("\n\n> Yes/No? (y/n) ---> ")
	fmt.Scanln(&userAnswer)
	if strings.ToLower(userAnswer) == "y" {
		runState = true
	} else {
		runState = false
	}
	return runState
}