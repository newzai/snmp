package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DeviceShellWarning  usl_device_shell_warning
type DeviceShellWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_shell_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *DeviceShellWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceShellWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceShellWarning) WarningType() string {
	return "shell"
}

//WarningStatus get warning type
func (r *DeviceShellWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceShellWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceShellWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *DeviceShellWarning) GetDemo() string {
	return r.Demo
}
