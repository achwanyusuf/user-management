package main

import (
	"os"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	file, _ := os.OpenFile("log/appsTestLogging.log",  os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	log.SetOutput(file)
	var logger = log.New()
	logger.Out = file
	log.SetLevel(log.DebugLevel)
}

func main(){
	log.Info("Running All Cases")

	log.Info("Testing hash")
	var generatedHash = hashAndSalt("12345678")
	log.WithFields(log.Fields{
		"hash": generatedHash,
	}).Info("Hash is Generated")

	log.Info("Validating hash True")
	var validHash = validatingHash("$2a$04$Zf3eMfnvIU5b61s8DN0miOIGVig0n7VxcuYnBWIqMB1I41JDpzlUG","12345678")
	log.WithFields(log.Fields{
		"result": validHash,
	}).Info("Validation Result")

	log.Info("Validating hash False")
	var invalidHash = validatingHash("$2a$04$Zf3eMfnvIU5b61s8DN0miOIGVig0n7VxcuYnBWIqMB1I41JDpzlUG","84723773")
	log.WithFields(log.Fields{
		"result": invalidHash,
	}).Info("Validation Result")

	log.Info("Testing generate Token")
	var token = generateToken("achwan.yusuf@gmail.com")
	log.Info("Token is generated " + token)
	log.Info("Testing is completed. Check the error!")
	
}