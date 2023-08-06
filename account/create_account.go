package account

import (
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
)

func CreateAccount() {
	utils.PtermHeaders("Create Account", pterm.FgLightRed)
	pterm.DefaultParagraph.Println("Welcome to the account or profile creation page! Here you can create a profile in order to use this app properly. Your profile will be saved in the documents folder and it will contain the hashed form of your passwords or private keys and your public keys as well.")
	pterm.DefaultParagraph.Println("")
	pterm.DefaultParagraph.Printfln("For enhancing security further, if your profile is somehow lost or deleted, no one will ever be able to access your \"Aetherius\" encrypted files as the public and private keys used to encrypt your files will be lost along with your profile folder. A feature to get back your lost profile and keys will be added to this app some time in the future.")
	pterm.DefaultParagraph.Println("")
	
}