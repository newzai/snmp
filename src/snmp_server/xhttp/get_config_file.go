package xhttp

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"snmp_server/globalvars"
	"snmp_server/model"
	"snmp_server/xdb"
	"strings"

	"github.com/cihub/seelog"

	"github.com/gin-gonic/gin"
)

type getConfigFileData struct {
	ID int `json:"itemid"`
}
type getConfigFileRequest struct {
	Token string            `json:"token"`
	Data  getConfigFileData `json:"data"`
}

func getConfigFile(c *gin.Context) {

	var request getConfigFileRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	t, err := model.GetTerminalByID(request.Data.ID, xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	configDir := fmt.Sprintf("%s%s/", globalvars.FTPDir, t.NTID)
	files := make([]string, 0, 0)
	filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			seelog.Warn(err)
			return err
		}
		if !info.IsDir() {
			files = append(files, strings.TrimPrefix(path, globalvars.FTPDir))
		}
		return nil
	})
	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"files": files,
		},
	}
	c.JSON(http.StatusOK, result)
}
