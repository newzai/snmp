package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DeviceLCDWarning  usl_device_lcd_warning
type DeviceLCDWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_lcd_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *DeviceLCDWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceLCDWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceLCDWarning) WarningType() string {
	return "lcd"
}

//WarningStatus get warning type
func (r *DeviceLCDWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceLCDWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceLCDWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *DeviceLCDWarning) GetDemo() string {
	return r.Demo
}
