package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//RecordingFailWarning  usl_device_shell_warning
type RecordingFailWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_record_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *RecordingFailWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *RecordingFailWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *RecordingFailWarning) WarningType() string {
	return "recording_fail"
}

//WarningStatus get warning type
func (r *RecordingFailWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *RecordingFailWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *RecordingFailWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *RecordingFailWarning) GetDemo() string {
	return r.Demo
}
