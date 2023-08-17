package account

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
)

func CreateAccount() {
	for {
		var accName string
		var accPass string
		utils.PtermHeaders("Create Account", pterm.FgLightRed)
		pterm.DefaultParagraph.Println("Welcome to the account or profile creation page! Here you can create a profile in order to use this app properly. Your profile will be saved in the documents folder and it will contain the hashed form of your passwords or private keys and your public keys as well. The password you will provide here will be used as the private key by default unless defaults are overridden.")
		pterm.DefaultParagraph.Println("")
		pterm.DefaultParagraph.Printfln("For enhancing security further, if your profile is somehow lost or deleted, no one will ever be able to access your \"Aetherius\" encrypted files as the public and private keys used to encrypt your files will be lost along with your profile folder. A feature to get back your lost profile and keys will be added to this app some time in the future.")
		pterm.DefaultParagraph.Println("")
		for {
			pterm.DefaultBasicText.WithStyle(
				pterm.NewStyle(pterm.FgLightBlue),
			).Print("> Enter your account name ---> ")
			fmt.Scanln(&accName)
			if len(accName) < 5 {
				pterm.DefaultBasicText.WithStyle(
					pterm.NewStyle(pterm.FgRed),
				).Print("!> " + "The account name cannot be less than 5 characters!" + " (try again)\n\n")
				continue
			}
			break
		}
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
		fmt.Print("\n")
		pterm.Warning.Println("Creating profile ...")
		dirSuccess := directoryCreator(accName)
		if !dirSuccess {
			utils.ErrorMsg("There has been an error while attempting to create profile\nWould you like to attempt account creation again?")
			intSuccess := utils.Interruptor()
			if intSuccess {
				utils.ResetTerminal()
				continue
			}
		}
		pterm.Warning.Println("Creating key ...")
		keySuccess := keyCreator(accName, accPass)
		if !keySuccess {
			utils.ErrorMsg("There has been an error while attempting to create profile key\nWould you like to attempt account creation again?")
			intSuccess := utils.Interruptor()
			if intSuccess {
				utils.ResetTerminal()
				continue
			}
		}
		profDetSuccess := profileDetails(accName)
		if !profDetSuccess {
			utils.ErrorMsg("There has been an error while attempting to display profile information")
		}
		break
	}
}

func directoryCreator(accName string) bool {
	if runtime.GOOS != "windows" {
		utils.ErrorMsg("OS other than windows not currently supported!")
		return false
	}
	u, err := user.Current()
	if err != nil {
		utils.ErrorMsg(err.Error())
		return false
	}
	docxPath := u.HomeDir + "\\documents\\aetherius\\"
	file, errFile := os.Open(docxPath)
	if errFile != nil {
		utils.ErrorMsg(errFile.Error())
		return false
	}
	defer file.Close()
	listFiles, errReadDir := file.Readdirnames(0)
	if errReadDir != nil {
		utils.ErrorMsg(errReadDir.Error())
		return false
	}
	for i := 0; i < len(listFiles); i++ {
		if strings.TrimSpace(listFiles[i]) == strings.TrimSpace(accName) {
			utils.ErrorMsg("A profile of that name already exists!")
			return false
		}
	}
	errOs := os.Mkdir(docxPath + accName, 0755)
	if errOs != nil {
		utils.ErrorMsg(errOs.Error())
		return false
	}
	pterm.Success.Println("Profile directory created successfully!")
	return true
}

func keyCreator(profileName string, password string) bool {
	u, errUser := user.Current()
	if errUser != nil {
		utils.ErrorMsg(errUser.Error())
		return false
	}
	dirPath := u.HomeDir + "\\documents\\aetherius\\" + profileName
	f, errFile := os.Create(dirPath + "\\key.aetk")
	if errFile != nil {
		utils.ErrorMsg(errFile.Error())
		return false
	}
	defer f.Close()
	encrypted := sha256.New()
	_, errWriteFirst := encrypted.Write([]byte(password))
	if errWriteFirst != nil {
		utils.ErrorMsg(errWriteFirst.Error())
		return false
	}
	encPass := encrypted.Sum(nil)
	encEncrypt := sha256.New()
	_, errWriteSecond := encEncrypt.Write([]byte(fmt.Sprintf("%x", encPass)))
	if errWriteSecond != nil {
		utils.ErrorMsg(errWriteSecond.Error())
		return false
	}
	encEncryptPass := encEncrypt.Sum(nil)
	_, errWrite := f.WriteString(fmt.Sprintf("%x", encEncryptPass))
	if errWrite != nil {
		utils.ErrorMsg(errWrite.Error())
		return false
	}
	pterm.Success.Println("Created key successfully!")
	return true
}

func profileDetails(accName string) bool {
	u, errUser := user.Current()
	if errUser != nil {
		utils.ErrorMsg(errUser.Error())
		return false
	}
	dirPath := u.HomeDir + "\\documents\\aetherius\\" + accName
	pterm.DefaultSection.Println("New profile information --->")
	pterm.Info.Println("Profile name: " + accName + "\nProfile Location: " + dirPath)
	return true
}