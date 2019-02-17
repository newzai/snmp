package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DeviceSpeakerWarning  usl_device_speaker_warning
type DeviceSpeakerWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_speaker_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *DeviceSpeakerWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceSpeakerWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceSpeakerWarning) WarningType() string {
	return "speaker"
}

//WarningStatus get warning type
func (r *DeviceSpeakerWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceSpeakerWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceSpeakerWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *DeviceSpeakerWarning) GetDemo() string {
	return r.Demo
}
