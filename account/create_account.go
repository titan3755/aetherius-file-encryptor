package account

import (
	"fmt"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
	"crypto/sha256"
	
)

func CreateAccount() {
	var accName string
	var accPass string
	utils.PtermHeaders("Create Account", pterm.FgLightRed)
	pterm.DefaultParagraph.Println("Welcome to the account or profile creation page! Here you can create a profile in order to use this app properly. Your profile will be saved in the documents folder and it will contain the hashed form of your passwords or private keys and your public keys as well. The password you will provide here will be used as the private key by default unless defaults are overridden.")
	pterm.DefaultParagraph.Println("")
	pterm.DefaultParagraph.Printfln("For enhancing security further, if your profile is somehow lost or deleted, no one will ever be able to access your \"Aetherius\" encrypted files as the public and private keys used to encrypt your files will be lost along with your profile folder. A feature to get back your lost profile and keys will be added to this app some time in the future.")
	pterm.DefaultParagraph.Println("")
	pterm.DefaultBasicText.WithStyle(
		pterm.NewStyle(pterm.FgLightBlue),
	).Print("> Enter your account name ---> ")
	fmt.Scanln(&accName)
	for {
		pterm.DefaultBasicText.WithStyle(
			pterm.NewStyle(pterm.FgLightBlue),
		).Print("> Enter your new password ---> ")
		fmt.Scanln(&accPass)
		result, err := utils.PasswordVerifier(accPass)
		if result {
			break
		} else {
			pterm.DefaultBasicText.WithStyle(
				pterm.NewStyle(pterm.FgRed),
			).Print("!> " + err + " (try again)\n\n")
		}
	}
	encrypt := sha256.Sum256([]byte (accPass))
	pterm.DefaultBasicText.Println("\nYou typed: " + accName + "\nPassHash: " + fmt.Sprintf("%X", encrypt) + "\nPassword: " + accPass)
}

func directoryCreator() {
	// to do
}

func keyCreator() {
	// to do
}