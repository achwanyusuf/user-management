package main

import (
	"github.com/achwanyusuf/user-management/utils/cryptography"
)

func validatingHash(hashed string,input string) bool{
	result := cryptography.ValidatingHash(hashed,[]byte(input))
	return result
}