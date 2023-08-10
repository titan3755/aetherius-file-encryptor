package utils

import "unicode"

func PasswordVerifier(pass string) (result bool, correction string) {
	const lengthErr = "Length of password must be greater than 10 characters!"
	const upLowNumSymErr = "Password must contain at least one upper case letters, lower case letters, numbers and symbols!"
	isUpperChar := false
	isLowerChar := false
	isNumber := false
	isSymbol := false
	for _, char := range pass {
		if unicode.IsUpper(char) {
			isUpperChar = true
		}
		if unicode.IsLower(char) {
			isLowerChar = true
		}
		if unicode.IsNumber(char) {
			isNumber = true
		}
		if unicode.IsSymbol(char) {
			isSymbol = true
		}
	}
	if len(pass) < 10 {
		return false, lengthErr
	}
	if !(isUpperChar && isLowerChar && isNumber && isSymbol) {
		return false, upLowNumSymErr
	}
	return true, ""
}