package globalvars

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/cihub/seelog"

	"github.com/beevik/ntp"
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

//Configure config for snmp_server
type Configure struct {
	NTPDEnable bool   `json:"ntpd_enable"`
	NTPServer1 string `json:"ntp_server_1"`
	NTPServer2 string `json:"ntp_server_2"`
	WebPort    int    `json:"web_port"`
	SnmpPort   int    `json:"snmp_port"`
}

//Exec  后端线程
func (c *Configure) Exec() {

	seelog.Infof("Start Configure, NTPDEnable %t", c.NTPDEnable)
	for {
		select {
		case <-time.After(60 * time.Second):
			seelog.Infof("NTP Enable %t", c.NTPDEnable)
			if c.NTPDEnable {
				time, err := ntp.Time(c.NTPServer1)
				if err == nil {
					seelog.Infof("get time (%s) from NTPServer1(%s)", time.Format("01/02/2006 15:04:05"), c.NTPServer1)
					cmd := exec.Command("date", "-s", time.Format("01/02/2006 15:04:05"))
					cmd.Run()
				} else {
					time, err = ntp.Time(c.NTPServer2)
					if err == nil {
						seelog.Infof("get time (%s) from NTPServer2(%s)", time.Format("01/02/2006 15:04:05"), c.NTPServer2)
						cmd := exec.Command("date", "-s", time.Format("01/02/2006 15:04:05"))
						cmd.Run()
					} else {
						seelog.Errorf("get time error:%v", err)
					}
				}
			}
		}
	}
}

//Load load config form local
func (c *Configure) Load() {

	jdata, err := ioutil.ReadFile("configure.json")
	if err == nil {
		json.Unmarshal(jdata, c)
	} else {
		c.Save()
	}
}

//Save save to local
func (c *Configure) Save() error {

	jdata, _ := json.MarshalIndent(c, "", " ")
	err := ioutil.WriteFile("configure.json", jdata, os.ModePerm)

	return err
}

//Default default configure
var Default *Configure

func init() {

	Default = &Configure{
		NTPDEnable: false,
		WebPort:    9192,
		SnmpPort:   162,
	}

	Default.Load()
	go Default.Exec()
}
