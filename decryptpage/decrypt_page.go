package decryptpage

import (
	"fmt"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
)

func DecryptPage(accName string, decType string) {
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