package xtraphandler

import (
	"net"
	"snmp_server/mibs"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xtrap"
	"snmp_server/xwarning"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DoWarningTest DoWarningTest
func DoWarningTest(msg mibs.IWarning) {
	doWarning(msg)
}

func doWarning(msg mibs.IWarning) {
	if msg.IsClear() {
		xwarning.ClearWarning(msg.GetNTID(), msg.WarningType(), xdb.EngineWarning)
		return
	}

	w, has := xwarning.GetWarning(msg.GetNTID(), msg.WarningType(), xdb.EngineWarning)
	if !has {
		t, err := model.GetTerminalByNTID(msg.GetNTID(), xdb.Engine)
		if err != nil {
			seelog.Errorf("GetTerminalByNTID err:%v", err)
			return
		}
		if t == nil {
			seelog.Errorf("GetTerminalByNTID nil")
			return
		}

		w = &xwarning.Warning{
			TID:    t.ID,
			TName:  t.Name,
			NTID:   t.NTID,
			Path:   t.Path,
			WType:  msg.WarningType(),
			WValue: msg.WarningStatus(),
		}
		err = xwarning.InsertWarning(w, xdb.EngineWarning)
		if err != nil {
			seelog.Errorf("insert %s %s warning error:%v", w.NTID, w.WType, err)
		} else {
			seelog.Infof("insert %s %s warning ok", w.NTID, w.WType)
		}
	}
}

func deviceMICWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {

	msg := new(mibs.DeviceMICWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)

}

func deviceSpeakerWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceSpeakerWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func deviceLEDWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {

	msg := new(mibs.DeviceLEDWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

func deviceKeypWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceKeypWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

func deviceQuickWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceQuickWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func deviceCameraWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceCameraWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func deviceLCDWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceLCDWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func deviceShellWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceShellWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func init() {

	xtrap.RegisterHandler(trapTypeDeviceMICWarning, deviceMICWarning)
	xtrap.RegisterHandler(trapTypeDeviceSpeakerWaring, deviceSpeakerWarning)
	xtrap.RegisterHandler(trapTypeDeviceLEDWarning, deviceLEDWarning)
	xtrap.RegisterHandler(trapTypeDeviceKeypWarning, deviceKeypWarning)
	xtrap.RegisterHandler(trapTypeDeviceQuickWarning, deviceQuickWarning)
	xtrap.RegisterHandler(trapTypeDeviceCameraWarning, deviceCameraWarning)
	xtrap.RegisterHandler(trapTypeDeviceLCDWarning, deviceLCDWarning)
	xtrap.RegisterHandler(trapTypeDeviceShellWarning, deviceShellWarning)
}
