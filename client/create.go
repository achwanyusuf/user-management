package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"github.com/achwanyusuf/user-management/proto"
	"net/http"
)

type CreateRequest struct {
	Email string
	Password string
	Address string
	Token string
}

func createProcess(client proto.AddServiceClient) gin.HandlerFunc {
	return func(c *gin.Context){
		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil{
			log.Error(err)
			c.JSON(500, gin.H{
				"message": "request is invalid",
			})
		}
		token := c.Request.Header["Authorization"]
		var createRequest CreateRequest
		reqBody := json.Unmarshal(value, &createRequest)
		if reqBody != nil {
			c.JSON(500, gin.H{
				"message": "request is invalid",
			})
		}
		req := &proto.CreateRequest{Email:createRequest.Email, Password: createRequest.Password, Address: createRequest.Address, Token: token[0]}
		if response, err := client.Create(c, req); err == nil {
			if response.Message == ""{
				c.JSON(http.StatusOK, response)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": response.Message})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
}