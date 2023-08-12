package utils

import "github.com/pterm/pterm"

func ErrorMsg(msg string) {
	pterm.DefaultSection.Println("An error has occured!")
	pterm.Error.Println(msg)
}