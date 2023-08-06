package account

import (
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
)

func LoginAccount() {
	utils.PtermHeaders("Login", pterm.FgLightMagenta)
	pterm.DefaultParagraph.Println("Welcome to the account login page! Here you can login to your account or profile which you created previously via the account creation page.")
	pterm.DefaultParagraph.Println("")
	
}
