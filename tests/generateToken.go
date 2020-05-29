package main

import (
	"github.com/achwanyusuf/user-management/utils/cryptography"
)

func generateToken(email string) string{
	result := cryptography.GenerateToken(email)
	return result
}