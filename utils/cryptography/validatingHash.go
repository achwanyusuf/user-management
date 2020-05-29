package cryptography

import (
	"golang.org/x/crypto/bcrypt"
	log "github.com/sirupsen/logrus"
)

func ValidatingHash(hashed string, plain []byte) bool {
	log.Info("Validating Hash")
    byteHash := []byte(hashed)
    err := bcrypt.CompareHashAndPassword(byteHash, plain)
    if err != nil {
		log.Error("Has is not Validated")
        return false
	}
	log.Info("Hash is validated")
    return true
}