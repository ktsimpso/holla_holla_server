package auth

import (
	"code.google.com/p/go.crypto/bcrypt"
	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/sha512"
)

func HashPassword(password []byte) ([]byte, error) {
	// Do not need to provide a salt here, bcrypt will generate one
	key := pbkdf2.Key(password, []byte(""), 1024, 64, sha512.New)
	return bcrypt.GenerateFromPassword(key, bcrypt.DefaultCost)
}
