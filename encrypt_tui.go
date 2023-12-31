package main

import (
	"fmt"

	"example.com/encrypt_tui/account"
	"example.com/encrypt_tui/userpages"
	"example.com/encrypt_tui/utils"
	"github.com/pterm/pterm"
);

func main() {
	var runVar bool = true
	for runVar {
		utils.ResetTerminal()
		runMain()
		exitProcs(&runVar)
	}
}

func runMain() {
	utils.PtermHeaders("Info Page", pterm.FgLightBlue)
	pterm.DefaultParagraph.Println("Welcome to Aetherius File Encryptor CLI app! Using this command line application, you can encrypt any text file or any readable non-binary file with on your computer and access them easily with a provided password. This application will not store or withhold any of your passwords or personal information and the code is free for anyone to inspect and find security vulnerabilities.")
	pterm.DefaultParagraph.Println()
	pterm.DefaultParagraph.Println("By using this app, you will be able to encrypt your readable non-binary files using different encryption algorithms. This app is still in development, so only 2 encryption modes will be available at first. They are listed below.")
	pterm.DefaultParagraph.Println()
	var viewEnAlDes string
	pterm.DefaultBasicText.WithStyle(
		pterm.NewStyle(pterm.FgLightBlue),
	).Print("> View encryption algorithm descriptions(y/n) ---> ")
	fmt.Scanln(&viewEnAlDes)
	if (viewEnAlDes == "y") {
		pterm.Info.Println("Encryption Algorithm Description Boxes --->")
		panel1 := pterm.DefaultBox.WithTitle("RSA").WithTitleTopCenter(true).Sprint("RSA (Rivest-Shamir-Adleman) is\n a public-key cryptosystem, one of the oldest,\n that is widely used for secure data transmission.\n The acronym \"RSA\" comes from the surnames of Ron\n Rivest, Adi Shamir and Leonard Adleman, who\n publicly described the algorithm in 1977.")
		panel2 := pterm.DefaultBox.WithTitle("AES").WithTitleTopCenter(true).Sprint("The Advanced Encryption Standard\n (AES) is a symmetric block cipher chosen by \nthe U.S. government to protect\n classified information.\n AES is implemented in software and hardware\n throughout the world to\n encrypt sensitive data.")
		panels, _ := pterm.DefaultPanel.WithPanels(pterm.Panels{
			{{Data: panel1}, {Data: panel2}},
		}).Srender()
		pterm.DefaultBox.WithTitle("Algorithm Descriptions").WithTitleTopCenter(true).WithRightPadding(0).WithBottomPadding(0).Println(panels)
		pterm.Print("\n\n")
	}
	pterm.DefaultBasicText.WithStyle(
		pterm.NewStyle(pterm.FgGreen),
	).Println("> Select a user account option ---> ")
	userAccountChallenge, _ := pterm.DefaultInteractiveSelect.
	WithOptions([]string{"Create a new account", "Login to existing account"}).
	Show()
	utils.ResetTerminal()
	if (userAccountChallenge == "Create a new account") {
		dettSucc := utils.Detector()
		if !dettSucc {
			fmt.Println("An error occured!")
		} else {
			account.CreateAccount()
		}
	} else if (userAccountChallenge == "Login to existing account") {
		dettSucc := utils.Detector()
		if !dettSucc {
			fmt.Println("An error occured!")
		} else {
			loggedAccount, success := account.LoginAccount()
			if success {
				userpages.UserHome(loggedAccount)
			}
		}
	} else {
		fmt.Println("An error occured!")
	}
}

func exitProcs(runIf *bool) {
	var runState string
	pterm.DefaultBasicText.WithStyle(
		pterm.NewStyle(pterm.FgRed),
	).Print("\n\n> Exit or reset the program? (e/r) ---> ")
	fmt.Scanln(&runState)
	if (runState == "e") {
		*runIf = false
	}
}
