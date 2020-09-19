package cod_test

import (
	"testing"
	"zleague/api/cod"

	"github.com/matryer/is"
)

func TestValidUser(t *testing.T) {
	is := is.New(t)
	username := "Temporis%231318655"

	valid := cod.IsValid(username)
	is.Equal(valid, true)
}

func TestInvalidUser(t *testing.T) {
	is := is.New(t)
	username := "Temporis%2312345"

	valid := cod.IsValid(username)
	is.Equal(valid, false)
}
