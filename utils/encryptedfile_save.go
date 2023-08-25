package utils

import (
	"os"
	"os/user"
	"runtime"
	"strings"
)

func EncryptedFileSave(fileName string, accName string, encryptedData string, encryptionType string) (bool, string) {
	var errorMain string = ""
	if runtime.GOOS != "windows" {
		errorMain = "Only windows currently supported!"
		return false, errorMain
	}
	u, errUser := user.Current()
	if errUser != nil {
		errorMain = errUser.Error()
		return false, errorMain
	}
	var fullPath string
	// encryptPath := u.HomeDir + "\\documents\\aetherius\\" + accName + "\\encrypted"
	if encryptionType == "aes" {
		fullPath = u.HomeDir + "\\documents\\aetherius\\" + accName + "\\encrypted\\" + strings.TrimSpace(fileName) + ".aeta"
	} else if encryptionType == "rsa" {
		fullPath = u.HomeDir + "\\documents\\aetherius\\" + accName + "\\encrypted\\" + strings.TrimSpace(fileName) + ".aetr"
	} else {
		errorMain = "Correct encryption type not found/selected!"
		return false, errorMain
	}
	filePointer, errFile := os.Create(fullPath)
	if errFile != nil {
		errorMain = errFile.Error()
		return false, errorMain
	}
	defer filePointer.Close()
	b := []byte(encryptedData)
	errWriting := os.WriteFile(fullPath, b, 0755)
	if errWriting != nil {
		errorMain = errWriting.Error()
		return false, errorMain
	}
	return true, errorMain
}