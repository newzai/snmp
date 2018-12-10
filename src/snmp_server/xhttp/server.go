package xhttp

import (
	"fmt"
	"os"
	"path/filepath"

	"snmp_server/asset"

	"github.com/gin-gonic/gin"
)

//Run run http server
func Run(httpPort int) {

	isSuccess := true
	dirs := []string{"dist"} // 设置需要释放的目录
	for _, dir := range dirs {
		fmt.Println("remove ", filepath.Join("./", dir))
		os.RemoveAll(filepath.Join("./", dir))
	}
	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := asset.RestoreAssets("./", dir); err != nil {
			isSuccess = false
			fmt.Println("RestoreAssets error: ", filepath.Join("./", dir), " ", err.Error())
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("./", dir))
		}
	}

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
		v1.POST("/get_warnings", getWarning)
		v1.POST("/warning_test", warningTest)
		v1.POST("/get_commands", getCommands)
		v1.POST("/run_command", runCommand)
	}

	r.Run(fmt.Sprintf(":%d", httpPort))
}
