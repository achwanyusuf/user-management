package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"github.com/achwanyusuf/user-management/proto"
	"net/http"
)

type ReadAllRequest struct {
	Token string
}

func readAllProcess(client proto.AddServiceClient) gin.HandlerFunc {
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
		var readAllRequest ReadAllRequest
		reqBody := json.Unmarshal(value, &readAllRequest)
		if reqBody != nil {
			c.JSON(500, gin.H{
				"message": "request is invalid",
			})
		}
		req := &proto.ReadAllRequest{Token: token[0]}
		if response, err := client.ReadAll(c, req); err == nil {
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