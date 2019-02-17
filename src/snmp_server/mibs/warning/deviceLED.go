package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DeviceLEDWarning  usl_device_led_warning
type DeviceLEDWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_led_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *DeviceLEDWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceLEDWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceLEDWarning) WarningType() string {
	return "led"
}

//WarningStatus get warning type
func (r *DeviceLEDWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceLEDWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceLEDWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *DeviceLEDWarning) GetDemo() string {
	return r.Demo
}
