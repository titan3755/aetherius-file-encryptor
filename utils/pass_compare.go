package utils

import (
	"crypto/sha256"
	"fmt"
)

func PassCompare(pass string, hashedPass string) bool {
	encrypted := sha256.New()
	_, errWriteFirst := encrypted.Write([]byte(pass))
	if errWriteFirst != nil {
		return false
	}
	encPass := encrypted.Sum(nil)
	encEncrypt := sha256.New()
	_, errWriteSecond := encEncrypt.Write([]byte(fmt.Sprintf("%x", encPass)))
	if errWriteSecond != nil {
		return false
	}
	encEncryptPass := encEncrypt.Sum(nil)
	return fmt.Sprintf("%x", encEncryptPass) == hashedPass
}