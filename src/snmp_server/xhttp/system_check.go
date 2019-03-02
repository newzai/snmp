package xhttp

import (
	"fmt"
	"net/http"
	"os"
	"snmp_server/globalvars"
	"time"

	"github.com/shirou/gopsutil/process"

	"github.com/shirou/gopsutil/disk"

	"github.com/gin-gonic/gin"
)

type systemCheckRequest struct {
	Token string `json:"token"`
}

func systemCheck(c *gin.Context) {
	var request systemCheckRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	diskInfo, err := disk.Usage("/")
	os.Remove("./testfile.txt")
	var diskNormal bool
	if err == nil {
		diskNormal = true
	}
	var free uint64
	if diskNormal {
		free = diskInfo.Free
	}
	checkPid := os.Getpid() // process.test
	selfProcess, _ := process.NewProcess(int32(checkPid))
	cpuPercent, _ := selfProcess.CPUPercent()
	memoryPercent, _ := selfProcess.MemoryPercent()
	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"start_time":     globalvars.StartTime.Format("2006-01-02 15:04:05"),
			"run_time":       fmt.Sprintf("%f hours", time.Since(globalvars.StartTime).Hours()),
			"disk_status":    diskNormal,
			"disk_free":      fmt.Sprintf("%d G", free/(1024*1024*1024)),
			"net_status":     true,
			"cpu_percent":    fmt.Sprintf("%f %%", cpuPercent),
			"memory_percent": fmt.Sprintf("%f %%", memoryPercent),
		},
	}
	c.JSON(http.StatusOK, result)

}
