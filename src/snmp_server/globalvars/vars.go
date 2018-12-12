package globalvars

import (
	"fmt"
)

//FTPDir ftp dir
var FTPDir = "/home/klsnmp/ftpfile/"

//FTPUser ftp User
var FTPUser = "uftp"

//FTPGroup ftp group
var FTPGroup = "uftp"

//FTPUID ftp user uid
var FTPUID = 0

//FTPGID ftp group  gid
var FTPGID = 0

//UpgradeDir 防止升级文件的目录
var UpgradeDir = "/home/klsnmp/upgrade"

var (
	//AppVersion app 版本号
	AppVersion = "0.0.1"
	//AppBuildTime 编译时间
	AppBuildTime = "2017-12-01T00:03:18+0800"
	//AppGitHash git hash
	AppGitHash = "undefined"
	//GoVersion go version
	GoVersion = "undefined"
)

//GetFTPChown chown for ftp user:group
func GetFTPChown() string {
	return fmt.Sprintf("%s:%s", FTPUser, FTPGroup)
}
