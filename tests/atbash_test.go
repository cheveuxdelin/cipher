package tests

import (
	"strconv"
	"testing"

	"github.com/cheveuxdelin/cipher/atbash"
	"github.com/cheveuxdelin/cipher/errors"
)

type testStruct struct {
	s    string
	want wantStruct
}

type wantStruct struct {
	rtn string
	err error
}

func newTestStruct(s string, rtn string, err error) testStruct {
	return testStruct{
		s: s,
		want: wantStruct{
			rtn: rtn,
			err: err,
		},
	}
}

func TestAtbashEncoding(t *testing.T) {

	var tests = [...]testStruct{
		newTestStruct("a", "=", nil),
		newTestStruct("=", "a", nil),
		newTestStruct("ƒ", "", errors.ErrStringNotASCII),
		newTestStruct("", "", errors.ErrEmptyString),
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ans, err := atbash.Encode(tt.s)

			if err != tt.want.err {
				t.Errorf("got %s, want %s", err.Error(), tt.want.err.Error())
			}

			if ans != tt.want.rtn {
				t.Errorf("got %s, want %s", ans, tt.want.rtn)
			}
		})
	}
}
