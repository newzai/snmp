package xhttp

import (
	"fmt"
	"os"
	"path/filepath"
	"snmp_server/asset"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

//Run run http server
func Run(httpPort int) {

	restoreAssets()

	r := gin.Default()
	r.StaticFile("/", "./dist/index.html")
	r.Static("/static", "./dist/static")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	v1 := r.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/logout", logout)
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
		v1.POST("/clear_warning", clearWarning)
		v1.POST("/warning_test", warningTest)
		v1.POST("/get_commands", getCommands)
		v1.POST("/run_command", runCommand)
		v1.POST("/get_configure", getConfigure)
		v1.POST("/set_configure", setConfigure)
		v1.POST("/get_logs", getlogs)
		v1.POST("/system_check", systemCheck)
		v1.POST("/get_all_users", getAllUsers)
		v1.POST("/get_all_terminals", getAllTerminals)
	}
	pprof.Register(r)

	r.Run(fmt.Sprintf(":%d", httpPort))
}

//restoreAssets 解压web静态资源
func restoreAssets() {
	isSuccess := true
	dirs := []string{"dist"} // 设置需要释放的目录
	for _, dir := range dirs {
		os.RemoveAll(filepath.Join("./", dir))
	}
	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := asset.RestoreAssets("./", dir); err != nil {
			isSuccess = false
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("./", dir))
		}
	}
}
