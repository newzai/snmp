package mibs

import (
	"encoding/json"

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
	value := PDU2JSON(packet.Variables)
	for key, value := range value {
		seelog.Infof("kv %s=%v", key, value)
	}
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//Keepalive Keepalvie
type Keepalive struct {
	DeviceIP   string `json:"usl_device_ip"`
	DevicePort int    `json:"usl_device_port"`
	NTID       string `json:"usl_ntid"`
}

func (r *Keepalive) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *Keepalive) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

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
	value := PDU2JSON(packet.Variables)
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
	value := PDU2JSON(packet.Variables)
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
	value := PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//DeviceRebootReport 设备重启
type DeviceRebootReport struct {
	NTID string `json:"usl_ntid"`
}

func (r *DeviceRebootReport) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceRebootReport) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//DeviceMICWarning  usl_device_mic_warning
type DeviceMICWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_mic_status"`
}

func (r *DeviceMICWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceMICWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceMICWarning) WarningType() string {
	return "mic"
}

//WarningStatus get warning type
func (r *DeviceMICWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceMICWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceMICWarning) GetNTID() string {
	return r.NTID
}

//DeviceSpeakerWarning  usl_device_speaker_warning
type DeviceSpeakerWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_speaker_status"`
}

func (r *DeviceSpeakerWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceSpeakerWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
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

//DeviceLEDWarning  usl_device_led_warning
type DeviceLEDWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_led_status"`
}

func (r *DeviceLEDWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceLEDWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceLEDWarning) WarningType() string {
	return "led"
}

//WarningStatus get warning type
func (r *DeviceLEDWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceLEDWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceLEDWarning) GetNTID() string {
	return r.NTID
}

//DeviceKeypWarning  usl_device_keyp_warning
type DeviceKeypWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_keyp_status"`
}

func (r *DeviceKeypWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceKeypWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceKeypWarning) WarningType() string {
	return "keyp"
}

//WarningStatus get warning type
func (r *DeviceKeypWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceKeypWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceKeypWarning) GetNTID() string {
	return r.NTID
}

//DeviceQuickWarning  usl_device_quick_warning
type DeviceQuickWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_quick_status"`
}

func (r *DeviceQuickWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceQuickWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceQuickWarning) WarningType() string {
	return "quick"
}

//WarningStatus get warning type
func (r *DeviceQuickWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceQuickWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceQuickWarning) GetNTID() string {
	return r.NTID
}

//DeviceCameraWarning  usl_device_camera_warning
type DeviceCameraWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_camera_status"`
}

func (r *DeviceCameraWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceCameraWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
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

//DeviceLCDWarning  usl_device_lcd_warning
type DeviceLCDWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_lcd_status"`
}

func (r *DeviceLCDWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceLCDWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
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

//DeviceShellWarning  usl_device_shell_warning
type DeviceShellWarning struct {
	NTID   string `json:"usl_ntid"`
	Clear  int    `json:"usl_device_warning_clear"`
	Status int    `json:"usl_device_shell_status"`
}

func (r *DeviceShellWarning) String() string {
	jdata, _ := json.MarshalIndent(r, "", " ")
	return string(jdata)
}

//FromSnmpPackage FromSnmpPackage
func (r *DeviceShellWarning) FromSnmpPackage(packet *gosnmp.SnmpPacket) {
	value := PDU2JSON(packet.Variables)
	jdata, _ := json.MarshalIndent(value, "", " ")
	seelog.Info("PDU json:", string(jdata))
	json.Unmarshal(jdata, r)

}

//WarningType get warning type
func (r *DeviceShellWarning) WarningType() string {
	return "shell"
}

//WarningStatus get warning type
func (r *DeviceShellWarning) WarningStatus() int {
	return r.Status
}

//IsClear is clear message
func (r *DeviceShellWarning) IsClear() bool {
	return r.Clear == 1
}

//GetNTID is clear message
func (r *DeviceShellWarning) GetNTID() string {
	return r.NTID
}

//IWarning warning interface
type IWarning interface {
	WarningType() string
	WarningStatus() int
	IsClear() bool
	GetNTID() string
}
