package xhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xsnmp"
	"snmp_server/xtask"
	"time"

	"github.com/cihub/seelog"

	"github.com/gin-gonic/gin"
)

type snmpRequestData struct {
	SnmpType string                 `json:"snmp_type"`
	ItemID   int                    `json:"itemid"`
	Index    int                    `json:"index"`
	OIDs     map[string]interface{} `json:"oids"`
}

type snmpRequest struct {
	Token string          `json:"token"`
	Data  snmpRequestData `json:"data"`
}

func snmp(c *gin.Context) {
	var request snmpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	t, err := model.GetTerminalByID(request.Data.ItemID, xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if t == nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "termianl not found."})
		return
	}
	if !t.IsOnline() {

	}

	if len(request.Data.OIDs) == 0 {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "oids is empty."})
		return
	}

	switch request.Data.SnmpType {
	case "get":
		snmpResult, err := xsnmp.Default.Get(request.Data.OIDs, request.Data.Index, t.Remote())
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		} else {
			result := gin.H{
				"result":  0,
				"message": "OK",
				"data": gin.H{
					"oids": snmpResult,
				},
			}
			c.JSON(http.StatusOK, result)
		}
	case "set":
		var ftpUpgrade *xtask.Upgrade
		if _, ok := request.Data.OIDs["usl_ftp_soft_file_name"]; ok {
			//ftp 升级任务
			ftpUpgrade, err = initUpgrade(t, c)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
				return
			}
		}
		if _, ok := request.Data.OIDs["usl_ftp_save_cfg_file_name"]; ok {
			request.Data.OIDs["usl_ftp_save_cfg_file_name"] = fmt.Sprintf("%s/%s.cfg", t.NTID, time.Now().Format("20060102_150405"))
		}

		snmpResult, err := xsnmp.Default.Set(request.Data.OIDs, request.Data.Index, t.Remote())
		if err != nil {
			logInfo := &model.LogInfo{
				User:     getUsernameByToken(request.Token),
				NTID:     t.NTID,
				Event:    "snmp",
				SubEvent: "set_err",
				Info:     err.Error(),
			}
			logInfo.Insert()
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})

			if ftpUpgrade != nil {
				seelog.Warnf("ftp upgrade task Delete  %s for snmp set error", ftpUpgrade.NTID)
				ftpUpgrade.Delete(xdb.EngineTask)
			}

		} else {
			result := gin.H{
				"result":  0,
				"message": "OK",
				"data": gin.H{
					"oids": snmpResult,
				},
			}
			jdata, _ := json.Marshal(snmpResult)
			logInfo := &model.LogInfo{
				User:     getUsernameByToken(request.Token),
				NTID:     t.NTID,
				Event:    "snmp",
				SubEvent: "set_ok",
				Info:     string(jdata),
			}
			logInfo.Insert()
			c.JSON(http.StatusOK, result)
		}
	default:
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "not support snmp_type ,only get|set"})
		return
	}
}

func initUpgrade(t *model.Terminal, c *gin.Context) (*xtask.Upgrade, error) {
	upgrade := &xtask.Upgrade{NTID: t.NTID}
	has, err := upgrade.Get(xdb.EngineTask)
	if err != nil {
		return nil, err
	}
	if has {
		if upgrade.Completed {
			seelog.Infof("delete completed upgrade task %s", upgrade.NTID)
			upgrade.Delete(xdb.EngineTask)
			upgrade = &xtask.Upgrade{NTID: t.NTID}
		} else {
			return nil, fmt.Errorf("ftp task %v ", xtask.UpgradeResult(upgrade.Result))
		}
	}
	upgrade.TID = t.ID
	upgrade.NTID = t.NTID
	upgrade.Name = t.Name
	upgrade.Path = t.Path
	upgrade.Type = t.Type
	upgrade.Result = int(xtask.Downing)
	err = upgrade.Insert(xdb.EngineTask)
	if err != nil {
		return nil, err
	}
	return upgrade, nil
}
