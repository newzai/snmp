package xhttp

import (
	"net/http"
	"snmp_server/globalvars"
	"snmp_server/model"

	"github.com/gin-gonic/gin"
)

type getConfigureRequest struct {
	Token string `json:"token"`
}

func getConfigure(c *gin.Context) {
	var request getConfigureRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"configure": globalvars.Default,
		},
	}
	c.JSON(http.StatusOK, result)
}

type setConfigureData struct {
	Configure       globalvars.Configure `json:"configure"`
	DiskUsedPrecent float64              `json:"disk_used_precent"`
}

type setConfigureRequest struct {
	Token string           `json:"token"`
	Data  setConfigureData `json:"data"`
}

func setConfigure(c *gin.Context) {

	var request setConfigureRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	globalvars.Default.NTPServer1 = request.Data.Configure.NTPServer1
	globalvars.Default.NTPServer2 = request.Data.Configure.NTPServer2
	globalvars.Default.NTPDEnable = request.Data.Configure.NTPDEnable
	globalvars.Default.WebPort = request.Data.Configure.WebPort
	globalvars.Default.SnmpPort = request.Data.Configure.SnmpPort
	if request.Data.DiskUsedPrecent > 1.0 {
		if request.Data.DiskUsedPrecent > 100.0 {
			globalvars.SetDiskHighUsedPrecent(90.0)
		} else {
			globalvars.SetDiskHighUsedPrecent(request.Data.DiskUsedPrecent)
		}
	} else {
		globalvars.SetDiskHighUsedPrecent(90.0)
	}

	err := globalvars.Default.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"configure":         globalvars.Default,
			"disk_used_precent": globalvars.GetDiskHighUsedPrecent(),
		},
	}

	logInfo := &model.LogInfo{
		User:     getUsernameByToken(request.Token),
		NTID:     "NA",
		Event:    "config",
		SubEvent: "set_config",
		Info:     globalvars.Default.String(),
	}
	logInfo.Insert()
	c.JSON(http.StatusOK, result)
}
