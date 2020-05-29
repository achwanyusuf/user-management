package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"github.com/achwanyusuf/user-management/proto"
	"net/http"
)

type LoginRequest struct {
	Email string
	Password string
}

func loginProcess(client proto.AddServiceClient) gin.HandlerFunc {
	return func(c *gin.Context){
		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil{
			log.Error(err)
			c.JSON(500, gin.H{
				"message": "request is invalid",
			})
		}
		var loginRequest LoginRequest
		reqBody := json.Unmarshal(value, &loginRequest)
		if reqBody != nil {
			c.JSON(500, gin.H{
				"message": "request is invalid",
			})
		}
		req := &proto.LoginRequest{Email:loginRequest.Email, Password: loginRequest.Password}
		if response, err := client.Login(c, req); err == nil {
			if response.Message == "" {
				c.JSON(http.StatusOK, response)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": response.Message})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
}