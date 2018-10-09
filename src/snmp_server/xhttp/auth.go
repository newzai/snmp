package xhttp

import (
	"net/http"
	"snmp_server/sessions"

	"github.com/gin-gonic/gin"
)

type tokenParam struct {
	Token string `json:"token"`
}

func authToken(token string, c *gin.Context) bool {

	_, ok := sessions.GetUserSession(token)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"result": 401, "message": "user not login"})
		return false
	}
	return true
}
