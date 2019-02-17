package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DeviceQuickWarning  usl_device_quick_warning
type DeviceQuickWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_quick_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *DeviceQuickWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceQuickWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceQuickWarning) WarningType() string {
	return "quick"
}

//WarningStatus get warning type
func (r *DeviceQuickWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceQuickWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceQuickWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *DeviceQuickWarning) GetDemo() string {
	return r.Demo
}
