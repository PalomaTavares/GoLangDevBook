package security

import "golang.org/x/crypto/bcrypt"

// transforms string in hash
func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// compares password and hash and returns if they are the same
func VerifyPassword(passHash string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(pass))
}
