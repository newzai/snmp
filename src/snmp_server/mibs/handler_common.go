package mibs

//uslRegisterAgain usl_register_again Octstring
func uslRegisterAgain(value interface{}) error {
	//“register”
	return nil
}

//uslRegisterTtl usl_register_ttl int
func uslRegisterTtl(value interface{}) error {
	//5-600
	return nil
}

//uslRegisterResult usl_register_result int
func uslRegisterResult(value interface{}) error {
	//0-注册失败 1-注册成功
	return nil
}

//uslRebootDevice usl_reboot_device Octstring
func uslRebootDevice(value interface{}) error {
	//“reboot”
	return nil
}

//uslSetDefault usl_set_default Octstring
func uslSetDefault(value interface{}) error {
	//“default”
	return nil
}

//uslFtpServerIp usl_ftp_server_ip Octstring
func uslFtpServerIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//uslFtpServerPort usl_ftp_server_port int
func uslFtpServerPort(value interface{}) error {
	//0-65535
	return nil
}

//uslFtpUserName usl_ftp_user_name Octstring
func uslFtpUserName(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//uslFtpUserPasswd usl_ftp_user_passwd Octstring
func uslFtpUserPasswd(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//uslFtpFileSize usl_ftp_file_size Octstring
func uslFtpFileSize(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//uslFtpSoftFileName usl_ftp_soft_file_name Octstring
func uslFtpSoftFileName(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//uslFtpSaveCfgFileName usl_ftp_save_cfg_file_name Octstring
func uslFtpSaveCfgFileName(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//uslFtpRestoreCfgFileName usl_ftp_restore_cfg_file_name Octstring
func uslFtpRestoreCfgFileName(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//uslKeepliveResult usl_keeplive_result int
func uslKeepliveResult(value interface{}) error {
	//0-保活失败 1-保活成功
	return nil
}
