package cod

import (
	"testing"

	"github.com/matryer/is"
)

// TestIsValid checks returned values from IsValid function from cod package
func TestIsValid(t *testing.T) {

	is := is.New(t)

	// table of invalid and or valid users
	var userTable = []struct {
		username string
		expected bool
	}{
		{"Temporis%231318655", true},
		{"onesicksikh%231221896", true},
		{"ghost%232891963", true},
		{"fehyifue8", false},
		{"%23", false},
		{"\n", false},
		{"Temporis#1318655", false},
		{"ghost#2891963", false},
		{"onesicksikh#1221896", false},
	}
	// each user in table
	for _, user := range userTable {
		result := IsValid(user.username)
		expected := user.expected
		is.Equal(result, expected)
	}
}
