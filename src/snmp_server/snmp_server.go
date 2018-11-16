package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"reflect"
	"snmp_server/globalvars"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xhttp"
	"snmp_server/xsnmp"
	"snmp_server/xtask"
	"snmp_server/xtrap"
	"snmp_server/xwarning"
	"strconv"
	"xconf"

	"github.com/cihub/seelog"

	_ "snmp_server/mibs"
	_ "snmp_server/xtrap/xtraphandler"
)

var ftpDir = flag.String("ftpdir", "/home/klsnmp/ftpfile/", "-ftpdir=/home/klsnmp/ftpfile")
var snmpPort = flag.Uint("snmp_port", 162, "-snmp_port=162")
var httpPort = flag.Int("http_port", 9192, "-http_port=9192")
var showSQL = flag.Bool("show_sql", false, "-show_sql=false|true")
var ftpUser = flag.String("ftp_user", "uftp", "-ftp_user=uftp")
var ftpGroup = flag.String("ftp_group", "uftp", "-ftp_user=uftp")
var help = flag.Bool("help", false, "-help=true")
var version = flag.Bool("version", false, "-version")

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	if *version {
		fmt.Printf("Version: %v\n", globalvars.AppVersion)
		fmt.Printf("Git Hash: %v\n", globalvars.AppGitHash)
		fmt.Printf("Build Time: %v\n", globalvars.AppBuildTime)
		fmt.Printf("Go Version: %v\n", globalvars.GoVersion)
		os.Exit(0)
	}
	globalvars.FTPDir = *ftpDir
	globalvars.FTPUser = *ftpUser
	globalvars.FTPGroup = *ftpGroup

	// start ftp dir, user, group check
	ftpCheck()
	//end ftp dir, ftp user , ftp group check

	xconf.InitSeelog("./conf/", "./log/", "snmp", false, true, xconf.Info)
	seelog.Info("snmp server start..")
	xdb.Init(*showSQL)
	model.InitDatabase(xdb.Engine)
	xtask.InitDatabase(xdb.EngineTask)
	xwarning.InitDatabase(xdb.EngineWarning)
	go xhttp.Run(*httpPort)

	xsnmp.Default.Start("0.0.0.0", uint16(*snmpPort), xtrap.OnTrapHandler)
}

func ftpCheck() {
	{
		fileinfo, err := os.Stat(globalvars.FTPDir)
		if err != nil {
			if os.IsNotExist(err) {
				flag.Usage()
				panic(err)
			}
		} else {
			if !fileinfo.IsDir() {
				flag.Usage()
				panic(globalvars.FTPDir + " is not dir.")
			}
		}
		u, err := user.Lookup(globalvars.FTPUser)
		if err != nil {
			flag.Usage()
			panic(err)
		}
		g, err := user.LookupGroup(globalvars.FTPGroup)
		if err != nil {
			flag.Usage()
			panic(err)
		}
		uid, _ := strconv.Atoi(u.Uid)
		gid, _ := strconv.Atoi(g.Gid)

		dirUID := int(reflect.ValueOf(fileinfo.Sys()).Elem().FieldByName("Uid").Uint())
		dirGid := int(reflect.ValueOf(fileinfo.Sys()).Elem().FieldByName("Gid").Uint())

		if dirGid != gid || dirUID != uid {
			flag.Usage()
			panic("ftp dir  own uid or gid not match.")
		}
	}
}
