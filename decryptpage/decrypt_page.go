package decryptpage

import (
	"fmt"
	"example.com/encrypt_tui/cryptography"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
	"os"
	"os/user"
	"runtime"
)

func DecryptPage(accName string, decType string) {
	for {
	utils.ResetTerminal()
	dncHeaders(decType)
	if runtime.GOOS != "windows" {
		utils.ErrorMsg("Only windows OS is supported right now!")
		return
	}
	u, errUser := user.Current()
	if errUser != nil {
		utils.ErrorMsg("There has been an error while trying to find user!")
		return
	}
	encryptPath := u.HomeDir + "\\documents\\aetherius\\" + accName + "\\encrypted"
	if _, err := os.Stat(encryptPath); os.IsNotExist(err) {
		fileErr := os.Mkdir(encryptPath, 0755)
		if fileErr != nil {
			utils.ErrorMsg("There has been an error while trying to create the encrypt folder!")
			return
		}
	}
	var dncPass string
	var pathName string
	pathSuccess := false
	response, successSelector := dncSelector(decType, encryptPath)
	if !successSelector {
		utils.ErrorMsg("An error has occured during file selection!\nMost likely you do not have any encrypted files in your profile folder")
		return
	}
	pathName = encryptPath + "\\" + response
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
	if pathSuccess {
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
				fmt.Print("\n")
				break
			} else {
				fmt.Print("\n")
				pterm.Error.Println("Decryption was not successful!\nCheck if the provided file or password is valid")
				fmt.Print("\n")
				fmt.Println("Try again? ")
				userResponse := utils.Interruptor()
				if userResponse {
					continue
				} else {
					break
				}
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

func dncSelector(decType string, path string) (string, bool) {
	var fileArray []string
	entries, err := os.ReadDir(path)
    if err != nil {
        return "", false
    }
	if len(entries) == 0 {
		return "", false
	}
    for _, e := range entries {
        fileArray = append(fileArray, e.Name())
    }
	pterm.DefaultBasicText.WithStyle(
		pterm.NewStyle(pterm.FgGreen),
	).Println("> Select a decryption file option ---> ")
	userDecryptionFileChallenge, _ := pterm.DefaultInteractiveSelect.
	WithOptions(fileArray).
	Show()
	return userDecryptionFileChallenge, true
}