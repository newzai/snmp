package mibs

import (
	"github.com/soniah/gosnmp"
)

//1.3.6.1.4.1
var oidCommons = []OIDAttr{
	OIDAttr{Name: "usl_register", OID: "1800.50.1.1", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_keeplive", OID: "1800.50.1.2", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_upgrade_report", OID: "1800.50.1.3", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_restore_cfg_report", OID: "1800.50.1.4", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_save_cfg_report", OID: "1800.50.1.5", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_reboot_report", OID: "1800.50.1.6", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_mic_warning", OID: "1800.50.1.7", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_speaker_warning", OID: "1800.50.1.8", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_led_warning", OID: "1800.50.1.9", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_keyp_warning", OID: "1800.50.1.10", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_quick_warning", OID: "1800.50.1.11", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_camera_warning", OID: "1800.50.1.12", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_lcd_warning", OID: "1800.50.1.13", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_shell_warning", OID: "1800.50.1.14", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_power_warning", OID: "1800.50.1.15", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_callout_warning", OID: "1800.50.1.16", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_register_warning", OID: "1800.50.1.17", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_drycontact_warning", OID: "1800.50.1.18", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_main_backup_switch_warning", OID: "1800.50.1.19", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_ntp_drop_warning", OID: "1800.50.1.20", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_recording_fail_warning", OID: "1800.50.1.21", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_main_power_drop_warning", OID: "1800.50.1.22", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_backup_power_drop_warning", OID: "1800.50.1.23", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_main_net_drop_warning", OID: "1800.50.1.24", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_backup_net_drop_warning", OID: "1800.50.1.25", Type: gosnmp.ObjectIdentifier, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_procduct_type", OID: "1800.50.2.1", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_ip", OID: "1800.50.2.2", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_port", OID: "1800.50.2.3", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_ntid", OID: "1800.50.2.4", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_procduct_name", OID: "1800.50.2.5", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_zone", OID: "1800.50.2.6", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_procduct_version", OID: "1800.50.2.7", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_upgrade_status", OID: "1800.50.2.8", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_upgrade_deep", OID: "1800.50.2.9", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_restore_cfg_status", OID: "1800.50.2.10", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_restore_cfg_deep", OID: "1800.50.2.11", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_save_cfg_status", OID: "1800.50.2.12", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_save_cfg_deep", OID: "1800.50.2.13", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_save_cfg_file_name", OID: "1800.50.2.14", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_ftp_result_string", OID: "1800.50.2.15", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_warning_clear", OID: "1800.50.2.16", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_mic_status", OID: "1800.50.2.17", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_speaker_status", OID: "1800.50.2.18", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_led_status", OID: "1800.50.2.19", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_keyp_status", OID: "1800.50.2.20", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_quick_status", OID: "1800.50.2.21", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_camera_status", OID: "1800.50.2.22", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_lcd_status", OID: "1800.50.2.23", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_shell_status", OID: "1800.50.2.24", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_power_status", OID: "1800.50.2.25", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_callout_status", OID: "1800.50.2.26", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_register_status", OID: "1800.50.2.27", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_drycontact_status", OID: "1800.50.2.28", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_warning_demo", OID: "1800.50.2.29", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_main_backup_switch_status", OID: "1800.50.2.30", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_ntp_status", OID: "1800.50.2.31", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_record_status", OID: "1800.50.2.32", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_main_power_status", OID: "1800.50.2.33", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_backup_power_status", OID: "1800.50.2.34", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_main_net_status", OID: "1800.50.2.35", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_device_backup_net_status", OID: "1800.50.2.36", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "usl_register_again", OID: "1800.50.3.1", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslRegisterAgain},
	OIDAttr{Name: "usl_register_ttl", OID: "1800.50.3.2", Type: gosnmp.Integer, ReadOnly: false, ValidHander: uslRegisterTtl},
	OIDAttr{Name: "usl_register_result", OID: "1800.50.3.3", Type: gosnmp.Integer, ReadOnly: false, ValidHander: uslRegisterResult},
	OIDAttr{Name: "usl_reboot_device", OID: "1800.50.3.4", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslRebootDevice},
	OIDAttr{Name: "usl_set_default", OID: "1800.50.3.5", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslSetDefault},
	OIDAttr{Name: "usl_ftp_server_ip", OID: "1800.50.3.6", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslFtpServerIp},
	OIDAttr{Name: "usl_ftp_server_port", OID: "1800.50.3.7", Type: gosnmp.Integer, ReadOnly: false, ValidHander: uslFtpServerPort},
	OIDAttr{Name: "usl_ftp_user_name", OID: "1800.50.3.8", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslFtpUserName},
	OIDAttr{Name: "usl_ftp_user_passwd", OID: "1800.50.3.9", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslFtpUserPasswd},
	OIDAttr{Name: "usl_ftp_file_size", OID: "1800.50.3.10", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslFtpFileSize},
	OIDAttr{Name: "usl_ftp_soft_file_name", OID: "1800.50.3.11", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslFtpSoftFileName},
	OIDAttr{Name: "usl_ftp_save_cfg_file_name", OID: "1800.50.3.12", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslFtpSaveCfgFileName},
	OIDAttr{Name: "usl_ftp_restore_cfg_file_name", OID: "1800.50.3.13", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: uslFtpRestoreCfgFileName},
	OIDAttr{Name: "usl_keeplive_result", OID: "1800.50.3.14", Type: gosnmp.Integer, ReadOnly: false, ValidHander: uslKeepliveResult},
}

func init() {

	for _, attr := range oidCommons {
		cAttr := attr
		cAttr.OID = oidPrefix + cAttr.OID

		jsonKeyAttr[cAttr.Name] = &cAttr
		oidKeyAttr[cAttr.OID] = &cAttr
	}
}
