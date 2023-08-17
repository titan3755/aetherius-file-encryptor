package account

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
)

func LoginAccount() (string, bool) {
	var accName string
	var accPass string
	loginHeaders()
	for {
		pterm.DefaultBasicText.WithStyle(
			pterm.NewStyle(pterm.FgLightBlue),
		).Print("> Enter your account name ---> ")
		fmt.Scanln(&accName)
		pterm.DefaultBasicText.WithStyle(
			pterm.NewStyle(pterm.FgLightBlue),
		).Print("> Enter your account password ---> ")
		fmt.Scanln(&accPass)
		accVerificationSuccess := userVerification(accName, accPass)
		if !accVerificationSuccess {
			utils.ErrorMsg("Wrong account name or password!\nTry again?")
			intSuccess := utils.Interruptor()
			if intSuccess {
				utils.ResetTerminal()
				loginHeaders()
				continue
			}
		} else if accVerificationSuccess {
			utils.ResetTerminal()
			return accName, true
		}
		break
	}
	return "", false
}

func loginHeaders() {
	utils.PtermHeaders("Login", pterm.FgLightMagenta)
	pterm.DefaultParagraph.Println("Welcome to the account login page! Here you can login to your account or profile which you created previously via the account creation page.")
	pterm.DefaultParagraph.Println("")
}

func userVerification(accName string, accPass string) bool {
	u, errUser := user.Current()
	if errUser != nil {
		utils.ErrorMsg(errUser.Error())
		return false
	}
	folderPath := u.HomeDir + "\\documents\\" + "\\aetherius\\" + accName
	file, errFile := os.Open(folderPath)
	if errFile != nil {
		utils.ErrorMsg("The account name you have given most likely does not exist!\nError: " + errFile.Error())
		return false
	}
	defer file.Close()
	listFiles, errReadDir := file.Readdirnames(0)
	if errReadDir != nil {
		utils.ErrorMsg(errReadDir.Error())
		return false
	}
	for i := 0; i < len(listFiles); i++ {
		if strings.TrimSpace(listFiles[i]) == "key.aetk" {
			data, errRead := os.ReadFile(folderPath + "\\key.aetk")
			if errRead != nil {
				utils.ErrorMsg(errRead.Error())
				return false
			}
			verResult := utils.PassCompare(accPass, string(data))
			return verResult
		}
	}
	return false
}
