package utils

import (
	"os/user"
	"os"
	"strings"
)

const appFolder = "aetherius"

func Detector() bool {
	u, err := user.Current()
	if err != nil {
		return false
	}
	docxPath := u.HomeDir + "\\documents\\"
	file, errFile := os.Open(docxPath)
	if errFile != nil {
		return false
	}
	defer file.Close()
	listFiles, errReadDir := file.Readdirnames(0)
	if errReadDir != nil {
		return false
	}
	for i := 0; i < len(listFiles); i++ {
		if strings.TrimSpace(listFiles[i]) == appFolder {
			return true
		}
	}
	errOs := os.Mkdir(docxPath + appFolder, 0755)
	return errOs == nil
}