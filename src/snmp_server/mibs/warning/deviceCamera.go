package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DeviceCameraWarning  usl_device_camera_warning
type DeviceCameraWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_camera_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *DeviceCameraWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceCameraWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceCameraWarning) WarningType() string {
	return "camera"
}

//WarningStatus get warning type
func (r *DeviceCameraWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceCameraWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceCameraWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *DeviceCameraWarning) GetDemo() string {
	return r.Demo
}
