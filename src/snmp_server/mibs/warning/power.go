package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//PowerWarning  usl_device_shell_warning
type PowerWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_power_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *PowerWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *PowerWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *PowerWarning) WarningType() string {
	return "power"
}

//WarningStatus get warning type
func (r *PowerWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *PowerWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *PowerWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *PowerWarning) GetDemo() string {
	return r.Demo
}
