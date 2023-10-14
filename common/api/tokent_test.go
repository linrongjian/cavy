package api

import (
	"testing"
)

func Test_Token(t *testing.T) {
	token := GenToken(1177397)
	t.Log(token)
	aa, err := ParseToken(token)
	t.Log(err)
	t.Log(aa)
}
