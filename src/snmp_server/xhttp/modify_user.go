package xhttp

import (
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"

	"github.com/gin-gonic/gin"
)

type modifyUserData struct {
	ID       int    `json:"userid"`
	Type     int    `json:"usertype"`
	Password string `json:"password"`
	Parent   int    `json:"zoneid"`
}

type modifyUserRequest struct {
	Token string         `json:"token"`
	Data  modifyUserData `json:"data"`
}

func modifyUser(c *gin.Context) {
	var request modifyUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	user, err := model.GetUserByID(request.Data.ID, xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "user not exist"})
		return
	}

	user.Type = request.Data.Type
	user.Parent = request.Data.Parent
	if len(request.Data.Password) > 0 {
		user.Password = model.MD5(request.Data.Password)
	}

	model.UpdateUser(user, xdb.Engine)

	c.JSON(http.StatusOK, gin.H{"result": 0, "message": "OK"})
}
