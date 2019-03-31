package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//BackupPowerDropWarning  usl_device_shell_warning
type BackupPowerDropWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_backup_power_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *BackupPowerDropWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *BackupPowerDropWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *BackupPowerDropWarning) WarningType() string {
	return "backup_power_drop"
}

//WarningStatus get warning type
func (r *BackupPowerDropWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *BackupPowerDropWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *BackupPowerDropWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *BackupPowerDropWarning) GetDemo() string {
	return r.Demo
}
