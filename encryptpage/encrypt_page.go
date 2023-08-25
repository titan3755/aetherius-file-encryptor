package encryptpage

import (
	"fmt"
	"example.com/encrypt_tui/cryptography"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
)

func EncryptPage(accName string, encType string) {
	var pathName string
	var encPass string
	pathSuccess := false
	pwSuccess := false
	encHeaders(encType)
	for {
		pterm.DefaultBasicText.WithStyle(
			pterm.NewStyle(pterm.FgLightBlue),
		).Print("> Enter the absolute path to file ---> ")
		fmt.Scanln(&pathName)
		successFile := utils.FileValidator(pathName)
		if !successFile {
			fmt.Print("\n")
			utils.ErrorMsg("An error occured while trying to read the file!\nThe path you provided is most likely wrong or invalid!\nTry again?")
			intSuccess := utils.Interruptor()
			if intSuccess {
				utils.ResetTerminal()
				encHeaders(encType)
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
			).Print("> Enter the password for encryption ---> ")
			fmt.Scanln(&encPass)
			passSuccess, errPass := utils.PasswordVerifier(encPass)
			if passSuccess {
				pwSuccess = true
				break
			} else {
				pterm.DefaultBasicText.WithStyle(
					pterm.NewStyle(pterm.FgRed),
				).Print("!> " + errPass + " (try again)\n\n")
			}
		}
	}
	if pwSuccess {
		fmt.Print("\n")
		pterm.Warning.Println("Encrypting files...")
		fmt.Print("\n")
		if encType == "Encrypt With AES" {
			cipherText, encryptSuccess := cryptography.EncryptAes(pathName, encPass)
			if cipherText != "" && encryptSuccess {
				var fileName string
				fmt.Print("\n")
				pterm.Success.Println("Encryption successful!")
				fmt.Print("\n")
				pterm.DefaultBasicText.WithStyle(
					pterm.NewStyle(pterm.FgLightBlue),
				).Print("> Enter the file name for saving ciphertext ---> ")
				fmt.Scanln(&fileName)
				fmt.Print("\n")
				pterm.Warning.Println("Saving ciphertext ...")
				cipherTextSavingSuccess, errorMain := utils.EncryptedFileSave(fileName, accName, cipherText, "aes")
				if cipherTextSavingSuccess {
					fmt.Print("\n")
					pterm.Success.Println("CipherText saved successfully!")
				} else {
					fmt.Print("\n")
					pterm.Error.Println("An error occured while saving the ciphertext!\n"+errorMain)
				}
			} else {
				fmt.Print("\n")
				pterm.Error.Println("Encryption was not successful!")
			}
		} else if encType == "Encrypt With RSA" {
			cryptography.EncryptRsa(pathName, encPass)
		}
	}
}

func encHeaders(encType string) {
	fmt.Print("\n\n")
	if encType == "Encrypt With AES" {
		utils.PtermHeaders("Encryption [AES]", pterm.FgYellow)
	} else if encType == "Encrypt With RSA" {
		utils.PtermHeaders("Encryption [RSA]", pterm.FgYellow)
	}
	fmt.Print("\n")
	pterm.DefaultParagraph.Println("Welcome to the encryption homepage! From here you can access the file encryption features of this application. Simply provide an absolute filepath to a readable non-binary file, provide a password and you file will be encrypted and relocated to a folder in your profile directory.")
	pterm.DefaultParagraph.Println("")
	pterm.DefaultParagraph.Printfln("You can also decrypt the same file by heading to the decryption page and providing the same password you provided to encrypt that particular file which you now want to decrypt. Note that if you forget or lose your encryption password, there will be no way to decrypt that file again.")
	pterm.DefaultParagraph.Println("")
}

