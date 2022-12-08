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
	regex, _ := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	return regex.MatchString(s)
}
