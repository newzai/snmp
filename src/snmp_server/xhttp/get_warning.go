package xhttp

import (
	"encoding/json"
	"net/http"
	"snmp_server/model"
	"snmp_server/sessions"
	"snmp_server/xdb"
	"snmp_server/xtrap/xtraphandler"
	"snmp_server/xwarning"

	"github.com/gin-gonic/gin"
)

type getWarningRequest struct {
	Token string `json:"token"`
}

func getWarning(c *gin.Context) {

	var request getWarningRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	user, _ := sessions.GetUserSession(request.Token)
	zone, err := model.GetZoneByID(user.Parent, xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	path := ""
	if zone.Parent == 0 {
		path = zone.Name
	} else {
		path = zone.Path + "." + zone.Name
	}

	warnings, err := xwarning.GetWarningsByPath(path, xdb.EngineWarning)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"warnings": warnings,
		},
	}
	c.JSON(http.StatusOK, result)
}

type cleanWarningData struct {
	ID int64 `json:"id"` // the warning id
}

type clearWarningRequest struct {
	Token string           `json:"token"`
	Data  cleanWarningData `json:"data"`
}

func clearWarning(c *gin.Context) {

	var request clearWarningRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	w, err := xwarning.GetWarningByID(request.Data.ID, xdb.EngineWarning)
	if err == nil {
		xwarning.ClearWarning(w.NTID, w.WType, xdb.EngineWarning)
	}
	c.JSON(http.StatusOK, gin.H{"result": 0, "message": "OK"})
}

//WarningTest WarningTest
type WarningTest struct {
	NTID   string `json:"usl_ntid"`
	Type   string `json:"warning_type"`
	Clear  int    `json:"clear"`
	Status int    `json:"status"`
	Demo   string `json:"demo"`
}

func (r *WarningTest) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//WarningType get warning type
func (r *WarningTest) WarningType() string {
	return r.Type
}

//WarningStatus get warning type
func (r *WarningTest) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *WarningTest) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *WarningTest) GetNTID() string {
	return r.NTID
}

//GetDemo demo for warning
func (r *WarningTest) GetDemo() string {
	return r.Demo
}

func warningTest(c *gin.Context) {

	var request WarningTest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	xtraphandler.DoWarningTest(&request)

	c.JSON(http.StatusOK, gin.H{"result": 0, "message": "OK"})
}
