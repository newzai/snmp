package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//BackupNetDropWarning  usl_device_shell_warning
type BackupNetDropWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_backup_net_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *BackupNetDropWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *BackupNetDropWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *BackupNetDropWarning) WarningType() string {
	return "backup_net_drop"
}

//WarningStatus get warning type
func (r *BackupNetDropWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *BackupNetDropWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *BackupNetDropWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *BackupNetDropWarning) GetDemo() string {
	return r.Demo
}
