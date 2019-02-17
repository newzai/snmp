package report

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//Register Register Message
type Register struct {
	ProcductType    string `json:"usl_procduct_type"`
	DeviceIP        string `json:"usl_device_ip"`
	DevicePort      int    `json:"usl_device_port"`
	NTID            string `json:"usl_ntid"`
	ProcductName    string `json:"usl_procduct_name"`
	Zone            string `json:"usl_zone"`
	ProcductVersion string `json:"usl_procduct_version"`
}

func (r *Register) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage set data from Snmp package
func (r *Register) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	for key, value := range value {
		seelog.Infof("kv %s=%v", key, value)
	}

	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}
