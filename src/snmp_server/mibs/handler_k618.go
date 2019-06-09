package mibs

//k618NetMode k618_net_mode int
func k618NetMode(value interface{}) error {
	//0-静态 1-dhcp 2-pppoe
	return nil
}

//k618NetStaticIp k618_net_static_ip Octstring
func k618NetStaticIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618NetStaticGateway k618_net_static_gateway Octstring
func k618NetStaticGateway(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618NetStaticMask k618_net_static_mask Octstring
func k618NetStaticMask(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618PppoeAccount k618_pppoe_account Octstring
func k618PppoeAccount(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618PppoePassword k618_pppoe_password Octstring
func k618PppoePassword(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618DnsMode k618_dns_mode int
func k618DnsMode(value interface{}) error {
	//0-关闭 1-开启
	return nil
}

//k618MasterDns k618_master_dns Octstring
func k618MasterDns(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618SlaveDns k618_slave_dns Octstring
func k618SlaveDns(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618Net0Adptive k618_net0_adptive int
func k618Net0Adptive(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618Net0Rate k618_net0_rate int
func k618Net0Rate(value interface{}) error {
	//0-10M 2-100M
	return nil
}

//k618Net0Mode k618_net0_mode int
func k618Net0Mode(value interface{}) error {
	//0-半工 1-双工
	return nil
}

//k618Net1Adptive k618_net1_adptive int
func k618Net1Adptive(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618Net1Rate k618_net1_rate int
func k618Net1Rate(value interface{}) error {
	//0-10M 2-100M
	return nil
}

//k618Net1Mode k618_net1_mode int
func k618Net1Mode(value interface{}) error {
	//0-半工 1-双工
	return nil
}

//k618AdminName k618_admin_name Octstring
func k618AdminName(value interface{}) error {
	//最大字符串长度：32
	return nil
}

//k618AdminPassword k618_admin_password Octstring
func k618AdminPassword(value interface{}) error {
	//最大字符串长度：32
	return nil
}

//k618LanguageMode k618_language_mode int
func k618LanguageMode(value interface{}) error {
	//0-中文 1-英文 2-韩文
	return nil
}

//k618HttpMode k618_http_mode int
func k618HttpMode(value interface{}) error {
	//0-http 1-https
	return nil
}

//k618HttpPort k618_http_port int
func k618HttpPort(value interface{}) error {
	//0-65535
	return nil
}

//k618SipLocalPort k618_sip_local_port int
func k618SipLocalPort(value interface{}) error {
	//0-65535
	return nil
}

//k618LogOuputMode k618_log_ouput_mode int
func k618LogOuputMode(value interface{}) error {
	//0-串口 1-telnet 2-网页
	return nil
}

//k618NtpMode k618_ntp_mode int
func k618NtpMode(value interface{}) error {
	//0-关闭 1-开启
	return nil
}

//k618PrimaryNtpAddress k618_primary_ntp_address Octstring
func k618PrimaryNtpAddress(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618PrimaryNtpPort k618_primary_ntp_port int
func k618PrimaryNtpPort(value interface{}) error {
	//0-65535
	return nil
}

//k618SecondNtpAddress k618_second_ntp_address Octstring
func k618SecondNtpAddress(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618SecondNtpPort k618_second_ntp_port int
func k618SecondNtpPort(value interface{}) error {
	//0-65535
	return nil
}

//k618TimeZone k618_time_zone int
func k618TimeZone(value interface{}) error {
	//0-25
	return nil
}

//k618SummerTimeMode k618_summer_time_mode int
func k618SummerTimeMode(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618SnmpEnable k618_snmp_enable int
func k618SnmpEnable(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618SnmpServerIp k618_snmp_server_ip Octstring
func k618SnmpServerIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618SnmpServerPort k618_snmp_server_port int
func k618SnmpServerPort(value interface{}) error {
	//0-65535
	return nil
}

//k618SnmpDeviceZone k618_snmp_device_zone Octstring
func k618SnmpDeviceZone(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SnmpDeviceName k618_snmp_device_name Octstring
func k618SnmpDeviceName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069WorkMode k618_tr069_work_mode int
func k618Tr069WorkMode(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618Tr069AcsServer k618_tr069_acs_server Octstring
func k618Tr069AcsServer(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069AcsAuthmode k618_tr069_acs_authmode int
func k618Tr069AcsAuthmode(value interface{}) error {
	//
	return nil
}

//k618Tr069AcsAccount k618_tr069_acs_account Octstring
func k618Tr069AcsAccount(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069AcsPassword k618_tr069_acs_password Octstring
func k618Tr069AcsPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069ConnectPeriod k618_tr069_connect_period int
func k618Tr069ConnectPeriod(value interface{}) error {
	//
	return nil
}

//k618Tr069CpeAuthmode k618_tr069_cpe_authmode int
func k618Tr069CpeAuthmode(value interface{}) error {
	//
	return nil
}

//k618Tr069CpeReqName k618_tr069_cpe_req_name Octstring
func k618Tr069CpeReqName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069CpeReqPassword k618_tr069_cpe_req_password Octstring
func k618Tr069CpeReqPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069InformTimeout k618_tr069_inform_timeout int
func k618Tr069InformTimeout(value interface{}) error {
	//
	return nil
}

//k618Tr069Manufacturer k618_tr069_manufacturer Octstring
func k618Tr069Manufacturer(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069Oui k618_tr069_oui Octstring
func k618Tr069Oui(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069ProductClass k618_tr069_product_class Octstring
func k618Tr069ProductClass(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069SerialNamber k618_tr069_serial_namber Octstring
func k618Tr069SerialNamber(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069HardwareVersion k618_tr069_hardware_version Octstring
func k618Tr069HardwareVersion(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069SoftwareVersion k618_tr069_software_version Octstring
func k618Tr069SoftwareVersion(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069ConreqPort k618_tr069_conreq_port int
func k618Tr069ConreqPort(value interface{}) error {
	//
	return nil
}

//k618Tr069ConregUri k618_tr069_conreg_uri Octstring
func k618Tr069ConregUri(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069ConreqUrl k618_tr069_conreq_url Octstring
func k618Tr069ConreqUrl(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069Dn0Number k618_tr069_dn0_number Octstring
func k618Tr069Dn0Number(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069Dn0Name k618_tr069_dn0_name Octstring
func k618Tr069Dn0Name(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618Tr069AutoProvision k618_tr069_auto_provision int
func k618Tr069AutoProvision(value interface{}) error {
	//
	return nil
}

//k618ControlOut1PutType k618_control_out1_put_type int
func k618ControlOut1PutType(value interface{}) error {
	//0-禁用 1-干接点 2-门禁
	return nil
}

//k618ControlOut1ConnectLev k618_control_out1_connect_lev int
func k618ControlOut1ConnectLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k618ControlOut1ConnectSec k618_control_out1_connect_sec int
func k618ControlOut1ConnectSec(value interface{}) error {
	//秒数
	return nil
}

//k618ControlOut1TriggerType k618_control_out1_trigger_type int
func k618ControlOut1TriggerType(value interface{}) error {
	//0-振铃 1-摘机 2-通话 3-振铃结束 4-通话结束
	return nil
}

//k618ControlOut1DoorLev k618_control_out1_door_lev int
func k618ControlOut1DoorLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k618ControlOut1DoorSec k618_control_out1_door_sec int
func k618ControlOut1DoorSec(value interface{}) error {
	//秒数
	return nil
}

//k618ControlOut1DoorNum k618_control_out1_door_num Octstring
func k618ControlOut1DoorNum(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k618ControlOut1DoorPsw k618_control_out1_door_psw Octstring
func k618ControlOut1DoorPsw(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k618ControlOut1DoorMaxTalkLen k618_control_out1_door_max_talk_len int
func k618ControlOut1DoorMaxTalkLen(value interface{}) error {
	//秒数
	return nil
}

//k618ControlOut2PutType k618_control_out2_put_type int
func k618ControlOut2PutType(value interface{}) error {
	//0-禁用 1-干接点 2-门禁
	return nil
}

//k618ControlOut2ConnectLev k618_control_out2_connect_lev int
func k618ControlOut2ConnectLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k618ControlOut2ConnectSec k618_control_out2_connect_sec int
func k618ControlOut2ConnectSec(value interface{}) error {
	//秒数
	return nil
}

//k618ControlOut2TriggerType k618_control_out2_trigger_type int
func k618ControlOut2TriggerType(value interface{}) error {
	//0-振铃 1-摘机 2-通话 3-振铃结束 4-通话结束
	return nil
}

//k618ControlOut2DoorLev k618_control_out2_door_lev int
func k618ControlOut2DoorLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k618ControlOut2DoorSec k618_control_out2_door_sec int
func k618ControlOut2DoorSec(value interface{}) error {
	//秒数
	return nil
}

//k618ControlOut2DoorNum k618_control_out2_door_num Octstring
func k618ControlOut2DoorNum(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k618ControlOut2DoorPsw k618_control_out2_door_psw Octstring
func k618ControlOut2DoorPsw(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k618ControlOut2DoorMaxTalkLen k618_control_out2_door_max_talk_len int
func k618ControlOut2DoorMaxTalkLen(value interface{}) error {
	//秒数
	return nil
}

//k618PnmSwitch k618_pnm_switch int
func k618PnmSwitch(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618PnmServerIp k618_pnm_server_ip Octstring
func k618PnmServerIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618PnmServerPort k618_pnm_server_port int
func k618PnmServerPort(value interface{}) error {
	//0-65535
	return nil
}

//k618PnmAreaName k618_pnm_area_name Octstring
func k618PnmAreaName(value interface{}) error {
	//最大字符串长度：20
	return nil
}

//k618PnmDeviceName k618_pnm_device_name Octstring
func k618PnmDeviceName(value interface{}) error {
	//最大字符串长度：20
	return nil
}

//k618PnmSelfDetectTimer k618_pnm_self_detect_timer int
func k618PnmSelfDetectTimer(value interface{}) error {
	//秒数
	return nil
}

//k618VideoMode k618_video_mode int
func k618VideoMode(value interface{}) error {
	//0-720p 1-480p
	return nil
}

//k618CodecTypeList k618_codec_type_list Octstring
func k618CodecTypeList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618HotlineEnable k618_hotline_enable int
func k618HotlineEnable(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618HotlineNumber k618_hotline_number Octstring
func k618HotlineNumber(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618HotlineAccount k618_hotline_account int
func k618HotlineAccount(value interface{}) error {
	//0-5
	return nil
}

//k618QuickAttrib k618_quick_attrib Octstring
func k618QuickAttrib(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618QuickNumberList k618_quick_number_list Octstring
func k618QuickNumberList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k618QuickAutoList k618_quick_auto_list Octstring
func k618QuickAutoList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618QuickAccountList k618_quick_account_list Octstring
func k618QuickAccountList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618BoardcastIpList k618_boardcast_ip_list Octstring
func k618BoardcastIpList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k618BoardcastPortList k618_boardcast_port_list Octstring
func k618BoardcastPortList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618RecvBoardcastVoicePriority k618_recv_boardcast_voice_priority int
func k618RecvBoardcastVoicePriority(value interface{}) error {
	//0-11
	return nil
}

//k618RecvBoardcastEnableList k618_recv_boardcast_enable_list Octstring
func k618RecvBoardcastEnableList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618RecvBoardcastPriorityList k618_recv_boardcast_priority_list Octstring
func k618RecvBoardcastPriorityList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618RecvBoardcastIpList k618_recv_boardcast_ip_list Octstring
func k618RecvBoardcastIpList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k618RecvBoardcastPortList k618_recv_boardcast_port_list Octstring
func k618RecvBoardcastPortList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618MicrophoneVolume k618_microphone_volume int
func k618MicrophoneVolume(value interface{}) error {
	//1-9
	return nil
}

//k618SpeakerVolume k618_speaker_volume int
func k618SpeakerVolume(value interface{}) error {
	//0-9
	return nil
}

//k618HookonWaitTime k618_hookon_wait_time int
func k618HookonWaitTime(value interface{}) error {
	//1-30
	return nil
}

//k618RingStyle k618_ring_style int
func k618RingStyle(value interface{}) error {
	//1-7
	return nil
}

//k618RingVolume k618_ring_volume int
func k618RingVolume(value interface{}) error {
	//0-9
	return nil
}

//k618RecordFtpEnable k618_record_ftp_enable int
func k618RecordFtpEnable(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618UploadFtpServer k618_upload_ftp_server Octstring
func k618UploadFtpServer(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618UploadFtpPort k618_upload_ftp_port int
func k618UploadFtpPort(value interface{}) error {
	//0-65535
	return nil
}

//k618UploadFtpUser k618_upload_ftp_user Octstring
func k618UploadFtpUser(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618UploadFtpPassword k618_upload_ftp_password Octstring
func k618UploadFtpPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618UploadFtpFilePath k618_upload_ftp_file_path Octstring
func k618UploadFtpFilePath(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618UploadFtpFilePrefix k618_upload_ftp_file_prefix Octstring
func k618UploadFtpFilePrefix(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618NoiseFilterEnable k618_noise_filter_enable int
func k618NoiseFilterEnable(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k618SipAccountEnable k618_sip_account_enable int
func k618SipAccountEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k618SipRegisterSwitch k618_sip_register_switch int
func k618SipRegisterSwitch(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k618SipUnregisterSwitch k618_sip_unregister_switch int
func k618SipUnregisterSwitch(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k618SipRegisterExpire k618_sip_register_expire int
func k618SipRegisterExpire(value interface{}) error {
	//秒数
	return nil
}

//k618SipDisplayName k618_sip_display_name Octstring
func k618SipDisplayName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SipAccountName k618_sip_account_name Octstring
func k618SipAccountName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SipAuthName k618_sip_auth_name Octstring
func k618SipAuthName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SipAccountPassword k618_sip_account_password Octstring
func k618SipAccountPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SipRegServer0 k618_sip_reg_server0 Octstring
func k618SipRegServer0(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618SipRegPort0 k618_sip_reg_port0 int
func k618SipRegPort0(value interface{}) error {
	//0-65535
	return nil
}

//k618SipRegDomain0 k618_sip_reg_domain0 Octstring
func k618SipRegDomain0(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SipRegServer1 k618_sip_reg_server1 Octstring
func k618SipRegServer1(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618SipRegPort1 k618_sip_reg_port1 int
func k618SipRegPort1(value interface{}) error {
	//0-65535
	return nil
}

//k618SipRegDomain1 k618_sip_reg_domain1 Octstring
func k618SipRegDomain1(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SipRegServer2 k618_sip_reg_server2 Octstring
func k618SipRegServer2(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k618SipRegPort2 k618_sip_reg_port2 int
func k618SipRegPort2(value interface{}) error {
	//0-65535
	return nil
}

//k618SipRegDomain2 k618_sip_reg_domain2 Octstring
func k618SipRegDomain2(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SipUserAgent k618_sip_user_agent Octstring
func k618SipUserAgent(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k618SipHeartBeatEnable k618_sip_heart_beat_enable int
func k618SipHeartBeatEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k618SipHeartBeatTime k618_sip_heart_beat_time int
func k618SipHeartBeatTime(value interface{}) error {
	//秒数
	return nil
}

//k618SipAutoAnswerEnable k618_sip_auto_answer_enable int
func k618SipAutoAnswerEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k618SipAutoAnswerTime k618_sip_auto_answer_time int
func k618SipAutoAnswerTime(value interface{}) error {
	//秒数
	return nil
}

//k618DigitRule15 k618_digit_rule_1_5 Octstring
func k618DigitRule15(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule610 k618_digit_rule_6_10 Octstring
func k618DigitRule610(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule1115 k618_digit_rule_11_15 Octstring
func k618DigitRule1115(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule1620 k618_digit_rule_16_20 Octstring
func k618DigitRule1620(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule2125 k618_digit_rule_21_25 Octstring
func k618DigitRule2125(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule2630 k618_digit_rule_26_30 Octstring
func k618DigitRule2630(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule3135 k618_digit_rule_31_35 Octstring
func k618DigitRule3135(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule3640 k618_digit_rule_36_40 Octstring
func k618DigitRule3640(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule4145 k618_digit_rule_41_45 Octstring
func k618DigitRule4145(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k618DigitRule4650 k618_digit_rule_46_50 Octstring
func k618DigitRule4650(value interface{}) error {
	//最大字符串长度：1200
	return nil
}
