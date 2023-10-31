package config_test

import (
	"testing"

	"github.com/Lucas-Linhar3s/debtManager/config"
)

func TestBcrypt(t *testing.T) {
	password := "123456"
	hash, err := config.GenerateHash(password)
	if err != nil {
		t.Error("Error to generate hash")
	}

	if !config.VerifyHash(hash, password) {
		t.Error("Error to verify hash")
	}

	t.Log("Hash: ", hash)

}
