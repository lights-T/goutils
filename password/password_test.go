package password

import "testing"

func TestHashPassword(t *testing.T) {
	t.Log(HashPassword("111111"))
	t.Log(HashPassword(""))
}

func TestCheckPasswordHash(t *testing.T) {
	t.Log(CheckPasswordHash("111111", "$2a$10$x3mXSITWvHROfqIH/Ti7rOw3vP39Wdu.DP1krcyUSpvQR3k8wOcIW"))
}
