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

//GetFTPChown chown for ftp user:group
func GetFTPChown() string {
	return fmt.Sprintf("%s:%s", FTPUser, FTPGroup)
}
