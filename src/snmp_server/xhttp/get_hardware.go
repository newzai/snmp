package xhttp

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"snmp_server/globalvars"
	"strings"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

type getHardwareData struct {
	Type string `json:"dev_type"`
}

type getHardwareRequest struct {
	Token string          `json:"token"`
	Data  getHardwareData `json:"data"`
}

func getAllHardware(c *gin.Context) {

	var request getHardwareRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	configDir := fmt.Sprintf("%s%s/", globalvars.FTPDir, request.Data.Type)
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
