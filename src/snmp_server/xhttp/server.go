package xhttp

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Run run http server
func Run(httpPort int) {

	r := gin.Default()

	r.Static("/", "./dist")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	v1 := r.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/getitem", getitem)
		v1.POST("/snmp", snmp)
		v1.POST("/snmp_batch", snmpBatch)
		v1.POST("/ftp_upgrade_status", ftpUpgradeStatus)
		v1.POST("/create_user", createUser)
		v1.POST("/getusers", getUsers)
		v1.POST("/delete_user", deleteUser)
		v1.POST("/modify_password", modifyPassword)
		v1.POST("/modify_user", modifyUser)
		v1.POST("/upload_hardware", uploadHardware)
		//v1.GET("/upload_hardware", uploadHardwareWeb)
		v1.POST("/get_all_hardware", getAllHardware)
		v1.POST("/get_config_file", getConfigFile)
	}

	r.Run(fmt.Sprintf(":%d", httpPort))
}
