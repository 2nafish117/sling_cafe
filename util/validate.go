package util

import (
	"errors"
	"regexp"
	"strconv"
)

// ValidateRequireAndLengthAndRegex is used to validate any input data but in string type
// @params value is the input value
// @params isRequired definds the wheather the input value required or not
// @params minLength definds minimum length of the input value, 0 value definds no min length check
// @params maxLength definds maximum length of the input value,  0 value definds no max length check
// @params maxLength definds regex of the input value, "" value definds no regex required
// @returns error if any
func ValidateRequireAndLengthAndRegex(value string, isRequired bool, minLength, maxLength int, regex, fieldName string) error {

	length := len(value)
	Re := regexp.MustCompile(regex)
	if isRequired == true && length < 1 {
		return errors.New(fieldName + " is Required")
	}

	// Min length check
	// If params min length value is zero that indecates, there will be no min length check
	if minLength != 0 && length > 1 && length < minLength {
		return errors.New(fieldName + " must be min " + strconv.Itoa(minLength))
	}

	// Max length check
	// If params max length value is zero that indecates, there will be no max length check
	if maxLength != 0 && length > 1 && length > maxLength {
		return errors.New(fieldName + " must be max " + strconv.Itoa(maxLength))
	}

	if !Re.MatchString(value) { // Regex check
		return errors.New("Invalid " + fieldName)
	}

	return nil
}
