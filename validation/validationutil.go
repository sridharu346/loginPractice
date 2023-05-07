package validation

import (
	"regexp"
)

func IsEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return match
}

func IsPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}
	return true
}

func IsPhoneNumberValid(Phone string) bool {
	pattern := `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`
	match, err := regexp.MatchString(pattern, Phone)
	if err != nil {
		return false
	}
	return match
}
