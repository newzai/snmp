package warning

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//MianBackupSwitchWarning  usl_device_shell_warning
type MianBackupSwitchWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_main_backup_switch_status"`
	Demo   string `json:"usl_device_warning_demo,omitempty"`
}

func (r *MianBackupSwitchWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *MianBackupSwitchWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *MianBackupSwitchWarning) WarningType() string {
	return "main_backup_switch"
}

//WarningStatus get warning type
func (r *MianBackupSwitchWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *MianBackupSwitchWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *MianBackupSwitchWarning) GetNTID() string {
	return r.NTID
}

//GetDemo get demo
func (r *MianBackupSwitchWarning) GetDemo() string {
	return r.Demo
}
