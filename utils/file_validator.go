package utils

import (
	"os"
)

func FileValidator(path string) bool {
	f, err := os.Open(path)
	f.Close()
	return err == nil
}