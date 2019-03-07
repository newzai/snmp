package xhttp

import (
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xwarning"

	"github.com/cihub/seelog"

	"github.com/gin-gonic/gin"
)

type getitemData struct {
	ItemID   int `json:"itemid"`
	ItemType int `json:"itemtype"`
}

type getitemRequest struct {
	Token string      `json:"token"`
	Data  getitemData `json:"data"`
}

type itemInfo struct {
	ItemID   int    `json:"itemid"`
	Parent   int    `json:"parent"`
	ItemName string `json:"itemname"`
	ItemPath string `json:"itempath"`
	ItemType int    `json:"itemtype"`
	Status   int    `json:"status"`
	DevType  string `json:"dev_type"`
	Warnings int    `json:"warnings"`
}

func getitem(c *gin.Context) {
	seelog.Info("getitem.")
	var request getitemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	switch request.Data.ItemType {
	case 1:
		var items []itemInfo
		zones, err := model.GetZonesByParent(request.Data.ItemID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			break
		}
		for _, zone := range zones {
			item := itemInfo{
				ItemID:   zone.ID,
				Parent:   zone.Parent,
				ItemName: zone.Name,
				ItemPath: zone.Path,
				ItemType: 1,
			}
			items = append(items, item)
		}

		terminals, _ := model.GetTerminalByParent(request.Data.ItemID, xdb.Engine)
		for _, terminal := range terminals {
			status := 0
			if terminal.IsOnline() {
				status = 1
			}
			warnings := 0
			if status == 1 {
				warnings = xwarning.Stats.GetCounts(terminal.ID)
			}
			item := itemInfo{
				ItemID:   terminal.ID,
				Parent:   terminal.Parent,
				ItemName: terminal.Name,
				ItemPath: terminal.Path,
				ItemType: 2,
				Status:   status,
				DevType:  terminal.Type,
				Warnings: warnings,
			}
			items = append(items, item)
		}
		result := gin.H{
			"result":  0,
			"message": "OK",
			"data": gin.H{
				"items": items,
			},
		}
		c.JSON(http.StatusOK, result)
	case 2:
		terminal, err := model.GetTerminalByID(request.Data.ItemID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			break
		}
		var items []itemInfo
		status := 0
		if terminal != nil {
			if terminal.IsOnline() {
				status = 1
			}
			warnings := 0
			if status == 1 {
				warnings = xwarning.Stats.GetCounts(terminal.ID)
			}
			item := itemInfo{
				ItemID:   terminal.ID,
				Parent:   terminal.Parent,
				ItemName: terminal.Name,
				ItemPath: terminal.Path,
				ItemType: 2,
				Status:   status,
				DevType:  terminal.Type,
				Warnings: warnings,
			}
			items = append(items, item)
		}
		result := gin.H{
			"result":  0,
			"message": "OK",
			"data": gin.H{
				"items": items,
			},
		}
		c.JSON(http.StatusOK, result)

	case 3:
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "not support user itemtype"})
	default:
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "not support itemtype"})
	}
}

type getAllTerminalsRequest struct {
	Token string `json:"token"`
}

func getAllTerminals(c *gin.Context) {
	var request getAllTerminalsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	terminals, err := model.GetAllTerminals(xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	tsinfos := make(map[string]string)
	for _, t := range terminals {
		tsinfos[t.Path+"."+t.Name] = t.NTID
	}

	result := gin.H{
		"result":  0,
		"message": "OK",
		"data":    tsinfos,
	}
	c.JSON(http.StatusOK, result)

}
