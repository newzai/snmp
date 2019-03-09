package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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
	"syscall"
	"xconf"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"

	_ "snmp_server/mibs"
	_ "snmp_server/xtrap/xtraphandler"
)

var upgradeDir = flag.String("upgradedir", "/home/klsnmp/upgrade", "-upgradedir=/home/klsnmp/upgrade")
var ftpDir = flag.String("ftpdir", "/home/klsnmp/ftpfile/", "-ftpdir=/home/klsnmp/ftpfile")
var showSQL = flag.Bool("show_sql", false, "-show_sql=false|true")
var ftpUser = flag.String("ftp_user", "uftp", "-ftp_user=uftp")
var ftpGroup = flag.String("ftp_group", "uftp", "-ftp_user=uftp")
var help = flag.Bool("help", false, "-help=true")
var version = flag.Bool("version", false, "-version")
var debug = flag.Bool("debug", false, "-debug=true|false")
var disableFtp = flag.Bool("disable_ftp", false, "-disable_ftp")

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
	if *debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	globalvars.FTPDir = *ftpDir
	fmt.Println(*ftpDir, ",", globalvars.FTPDir)
	globalvars.FTPUser = *ftpUser
	globalvars.FTPGroup = *ftpGroup
	globalvars.UpgradeDir = *upgradeDir
	globalvars.ShowSQL = *showSQL
	model.ShowSQL = *showSQL
	// start ftp dir, user, group check
	if !*disableFtp {
		ftpCheck()
	}
	//end ftp dir, ftp user , ftp group check

	xconf.InitSeelog("./conf/", "./log/", "snmp", false, true, xconf.Info)
	seelog.Info("snmp server start..")
	model.SystemStartLog()
	xdb.Init(*showSQL)
	model.InitDatabase(xdb.Engine)
	xtask.InitDatabase(xdb.EngineTask)
	xwarning.InitDatabase(xdb.EngineWarning)
	go xhttp.Run(globalvars.Default.WebPort)

	go xsnmp.Default.Start("0.0.0.0", uint16(globalvars.Default.SnmpPort), xtrap.OnTrapHandler)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	sig := <-sigs
	model.SystemStopLog(sig.String())
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
