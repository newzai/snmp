package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//UsualInput3Warning  usl_device_shell_warning
type UsualInput3Warning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_usual_input_three_warning"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *UsualInput3Warning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *UsualInput3Warning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *UsualInput3Warning) WarningType() string {
	return "usual_input_3"
}

//WarningStatus get warning type
func (r *UsualInput3Warning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *UsualInput3Warning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *UsualInput3Warning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *UsualInput3Warning) GetDemo() string {
	return r.Demo
}
