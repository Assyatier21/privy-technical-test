package utils

import "regexp"

func IsValidAlphabet(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z]+$)`)
	return regex.MatchString(s)
}

func IsValidNumeric(s string) bool {
	regex, _ := regexp.Compile(`([0-9])`)
	return regex.MatchString(s)
}

func IsValidAlphaNumeric(s string) bool {
	regex, _ := regexp.Compile(`(^[a-zA-Z0-9]*$)`)
	return regex.MatchString(s)
}

func IsValidAlphaNumericHyphen(s string) bool {
	regex, _ := regexp.Compile(`[a-zA-Z0-9-]+`)
	return regex.MatchString(s)
}

func IsValidFloatNumber(s string) bool {
	regex, _ := regexp.Compile(`^[-+]?[0-9]*\.[0-9]+$`)
	return regex.MatchString(s)
}

func IsValidLinkImage(s string) bool {
	regex, _ := regexp.Compile(`(http)?s?:?(\/\/[^"']*\.(?:png|jpg|jpeg|gif|png|svg))`)
	return regex.MatchString(s)
}
