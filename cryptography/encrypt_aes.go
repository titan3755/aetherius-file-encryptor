package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"example.com/encrypt_tui/utils"
)

func EncryptAes(pathName string, encPass string) (string, bool) {
	file, errFile := os.Open(pathName)
	if errFile != nil {
		utils.ErrorMsg("An error occured while trying to open the data file!")
		return "", false
	}
	b, errRead := io.ReadAll(file)
	if errRead != nil {
		utils.ErrorMsg("An error occured while trying to read the data file!")
		return "", false
	}
	byteString := make([]byte, 32)
	copy(byteString, []byte(encPass))
	key, decodeErr := hex.DecodeString(hex.EncodeToString(byteString))
	if decodeErr != nil {
		utils.ErrorMsg("An error occured while trying to decode the string!::"+decodeErr.Error())
		return "", false
	}
	plainText := b
	block, errBlock := aes.NewCipher(key)
	if errBlock != nil {
		utils.ErrorMsg("An error occured while trying to encrypt the data!")
		return "", false
	}
	aesGCM, errGCM := cipher.NewGCM(block)
	if errGCM != nil {
		utils.ErrorMsg("An error occured while trying to create new GCM!")
		return "", false
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, errIO := io.ReadFull(rand.Reader, nonce)
	if errIO != nil {
		utils.ErrorMsg("An error occured while trying to IO Read!")
		return "", false
	}
	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)
	defer file.Close()
	return fmt.Sprintf("%x", cipherText), true
}