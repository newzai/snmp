package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//MainNetDropWarning  usl_device_shell_warning
type MainNetDropWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_main_net_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *MainNetDropWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *MainNetDropWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *MainNetDropWarning) WarningType() string {
	return "main_net_drop"
}

//WarningStatus get warning type
func (r *MainNetDropWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *MainNetDropWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *MainNetDropWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *MainNetDropWarning) GetDemo() string {
	return r.Demo
}
