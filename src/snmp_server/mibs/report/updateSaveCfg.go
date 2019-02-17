package report

import (
	"encoding/json"
	"snmp_server/mibs"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//UpgradeReport UpgradeReport
type UpgradeReport struct {
	NTID   string `json:"usl_ntid"`
	Status int    `json:"usl_upgrade_status"`
	Deep   int    `json:"usl_upgrade_deep"`
	Result string `json:"usl_ftp_result_string"`
}

func (r *UpgradeReport) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *UpgradeReport) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

type UpgradeRestoreCfgReport struct {
	NTID   string `json:"usl_ntid"`
	Status int    `json:"usl_restore_cfg_status"`
	Deep   int    `json:"usl_restore_cfg_deep"`
	Result string `json:"usl_ftp_result_string"`
}

func (r *UpgradeRestoreCfgReport) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *UpgradeRestoreCfgReport) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//UpgradeSaveCfgReport 配置保存报告
type UpgradeSaveCfgReport struct {
	NTID     string `json:"usl_ntid"`
	Status   int    `json:"usl_save_cfg_status"`
	Deep     int    `json:"usl_save_cfg_deep"`
	FileName string `json:"usl_save_cfg_file_name"`
}

func (r *UpgradeSaveCfgReport) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *UpgradeSaveCfgReport) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := mibs.PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}
