package services

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidateEmail(email string) (bool, error) {

	if len(email) == 0 {
		return false, errors.New(fmt.Sprintf("Please provide email id"))
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email), nil
}
