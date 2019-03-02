package xtraphandler

import (
	"net"
	"snmp_server/mibs/report"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xsnmp"
	"snmp_server/xtrap"
	"time"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//Keepalive keepalive
func Keepalive(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	seelog.Info("recv OID_KEEPALIVE from ", remote)
	msg := new(report.Keepalive)
	msg.FromSnmpPackage(packet)
	seelog.Infof("msg:%s", msg)

	t, err := model.GetTerminalByNTID(msg.NTID, xdb.Engine)
	if err != nil {
		seelog.Errorf("get terminal error:%s", err)
	}
	if t == nil {

		seelog.Warnf("device is not exist %s, usl_register_again", msg.NTID)
		var RegisterAgain = map[string]interface{}{
			"usl_register_again": "register",
		}
		xsnmp.Default.Set(RegisterAgain, 0, remote)
		return
	}
	if !t.IsOnline() {

		seelog.Warnf("device is offline %s, usl_register_again", msg.NTID)
		var RegisterAgain = map[string]interface{}{
			"usl_register_again": "register",
		}
		xsnmp.Default.Set(RegisterAgain, 0, remote)
		return
	}
	setKeepalive(msg.NTID, remote.String())
	t.LastKeepalive = time.Now()
	t.IP = remote.IP.String()
	t.Port = remote.Port
	err = model.UpdateTerminal(t, true, xdb.Engine)
	if err != nil {
		seelog.Warnf("Update terminal error:%s", err)
	}

	TTLResult := map[string]interface{}{
		"usl_keeplive_result": 1,
	}
	result, err := xsnmp.Default.Set(TTLResult, 0, remote)
	seelog.Infof("set result :%v, %v", result, err)
	tryCreateFTPDir(t)

}

func init() {

	xtrap.RegisterHandler(trapTypeKEEPALIVE, Keepalive)
}
