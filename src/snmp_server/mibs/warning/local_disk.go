package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DiskWarning  usl_device_shell_warning
type DiskWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_power_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *DiskWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DiskWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DiskWarning) WarningType() string {
	return "disk"
}

//WarningStatus get warning type
func (r *DiskWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DiskWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DiskWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *DiskWarning) GetDemo() string {
	return r.Demo
}
