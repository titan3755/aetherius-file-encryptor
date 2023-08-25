package utils

import (
	"os"
	"path/filepath"
)

func FileValidator(path string) bool {
	f, err := os.Open(path)
	f.Close()
	return err == nil
}

func DFileValidator(path string, typeDec string) bool {
	if typeDec == "AES" && filepath.Ext(path) == ".aeta" {
		return true
	} else if typeDec == "RSA" && filepath.Ext(path) == ".aetr" {
		return true
	}
	return false
}