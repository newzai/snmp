package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//NtpDropWarning  usl_device_shell_warning
type NtpDropWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_ntp_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *NtpDropWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *NtpDropWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *NtpDropWarning) WarningType() string {
	return "ntp_drop"
}

//WarningStatus get warning type
func (r *NtpDropWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *NtpDropWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *NtpDropWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *NtpDropWarning) GetDemo() string {
	return r.Demo
}
