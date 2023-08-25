package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"io"
	"os"

	"example.com/encrypt_tui/utils"
)

func DecryptAes(pathName string, keyString string) (string, bool) {
	validateSuccess := utils.DFileValidator(pathName, "AES")
	if !validateSuccess {
		return "", false
	}
	file, errFile := os.Open(pathName)
	if errFile != nil {
		utils.ErrorMsg("An error occured while trying to open the data file!")
		return "", false
	}
	cipherText, errRead := io.ReadAll(file)
	if errRead != nil {
		utils.ErrorMsg("An error occured while trying to read the data file!")
		return "", false
	}
	byteString := make([]byte, 32)
	copy(byteString, []byte(keyString))
	key, decodeErr := hex.DecodeString(hex.EncodeToString(byteString))
	if decodeErr != nil {
		utils.ErrorMsg("An error occured while trying to decode the string!::"+decodeErr.Error())
		return "", false
	}
	cipTxt, cipTxtDecodeErr := hex.DecodeString(string(cipherText))
	if cipTxtDecodeErr != nil {
		utils.ErrorMsg("An error occured while trying to decode cipTxt!")
		return "", false
	}
	block, errBlock := aes.NewCipher(key)
	if errBlock != nil {
		utils.ErrorMsg("An error occured while trying to create new cipher!")
		return "", false
	}
	aesGCM, gcmErr := cipher.NewGCM(block)
	if gcmErr != nil {
		utils.ErrorMsg("An error occured while trying to generate GCM!")
		return "", false
	}
	nonceSize := aesGCM.NonceSize()
	nonce, cipherTxt := cipTxt[:nonceSize], cipTxt[nonceSize:]
	plainTxt, errPlainTxt := aesGCM.Open(nil, nonce, cipherTxt, nil)
	if errPlainTxt != nil {
		utils.ErrorMsg("An error occured while trying to generate plaintxt!")
		return "", false
	}
	return string(plainTxt), true
}