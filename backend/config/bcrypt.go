package config

import (
	"github.com/jameskeane/bcrypt"
)

// GenerateHash is a function to generate a hash from a password
func GenerateHash(password string) (string, error) {
	salt, err := bcrypt.Salt(10)
	if err != nil {
		return "", err
	}

	hash, err := bcrypt.Hash(password, salt)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// VerifyHash is a function to verify a hash from a password
func VerifyHash(hash, password string) bool {
	return bcrypt.Match(password, hash)
}
