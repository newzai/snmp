package xhttp

import (
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"

	"github.com/gin-gonic/gin"
)

type setItemData struct {
	ID         int                    `json:"itemid"`
	Type       int                    `json:"itemtype"`
	Attributes map[string]interface{} `json:"attrs"`
}

type setItemRequest struct {
	Token string      `json:"token"`
	Data  setItemData `json:"data"`
}

func setItem(c *gin.Context) {

	var request setItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	switch request.Data.Type {
	case 1:
		zone, err := model.GetZoneByID(request.Data.ID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return
		}
		if value, ok := request.Data.Attributes["imageurl"]; ok {
			switch tv := value.(type) {
			case string:
				zone.ImageURL = tv
			default:
				c.JSON(http.StatusOK, gin.H{"result": 1, "message": "imageurl is not string"})
				return
			}
		}
		xValue, xok := request.Data.Attributes["x"]
		yValue, yok := request.Data.Attributes["y"]
		if (xok && !yok) || (!xok && yok) {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": "x,y must include at the same time"})
			return
		}
		if xok {
			switch x := xValue.(type) {
			case float64:
				zone.X = int(x)
			default:
				c.JSON(http.StatusOK, gin.H{"result": 1, "message": "x is not int"})
				return
			}
		}
		if yok {
			switch y := yValue.(type) {
			case float64:
				zone.Y = int(y)
			default:
				c.JSON(http.StatusOK, gin.H{"result": 1, "message": "y is not int"})
				return
			}
		}
		err = model.UpdateZone(zone, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return
		}
	case 2:
		t, err := model.GetTerminalByID(request.Data.ID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return
		}
		xValue, xok := request.Data.Attributes["x"]
		yValue, yok := request.Data.Attributes["y"]
		if (xok && !yok) || (!xok && yok) {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": "x,y must include at the same time"})
			return
		}
		if xok {
			switch x := xValue.(type) {
			case float64:
				t.X = int(x)
			default:
				c.JSON(http.StatusOK, gin.H{"result": 1, "message": "x is not int"})
				return
			}
		}
		if yok {
			switch y := yValue.(type) {
			case float64:
				t.Y = int(y)
			default:
				c.JSON(http.StatusOK, gin.H{"result": 1, "message": "y is not int"})
				return
			}
		}

		err = model.UpdateTerminalXY(t, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return
		}
	default:
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "not support itemtype"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": 0, "message": "OK"})
}
