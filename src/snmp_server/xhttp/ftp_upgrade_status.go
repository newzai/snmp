package xhttp

import (
	"net/http"
	"snmp_server/xdb"
	"snmp_server/xtask"

	"github.com/gin-gonic/gin"
)

type ftpUpgradeStatusData struct {
	ID   int    `json:"itemid,omitempty"`
	Path string `json:"itempath,omitempty"`
	Type string `json:"dev_type,omitempty"`
}

type ftpUpgradeStatusRequest struct {
	Token string               `json:"token"`
	Data  ftpUpgradeStatusData `json:"data"`
}

func ftpUpgradeStatus(c *gin.Context) {

	var request ftpUpgradeStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	if request.Data.ID > 0 {
		upgrade := &xtask.Upgrade{TID: request.Data.ID}
		has, err := upgrade.Get(xdb.EngineTask)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return
		}
		if !has {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": "not ftp upgrade task"})
			return
		}
		result := gin.H{
			"result":  0,
			"message": "OK",
			"data": gin.H{
				"upgrades": []*xtask.Upgrade{upgrade},
			},
		}
		c.JSON(http.StatusOK, result)
	} else {
		upgrade := &xtask.Upgrade{Path: request.Data.Path, Type: request.Data.Type}
		upgrades, err := upgrade.Find(xdb.EngineTask)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return
		}
		result := gin.H{
			"result":  0,
			"message": "OK",
			"data": gin.H{
				"upgrades": upgrades,
			},
		}
		c.JSON(http.StatusOK, result)
	}
}
