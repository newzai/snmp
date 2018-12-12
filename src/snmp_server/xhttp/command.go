package xhttp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type runCommandData struct {
	Command int    `json:"command"`
	Params  string `json:"params"`
	Hash    string `json:"hash"`
}
type runCommandRequest struct {
	Token string         `json:"token"`
	Data  runCommandData `json:"data"`
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

	commandData, hash, err := getCommand(request.Data.Command)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	if hash != request.Data.Hash {
		c.JSON(http.StatusOK, gin.H{"result": 2, "message": "hash not match, refresh commands list"})
		return
	}

	commandParams := strings.Split(commandData, " ")
	commandName := commandParams[0]
	commandParams = commandParams[1:]
	if len(request.Data.Params) > 0 {
		userParams := strings.Split(request.Data.Params, " ")
		for _, param := range userParams {
			p := strings.TrimSpace(param)
			if len(p) > 0 {
				commandParams = append(commandParams, p)
			}
		}
	}

	command := exec.Command(commandName, commandParams...)

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

	commands, hash := getAllCommands()
	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"commands": commands,
			"hash":     hash,
		},
	}
	c.JSON(http.StatusOK, result)

}

func getCommand(index int) (string, string, error) {
	commandData, err := ioutil.ReadFile("command.txt")
	if err != nil {
		return "", "", err
	}

	hash := md5.New()
	hash.Write(commandData)

	lines := strings.Split(string(commandData), "\n")
	for _, line := range lines {

		if len(line) > 0 {
			commandLine := strings.Split(line, ":")
			if len(commandLine) == 3 {
				cIndex, err := strconv.Atoi(commandLine[0])
				if err == nil && cIndex == index {
					return commandLine[2], hex.EncodeToString(hash.Sum(nil)), nil
				}
			}

		}
	}
	return "", "", fmt.Errorf("command not found by index[%d]", index)
}

func getAllCommands() (map[int]string, string) {

	commands := make(map[int]string)
	commandData, err := ioutil.ReadFile("command.txt")
	if err != nil {
		return commands, ""
	}

	lines := strings.Split(string(commandData), "\n")
	for _, line := range lines {

		if len(line) > 0 {
			commandLine := strings.Split(line, ":")
			if len(commandLine) == 3 {
				cIndex, err := strconv.Atoi(commandLine[0])
				if err == nil {
					commands[cIndex] = commandLine[1]
				}
			}

		}
	}

	hash := md5.New()
	hash.Write(commandData)

	return commands, hex.EncodeToString(hash.Sum(nil))
}
