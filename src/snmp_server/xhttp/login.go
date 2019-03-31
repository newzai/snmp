package xhttp

import (
	"fmt"
	"net/http"
	"snmp_server/globalvars"
	"snmp_server/model"
	"snmp_server/sessions"
	"snmp_server/xdb"
	"strings"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*{
    "userid":int
    "usertype":int  -- 1 管理员， 2 普通用户
    "username": ""
    "password": ""  -- md5之后的值
    zoneinfo:
    {
        "zoneid": int   -- top(root) zone
        "zonename": ""  -- top(root) zone name
    }
}*/
type zoneInfo struct {
	ID    int    `json:"zoneid"`
	Name  string `json:"zonename"`
	Path  string `json:"zonepath"`
	Image string `json:"imageurl"`
}
type userInfo struct {
	UserID   int      `json:"userid"`
	UserType int      `json:"usertype"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	ZoneInfo zoneInfo `json:"zoneinfo"`
}

func login(c *gin.Context) {

	var login loginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	user, err := model.GetUserByName(login.Username, xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "user not exist"})
		return
	}
	if !strings.EqualFold(model.MD5(login.Password), user.Password) {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "password error"})
		return
	}
	parent, _ := model.GetZoneByID(user.Parent, xdb.Engine)
	if parent == nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": fmt.Sprintf("user parent(%d) not exist", user.Parent)})
		return
	}
	token := sessions.AllocateSession(user)

	info := new(userInfo)
	info.UserID = user.ID
	info.UserType = user.Type
	info.Username = user.Username
	info.Password = user.Password
	info.ZoneInfo.ID = parent.ID
	info.ZoneInfo.Name = parent.Name
	info.ZoneInfo.Path = parent.Path
	info.ZoneInfo.Image = parent.ImageURL

	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"token":      token,
			"userinfo":   info,
			"version":    globalvars.AppVersion,
			"build_date": globalvars.AppBuildTime,
			"git_hash":   globalvars.AppGitHash,
			"go_version": globalvars.GoVersion,
		},
	}
	c.JSON(http.StatusOK, result)

	logInfo := fmt.Sprintf("remote:%s, token:%s", c.Request.RemoteAddr, token)
	model.UserLoginLog(user.Username, logInfo)
}
