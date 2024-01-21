package security

import "golang.org/x/crypto/bcrypt"

// Hash generates a hash from a given password.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Check compares a given password with a hash.
func Check(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
