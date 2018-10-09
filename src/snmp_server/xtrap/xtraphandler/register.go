package xtraphandler

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"snmp_server/globalvars"
	"snmp_server/mibs"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xsnmp"
	"snmp_server/xtrap"
	"strings"
	"time"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//Register for trap register
func Register(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.Register)
	msg.FromSnmpPackage(packet)
	if !strings.HasPrefix(msg.Zone, "root.") {
		msg.Zone = "root." + msg.Zone
	}
	RegisterTTL := map[string]interface{}{
		"usl_register_result": 1,
		"usl_register_ttl":    15,
	}
	result, err := xsnmp.Default.Set(RegisterTTL, 0, remote)
	seelog.Infof("set result :%v, %v", result, err)

	t, err := model.GetTerminalByNTID(msg.NTID, xdb.Engine)
	if err != nil {
		seelog.Error("err:", err)
	}
	parent, err := model.CreatePath(msg.Zone, xdb.Engine)
	rawIP := fmt.Sprintf("%s:%d", msg.DeviceIP, msg.DevicePort)
	natIP := fmt.Sprintf("%s:%d", remote.IP.String(), remote.Port)
	if t == nil {
		if err != nil {
			panic(err)
		}
		t = &model.Terminal{
			ID:            0,
			Name:          msg.ProcductName,
			Path:          msg.Zone,
			NTID:          msg.NTID,
			IP:            msg.DeviceIP,
			Port:          msg.DevicePort,
			Version:       msg.ProcductVersion,
			Type:          msg.ProcductType,
			Parent:        parent,
			LastKeepalive: time.Now(),
		}

		if rawIP != natIP {
			seelog.Warnf("is Nat for %s  raw(%s) nat(%s)", t.NTID, rawIP, natIP)
			t.IP = remote.IP.String()
			t.Port = remote.Port
		}
		seelog.Infof("new terminal:%v", t)
		err = model.CreateTerminal(t, xdb.Engine)
		if err != nil {
			panic(err)
		}
		seelog.Infof("New Terminal %s ID:%s", t.NTID, t.ID)

	} else {
		t.Name = msg.ProcductName
		t.Path = msg.Zone
		t.IP = msg.DeviceIP
		t.Port = msg.DevicePort
		t.Version = msg.ProcductVersion
		t.Type = msg.ProcductType
		t.LastKeepalive = time.Now()
		t.Parent = parent
		if rawIP != natIP {
			seelog.Warnf("is Nat for %s  raw(%s) nat(%s)", t.NTID, rawIP, natIP)
			t.IP = remote.IP.String()
			t.Port = remote.Port
		}
		seelog.Infof("update terminal:%v", t)
		err = model.UpdateTerminal(t, false, xdb.Engine)
		if err != nil {
			seelog.Warnf("update %s error %s", t.NTID, err)
		}
	}

	tryCreateFTPDir(t)
}

func tryCreateFTPDir(t *model.Terminal) {
	ftpDir := fmt.Sprintf("%s%s", globalvars.FTPDir, t.NTID)
	ftpDirFile, err := os.Open(ftpDir)

	if err == nil {
		ftpDirFile.Close()
	} else {
		if os.IsNotExist(err) {
			err = os.MkdirAll(ftpDir, os.ModeDir|0777)
			if err != nil {
				seelog.Warnf("create %s ftp dir err %s", t.NTID, err)
			} else {
				cmd := exec.Command("chown", globalvars.GetFTPChown(), ftpDir)
				err = cmd.Run()
				if err != nil {
					seelog.Warnf("chown %s ftp dir error %s", t.NTID, err)
				}
			}
		}
	}
}
func init() {
	xtrap.RegisterHandler(trapTypeREGISTER, Register)
}
