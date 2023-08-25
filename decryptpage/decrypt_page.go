package decryptpage

import (
	"fmt"
	"example.com/encrypt_tui/cryptography"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
)

func DecryptPage(accName string, decType string) {
	var dncPass string
	var pathName string
	pathSuccess := false
	dncHeaders(decType)
	for {
		pterm.DefaultBasicText.WithStyle(
			pterm.NewStyle(pterm.FgLightBlue),
		).Print("> Enter the absolute path to ciphertext file ---> ")
		fmt.Scanln(&pathName)
		successFile := utils.FileValidator(pathName)
		if !successFile {
			fmt.Print("\n")
			utils.ErrorMsg("An error occured while trying to read the file!\nThe path you provided is most likely wrong or invalid!\nTry again?")
			intSuccess := utils.Interruptor()
			if intSuccess {
				utils.ResetTerminal()
				dncHeaders(decType)
				continue
			}
		} else if successFile {
			pathSuccess = true
		}
		break
	}
	if pathSuccess {
		for {
			pterm.DefaultBasicText.WithStyle(
				pterm.NewStyle(pterm.FgLightBlue),
			).Print("> Enter the password for decryption ---> ")
			fmt.Scanln(&dncPass)
			fmt.Print("\n")
			pterm.Warning.Println("Decrypting files ...")
			fmt.Print("\n")
			if decType == "Decrypt AES" {
				decryptedString, decryptSuccess := cryptography.DecryptAes(pathName, dncPass)
				if decryptedString != "" && decryptSuccess {
					fmt.Print("\n")
					pterm.Success.Println("Decryption successful!")
					fmt.Print("\n")
					fmt.Print("Decrypted Text --> \n\n" + decryptedString)
				} else {
					fmt.Print("\n")
					pterm.Error.Println("Decryption was not successful!")
				}
			} else if decType == "Decrypt RSA" {
				cryptography.DecryptRsa(pathName, dncPass)
			}
		}
	}
}

func dncHeaders(decType string) {
	fmt.Print("\n\n")
	if decType == "Decrypt AES" {
		utils.PtermHeaders("Decryption [AES]", pterm.FgLightCyan)
	} else if decType == "Decrypt RSA" {
		utils.PtermHeaders("Decryption [RSA]", pterm.FgLightCyan)
	}
	fmt.Print("\n")
	pterm.DefaultParagraph.Println("Welcome to the decryption homepage! From here you can access the file decryption features of this application. Simply provide an absolute filepath to a readable non-binary file, provide the password you used to encrypt it earlier and your file will be decrypted and the data will be displayed in the app. In the future, another feature will be added using which you can save the decrypted data to another file in a location you provide.")
	pterm.DefaultParagraph.Println("")
}