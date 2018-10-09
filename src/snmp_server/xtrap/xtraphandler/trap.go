package xtraphandler

import (
	"net"
	"snmp_server/mibs"
	"snmp_server/xtrap"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

func deviceMICWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {

	msg := new(mibs.DeviceMICWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
}

func deviceSpeakerWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceSpeakerWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
}
func deviceLEDWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {

	msg := new(mibs.DeviceLEDWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
}

func deviceKeypWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceKeypWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
}

func deviceQuickWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceQuickWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
}
func deviceCameraWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceCameraWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
}
func deviceLCDWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceLCDWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
}
func deviceShellWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(mibs.DeviceShellWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv upgradeReport:", msg)
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
