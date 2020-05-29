package cryptography

import (
	"golang.org/x/crypto/bcrypt"
	log "github.com/sirupsen/logrus"
)

func HashAndSalt(password []byte) string {
	log.Info("Generating Hash")
    hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
    if err != nil {
        log.Error(err)
	}
	log.Info("Hash is generated")
    return string(hash)
}