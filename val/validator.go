package val

import (
	"fmt"
	"net/mail"
	"regexp"
	"unicode"
)

var isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 5, 30); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}
	return nil
}

func ValidatePassword(value string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(value) >= 7 {
		hasMinLen = true
	}
	for _, char := range value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if !hasMinLen || !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return fmt.Errorf("password is not valid must include at least one special character,one numeric value, one uppercase and lowercase letter and have atleast 7 characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	if err := ValidateString(email, 3, 80); err != nil {
		return err
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	return nil
}
