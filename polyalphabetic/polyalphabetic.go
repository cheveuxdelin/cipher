package polyalphabetic

func Encode(s string, key string) (string, error) {
	var rtn []byte = []byte(s)

	for i := range s {
		rtn[i] = rtn[i] + key[len(key)%i]
	}

	return string(rtn), nil
}
