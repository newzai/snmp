package xhttp

import (
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"

	"github.com/gin-gonic/gin"
)

type createUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserType int    `json:"usertype"`
	Parent   int    `json:"zoneid"`
}

type createUserRequest struct {
	Token string         `json:"token"`
	Data  createUserData `json:"data"`
}

func createUser(c *gin.Context) {

	var request createUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	user, err := model.GetUserByName(request.Data.Username, xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if user != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "user has exist"})
		return
	}

	if !(request.Data.UserType == 1 || request.Data.UserType == 2) {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "usertype error must 1 or 2"})
		return
	}

	zone, err := model.GetZoneByID(request.Data.Parent, xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if zone == nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "zoneid(parent) not exist"})
		return
	}

	user = &model.User{
		Username: request.Data.Username,
		Password: model.MD5(request.Data.Password),
		Type:     request.Data.UserType,
		Parent:   request.Data.Parent,
	}

	_, err = model.CreateUser(user, xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": 0, "message": "OK"})
}
