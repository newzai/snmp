package xhttp

import (
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"

	"github.com/gin-gonic/gin"
)

type modifyPasswordData struct {
	ID          int    `json:"userid"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type modifyPasswordRequest struct {
	Token string             `json:"token"`
	Data  modifyPasswordData `json:"data"`
}

func modifyPassword(c *gin.Context) {

	var request modifyPasswordRequest
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

	if user.Password == request.Data.OldPassword {
		user.Password = model.MD5(request.Data.NewPassword)
		model.UpdateUser(user, xdb.EngineTask)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "user old password error"})
	}

	c.JSON(http.StatusOK, gin.H{"result": 0, "message": "OK"})
}
