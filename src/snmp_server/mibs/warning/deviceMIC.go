package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DeviceMICWarning  usl_device_mic_warning
type DeviceMICWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_mic_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *DeviceMICWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceMICWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceMICWarning) WarningType() string {
	return "mic"
}

//WarningStatus get warning type
func (r *DeviceMICWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceMICWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceMICWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *DeviceMICWarning) GetDemo() string {
	return r.Demo
}
