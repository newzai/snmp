package xhttp

import (
	"net/http"
	"snmp_server/sessions"

	"github.com/gin-gonic/gin"
)

type logoutRequest struct {
	Token string `json:"token"`
}

func logout(c *gin.Context) {
	var request logoutRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	sessions.Logout(request.Token)
	c.JSON(http.StatusOK, gin.H{"result": 0, "message": "OK"})

}
