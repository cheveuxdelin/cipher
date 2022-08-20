package tests

import (
	"testing"

	"github.com/cheveuxdelin/cipher/utils"
)

func TestToUpper(t *testing.T) {
	test_result := utils.ToUpper[byte]('a')

	if test_result != 'A' {
		t.Errorf("expected %c, but received %c", 'a', test_result)
	}
}
