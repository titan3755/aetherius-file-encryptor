package utils

import (
	"os"
)

func FileValidator(path string) bool {
	_, err := os.Open(path)
	return err == nil
}