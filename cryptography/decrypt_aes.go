package cryptography

import (
	"example.com/encrypt_tui/utils"
	"os"
	"io"
	"encoding/hex"
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
	_, errRead := io.ReadAll(file)
	if errRead != nil {
		utils.ErrorMsg("An error occured while trying to read the data file!")
		return "", false
	}
	byteString := make([]byte, 32)
	copy(byteString, []byte(keyString))
	_, decodeErr := hex.DecodeString(hex.EncodeToString(byteString))
	if decodeErr != nil {
		utils.ErrorMsg("An error occured while trying to decode the string!::"+decodeErr.Error())
		return "", false
	}
	return "", false
}