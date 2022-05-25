package utils

func ToUpper[T byte | rune](s T) T {
	if s >= 'a' && s <= 'z' {
		return s - ('a' - 'A')
	}
	return s
}

func IsASCIIPrintable[T byte | rune](s T) bool {
	return s >= ' ' && s <= '~'
}
