package userpages

import (
	"fmt"

	"github.com/pterm/pterm"
)

func UserHome(accName string) {
	pterm.Success.Println("Login has been successful!\nLogged in to: " + accName)
	fmt.Println("Welcome to user home page!!")
}