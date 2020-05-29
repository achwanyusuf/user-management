package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"github.com/achwanyusuf/user-management/proto"
	"net/http"
)

type DeleteRequest struct {
	UserId string
}

func deleteProcess(client proto.AddServiceClient) gin.HandlerFunc {
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
		var deleteRequest DeleteRequest
		reqBody := json.Unmarshal(value, &deleteRequest)
		if reqBody != nil {
			c.JSON(500, gin.H{
				"message": "request is invalid",
			})
		}
		req := &proto.DeleteRequest{UserId:deleteRequest.UserId, Token: token[0]}
		if response, err := client.Delete(c, req); err == nil {
			if response.Message == ""{	
				c.JSON(http.StatusOK, gin.H{"message": "Data is Deleted!"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": response.Message})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}
}