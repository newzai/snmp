package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//NtpWarning  usl_device_shell_warning
type NtpWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_power_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *NtpWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *NtpWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *NtpWarning) WarningType() string {
	return "ntp"
}

//WarningStatus get warning type
func (r *NtpWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *NtpWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *NtpWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *NtpWarning) GetDemo() string {
	return r.Demo
}
