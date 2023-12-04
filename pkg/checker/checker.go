package checker

import (
	"regexp"
	"unicode"
)

const emailPattern = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"

func IsValidEmail(email string) bool {
	m, err := regexp.MatchString(emailPattern, email)
	return m && err == nil
}

func IsValidPassword(password string) bool {
	return len(password) >= 8 && isLetter(password[0]) && onlyLettersAndDigits(password)
}

func isLetter(char byte) bool {
	return (97 <= char && char <= 122) || (65 <= char && char <= 90)
}

func onlyLettersAndDigits(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
