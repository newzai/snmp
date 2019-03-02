package xhttp

import (
	"net/http"
	"snmp_server/model"
	"strings"

	"github.com/gin-gonic/gin"
)

//LogWheres 查询条件
type logWheres struct {
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end,omitempty"`
	User      string `json:"username,omitempty"`
	NTID      string `json:"ntid,omitempty"`
	Event     string `json:"event,omitempty"`
	SubEvent  string `json:"sub_event,omitempty"`
}

//LogRequestData 请求data
type logRequestData struct {
	Wheres    logWheres `json:"wheres"`
	PageIndex int       `json:"index"`
	PageSize  int       `json:"size"`
}

type logRequest struct {
	Token string         `json:"token"`
	Data  logRequestData `json:"data"`
}

func getlogs(c *gin.Context) {

	var request logRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	//2006-02
	var yearMonth string
	if len(request.Data.Wheres.TimeStart) > 0 {
		yearMonth = request.Data.Wheres.TimeStart[:7]
	} else if len(request.Data.Wheres.TimeEnd) > 0 {
		yearMonth = request.Data.Wheres.TimeStart[:7]
	} else {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "missing where time_start or time_end"})
		return
	}
	if len(request.Data.Wheres.TimeStart) > 0 && len(request.Data.Wheres.TimeEnd) > 0 {
		if !strings.EqualFold(request.Data.Wheres.TimeStart[:7], request.Data.Wheres.TimeEnd[:7]) {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": "time_start and time_end must at the same month"})
			return
		}
	}
	engine, err := model.GetLogEngine(yearMonth)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	wheres := make([]string, 0, 10)
	args := make([]interface{}, 0, 10)
	if len(request.Data.Wheres.TimeStart) > 0 {
		wheres = append(wheres, "ts >= ?")
		args = append(args, request.Data.Wheres.TimeStart)
	}
	if len(request.Data.Wheres.TimeEnd) > 0 {
		wheres = append(wheres, "ts <= ?")
		args = append(args, request.Data.Wheres.TimeEnd)
	}
	if len(request.Data.Wheres.User) > 0 {
		wheres = append(wheres, "user = ?")
		args = append(args, request.Data.Wheres.User)
	}
	if len(request.Data.Wheres.NTID) > 0 {
		wheres = append(wheres, "ntid = ?")
		args = append(args, request.Data.Wheres.NTID)
	}
	if len(request.Data.Wheres.Event) > 0 {
		wheres = append(wheres, "event = ?")
		args = append(args, request.Data.Wheres.Event)
	}
	if len(request.Data.Wheres.SubEvent) > 0 {
		wheres = append(wheres, "sub_event = ?")
		args = append(args, request.Data.Wheres.SubEvent)

	}

	where := strings.Join(wheres, " and ")
	pageSize := request.Data.PageSize
	offset := (request.Data.PageSize) * (request.Data.PageIndex - 1)
	var logInfo model.LogInfo
	rows, err := engine.Where(where, args...).Limit(pageSize, offset).Rows(logInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	defer rows.Close()
	var logs []*model.LogInfo
	for rows.Next() {
		tmp := new(model.LogInfo)
		rows.Scan(tmp)
		logs = append(logs, tmp)
	}
	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"logs": logs,
		},
	}
	c.JSON(http.StatusOK, result)
}
