package xtraphandler

import (
	"net"
	"snmp_server/mibs/report"
	"snmp_server/xdb"
	"snmp_server/xtask"
	"snmp_server/xtrap"

	"github.com/cihub/seelog"

	"github.com/soniah/gosnmp"
)

func upgradeReport(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {

	msg := new(report.UpgradeReport)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
	upgrade := xtask.Upgrade{NTID: msg.NTID}
	has, err := upgrade.Get(xdb.EngineTask)
	if err != nil {
		seelog.Errorf("upgrade.Get %s error %s ", msg.NTID, err)
		return
	}
	if !has {
		seelog.Warnf("not upgrade task %s", msg.NTID)
		return
	}
	upgrade.Result = msg.Status
	upgrade.Progress = msg.Deep
	if upgrade.Result == int(xtask.OK) || upgrade.Result == int(xtask.ERROR) {
		upgrade.Completed = true
	}
	err = upgrade.Update(xdb.EngineTask)
	if err != nil {
		seelog.Warnf("upgrade.Update %s error %s ", msg.NTID, err)
	}
	seelog.Infof("upgrade task %s Result %s", upgrade.NTID, xtask.UpgradeResult(upgrade.Result))
}
func restoreCfgReport(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(report.UpgradeRestoreCfgReport)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv restoreCfgReport:", msg)
}

func saveCfgReport(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(report.UpgradeSaveCfgReport)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv saveCfgReport:", msg)
}

func deviceRebootReport(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(report.DeviceRebootReport)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv DeviceRebootReport:", msg)
}
func init() {

	xtrap.RegisterHandler(trapTypeUpgradeReport, upgradeReport)
	xtrap.RegisterHandler(trapTypeRestoreCfgReport, restoreCfgReport)
	xtrap.RegisterHandler(trapTypeSaveCfgReport, saveCfgReport)
	xtrap.RegisterHandler(trapTypeDeviceRebootReport, deviceRebootReport)
}
