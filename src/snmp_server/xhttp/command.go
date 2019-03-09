package xhttp

import (
	"bytes"
	"errors"
	"net/http"
	"os/exec"
	"snmp_server/model"
	"snmp_server/sessions"
	"strings"

	"github.com/gin-gonic/gin"
)

type runCommandData struct {
	Command string `json:"command"`
	Params  string `json:"params"`
}
type runCommandRequest struct {
	Token string         `json:"token"`
	Data  runCommandData `json:"data"`
}

var supportCommands = []string{
	"ls",
	"pwd",
	"cd",
	"date",
	"df",
	"du",
	"reboot",
	"ifconfig",
	"netstat",
	"ping",
	"ps",
}

func isValidCommand(command string) bool {
	for _, cmd := range supportCommands {
		if cmd == command {
			return true
		}
	}
	return false
}

func toParams(param string) ([]string, error) {

	mark := 0
	res := make([]string, 0, 10)
	buf := bytes.NewBuffer(nil)
	for _, c := range param {
		if mark == 0 {
			if c == ' ' {
				res = append(res, buf.String())
				buf = bytes.NewBuffer(nil)
			} else if c == '"' {
				res = append(res, buf.String())
				buf = bytes.NewBuffer(nil)
				mark = 1
			} else {
				buf.WriteRune(c)
			}
		} else if mark == 1 {
			if c == '"' {
				res = append(res, buf.String())
				buf = bytes.NewBuffer(nil)
				mark = 0
			} else {
				buf.WriteRune(c)
			}
		}
	}

	res = append(res, buf.String())

	ret := make([]string, 0, 10)
	for _, p := range res {
		if len(strings.TrimSpace(p)) > 0 {
			ret = append(ret, strings.TrimSpace(p))
		}
	}
	if mark == 0 {
		return ret, nil
	}
	return nil, errors.New("param missing \"")
}
func runCommand(c *gin.Context) {

	var request runCommandRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	if !isValidCommand(request.Data.Command) {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "no support command"})
		return
	}
	commandParams, err := toParams(request.Data.Params)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	user, _ := sessions.GetUserSession(request.Token)
	logInfo := &model.LogInfo{
		User:     user.Username,
		NTID:     "NA",
		Event:    "runcommand",
		SubEvent: request.Data.Command,
		Info:     strings.Join(commandParams, " "),
	}
	logInfo.Insert()
	if strings.EqualFold(request.Data.Command, "reboot") {
		logInfo.CloseDB()
	}
	command := exec.Command(request.Data.Command, commandParams...)

	output, err := command.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"output": string(output),
		},
	}
	c.JSON(http.StatusOK, result)
}

type getCommandRequest struct {
	Token string `json:"token"`
}

func getCommands(c *gin.Context) {
	var request runCommandRequest
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
			"commands": supportCommands,
		},
	}
	c.JSON(http.StatusOK, result)

}
