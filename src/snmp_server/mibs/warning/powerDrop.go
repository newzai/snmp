package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//MainPowerDropWarning  usl_device_shell_warning
type MainPowerDropWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_main_power_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *MainPowerDropWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *MainPowerDropWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *MainPowerDropWarning) WarningType() string {
	return "main_power_drop"
}

//WarningStatus get warning type
func (r *MainPowerDropWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *MainPowerDropWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *MainPowerDropWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *MainPowerDropWarning) GetDemo() string {
	return r.Demo
}
