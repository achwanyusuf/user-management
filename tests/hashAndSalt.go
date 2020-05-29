package main

import (
	"github.com/achwanyusuf/user-management/utils/cryptography"
)

func hashAndSalt(input string) string{
	return cryptography.HashAndSalt([]byte(input))
}