package xhttp

import (
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"

	"github.com/gin-gonic/gin"
)

type getUsersData struct {
	ID int `json:"zoneid,omitempty"`
}

type getUsersRequest struct {
	Token string       `json:"token"`
	Data  getUsersData `json:"data"`
}

func getUsers(c *gin.Context) {

	var request getUsersRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	if request.Data.ID > 0 {
		parent, err := model.GetZoneByID(request.Data.ID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return

		}
		users, err := model.GetUsersByParent(request.Data.ID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return
		}
		userinfos := make([]*userInfo, 0, len(users))
		for _, user := range users {
			info := new(userInfo)
			info.UserID = user.ID
			info.UserType = user.Type
			info.Username = user.Username
			info.Password = user.Password
			info.ZoneInfo.ID = parent.ID
			info.ZoneInfo.Name = parent.Name
			info.ZoneInfo.Path = parent.Path
			userinfos = append(userinfos, info)
		}
		result := gin.H{
			"result":  0,
			"message": "OK",
			"data": gin.H{
				"users": userinfos,
			},
		}
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "missing zoneid "})
	}
}

type getAllUsersRequest struct {
	Token string `json:"token"`
}

func getAllUsers(c *gin.Context) {
	var request getAllUsersRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	users, err := model.GetAllUsers(xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	usernames := make([]string, 0, len(users)+2)
	usernames = append(usernames, "system")
	usernames = append(usernames, "ntp")
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}

	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"usernames": usernames,
		},
	}
	c.JSON(http.StatusOK, result)
}
