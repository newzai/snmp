package xhttp

import (
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"

	"github.com/gin-gonic/gin"
)

type deleteUserData struct {
	ID int `json:"userid"`
}
type deleteUserRequest struct {
	Token string         `json:"token"`
	Data  deleteUserData `json:"data"`
}

func deleteUser(c *gin.Context) {
	var request deleteUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	model.RemoteUser(request.Data.ID, xdb.Engine)
	c.JSON(http.StatusOK, gin.H{"result": 0, "message": "OK"})
}
