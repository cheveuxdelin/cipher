package morse

import (
	"errors"
	"fmt"
	"strings"
)

type code struct {
	char byte
	code string
}

func toUpper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - ('a' - 'A')
	}
	return b
}
func isASCII(b byte) bool {
	return b >= ' ' && b <= '~'
}

var codes []code = []code{
	{'A', ".-"},
	{'B', "-..."},
	{'C', "-.-."},
	{'D', "-.."},
	{'E', "."},
	{'F', "..-."},
	{'G', "--."},
	{'H', "...."},
	{'I', ".."},
	{'J', ".---"},
	{'K', "-.-"},
	{'L', ".-.."},
	{'M', "--"},
	{'N', "-."},
	{'O', "---"},
	{'P', ".--."},
	{'Q', "--.-"},
	{'R', ".-."},
	{'S', "..."},
	{'T', "-"},
	{'U', "..-"},
	{'V', "...-"},
	{'W', ".--"},
	{'X', "-..-"},
	{'Y', "-.--"},
	{'Z', "--.."},
	{'!', "-.-.--"},
	{'"', ".-..-."},
	{'&', ".-..."},
	{'\'', ".----."},
	{'(', "-.--."},
	{')', "-.--.-"},
	{'+', ".-.-."},
	{',', "--..--"},
	{'-', "-....-"},
	{'.', ".-.-.-"},
	{'/', "-..-."},
	{'0', "-----"},
	{'1', ".----"},
	{'2', "..---"},
	{'3', "...--"},
	{'4', "....-"},
	{'5', "....."},
	{'6', "-...."},
	{'7', "--..."},
	{'8', "---.."},
	{'9', "----."},
	{':', "---..."},
	{'=', "-...-"},
	{'?', "..--.."},
	{'@', ".--.-."},
	{'_', "..--.-"},
}

func Encode(s string) (string, error) {
	// checking if s is not empty
	if len(s) == 0 {
		return "", errors.New("invalid string")
	}

	// creating map
	var d map[byte]string = make(map[byte]string, len(codes))
	for i := 0; i < len(codes); i++ {
		d[codes[i].char] = codes[i].code
	}

	// builder variable
	var rtn strings.Builder

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return "", errors.New("string containing only spaces")
	}

	for i := 0; i < len(s)-1; i++ {

		// checking if its ASCII
		if !isASCII(s[i]) {
			return "", fmt.Errorf("char not ASCII: %c", s[i])
		}

		if s[i] == ' ' {
			rtn.WriteString("/ ")
		} else {
			// getting char from d
			value, ok := d[toUpper(s[i])]

			if ok {
				rtn.WriteString(value)
				rtn.WriteString(" ")
			}
		}
	}

	value, ok := d[toUpper(s[len(s)-1])]
	if ok {
		rtn.WriteString(value)
	}

	return rtn.String(), nil
}

func Decode(s string) (string, error) {

	if len(s) == 0 {
		return "", errors.New("invalid string")
	}

	var d map[string]byte = make(map[string]byte, len(codes))
	for i := 0; i < len(codes); i++ {
		d[codes[i].code] = codes[i].char
	}

	var b, rtn strings.Builder
	var missingSpace bool

	var f func() = func() {
		if b.Len() != 0 {
			value, ok := d[b.String()]
			if ok {
				rtn.WriteByte(value)
			}
			b.Reset()
			missingSpace = true
		}
	}

	for i := range s {
		switch s[i] {
		case '.', '-':
			b.WriteByte(s[i])
		case ' ':
			f()
		case '/':
			f()
			if missingSpace {
				rtn.WriteByte(' ')
				missingSpace = false
			}
		default:
			return "", fmt.Errorf("invalid character: %c", s[i])
		}
	}

	f()

	return rtn.String(), nil
}
