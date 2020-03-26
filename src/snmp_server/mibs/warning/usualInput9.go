package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//UsualInput9Warning  usl_device_shell_warning
type UsualInput9Warning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_usual_input_nine_warning"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *UsualInput9Warning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *UsualInput9Warning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *UsualInput9Warning) WarningType() string {
	return "usual_input_9"
}

//WarningStatus get warning type
func (r *UsualInput9Warning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *UsualInput9Warning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *UsualInput9Warning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *UsualInput9Warning) GetDemo() string {
	return r.Demo
}
