package userpages

import (
	"fmt"

	"example.com/encrypt_tui/decryptpage"
	"example.com/encrypt_tui/encryptpage"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
)

func UserHome(accName string) {
	for {
		fmt.Print("\n\n")
		utils.PtermHeaders("Homepage", pterm.FgLightRed)
		fmt.Print("\n")
		pterm.Success.Println("Login has been successful!\nLogged in to: " + accName)
		fmt.Print("\n")
		pterm.DefaultParagraph.Println("Welcome to the user homepage! From here you can access the different features of this application such as file encryption and decryption functionalities as well as account control (deleting/modifying) capabilities.")
		pterm.DefaultParagraph.Println("")
		pterm.DefaultParagraph.Printfln("In case of encrypting files, you can provide an absolute path to the readable non-binary file you want to encrypt or provide a text in the app itself which is to be encrypted as well as the password which is to be used to unlock/decrypt the file. For decryption, you will do the same except you will have to provide the decryption key.")
		pterm.DefaultParagraph.Println("")
		pterm.DefaultBasicText.WithStyle(
			pterm.NewStyle(pterm.FgGreen),
		).Println("> Select an option ---> ")
		userFunctionlityChallenge, _ := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"Encrypt Files", "Decrypt Files", "User Account"}).
		Show()
		fmt.Print("\n")
		if userFunctionlityChallenge == "Encrypt Files" {
			pterm.DefaultBasicText.WithStyle(
				pterm.NewStyle(pterm.FgGreen),
			).Println("> Select an encryption option ---> ")
			userEncryptionChallenge, _ := pterm.DefaultInteractiveSelect.
			WithOptions([]string{"Encrypt With AES", "Encrypt With RSA"}).
			Show()
			fmt.Print("\n")
			pterm.Success.Println("You have selected p.op: " + userFunctionlityChallenge + " s.op: " + userEncryptionChallenge)
			utils.ResetTerminal()
			encryptpage.EncryptPage(accName, userEncryptionChallenge)
		} else if userFunctionlityChallenge == "Decrypt Files" {
			pterm.DefaultBasicText.WithStyle(
				pterm.NewStyle(pterm.FgGreen),
			).Println("> Select a decryption option ---> ")
			userDecryptionChallenge, _ := pterm.DefaultInteractiveSelect.
			WithOptions([]string{"Decrypt AES", "Decrypt RSA"}).
			Show()
			fmt.Print("\n")
			pterm.Success.Println("You have selected p.op: " + userFunctionlityChallenge + " s.op: " + userDecryptionChallenge)
			utils.ResetTerminal()
			decryptpage.DecryptPage(accName, userDecryptionChallenge)
		} else {
			fmt.Print("\n")
			pterm.Warning.Println("Sorry, that feature is not available yet!\nIt will be added in a future update")
		}
		fmt.Print("\nReset or Go back to homepage?\nSelecting \"no\" will log you out")
		intSuccess := utils.Interruptor()
		if intSuccess {
			utils.ResetTerminal()
			continue
		}
		break
	}
}