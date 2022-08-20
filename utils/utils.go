package utils

import (
	"strings"

	"github.com/cheveuxdelin/cipher/errors"
)

func ToUpper[T byte | rune](s T) T {
	if s >= 'a' && s <= 'z' {
		return s - ('a' - 'A')
	}
	return s
}

func IsASCIIPrintable[T byte | rune](s T) bool {
	return s >= ' ' && s <= '~'
}

func ValidateString(s *string) error {
	if len(*s) == 0 {
		return errors.ErrEmptyString
	}
	*s = strings.TrimSpace(*s)
	if len(*s) == 0 {
		return errors.ErrOnlySpaces
	}
	for i := range *s {
		if !IsASCIIPrintable((*s)[i]) {
			return errors.ErrStringNotASCII
		}
	}
	return nil
}
