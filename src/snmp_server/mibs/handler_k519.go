package mibs

//k519NetMode k519_net_mode int
func k519NetMode(value interface{}) error {
	//0-静态 1-dhcp 2-pppoe
	return nil
}

//k519NetStaticIp k519_net_static_ip Octstring
func k519NetStaticIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519NetStaticGateway k519_net_static_gateway Octstring
func k519NetStaticGateway(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519NetStaticMask k519_net_static_mask Octstring
func k519NetStaticMask(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519PppoeAccount k519_pppoe_account Octstring
func k519PppoeAccount(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519PppoePassword k519_pppoe_password Octstring
func k519PppoePassword(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519DnsMode k519_dns_mode int
func k519DnsMode(value interface{}) error {
	//0-关闭 1-开启
	return nil
}

//k519MasterDns k519_master_dns Octstring
func k519MasterDns(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519SlaveDns k519_slave_dns Octstring
func k519SlaveDns(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519Net0Adptive k519_net0_adptive int
func k519Net0Adptive(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k519Net0Rate k519_net0_rate int
func k519Net0Rate(value interface{}) error {
	//0-10M 2-100M
	return nil
}

//k519Net0Mode k519_net0_mode int
func k519Net0Mode(value interface{}) error {
	//0-半工 1-双工
	return nil
}

//k519Net1Adptive k519_net1_adptive int
func k519Net1Adptive(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k519Net1Rate k519_net1_rate int
func k519Net1Rate(value interface{}) error {
	//0-10M 2-100M
	return nil
}

//k519Net1Mode k519_net1_mode int
func k519Net1Mode(value interface{}) error {
	//0-半工 1-双工
	return nil
}

//k519AdminName k519_admin_name Octstring
func k519AdminName(value interface{}) error {
	//最大字符串长度：32
	return nil
}

//k519AdminPassword k519_admin_password Octstring
func k519AdminPassword(value interface{}) error {
	//最大字符串长度：32
	return nil
}

//k519LanguageMode k519_language_mode int
func k519LanguageMode(value interface{}) error {
	//0-中文 1-英文 2-韩文
	return nil
}

//k519HttpMode k519_http_mode int
func k519HttpMode(value interface{}) error {
	//0-http 1-https
	return nil
}

//k519HttpPort k519_http_port int
func k519HttpPort(value interface{}) error {
	//0-65535
	return nil
}

//k519TelnetPort k519_telnet_port int
func k519TelnetPort(value interface{}) error {
	//0-65535
	return nil
}

//k519SipLocalPort k519_sip_local_port int
func k519SipLocalPort(value interface{}) error {
	//0-65535
	return nil
}

//k519LogOuputMode k519_log_ouput_mode int
func k519LogOuputMode(value interface{}) error {
	//0-串口 1-telnet 2-网页
	return nil
}

//k519NtpMode k519_ntp_mode int
func k519NtpMode(value interface{}) error {
	//0-关闭 1-开启
	return nil
}

//k519PrimaryNtpAddress k519_primary_ntp_address Octstring
func k519PrimaryNtpAddress(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519PrimaryNtpPort k519_primary_ntp_port int
func k519PrimaryNtpPort(value interface{}) error {
	//0-65535
	return nil
}

//k519SecondNtpAddress k519_second_ntp_address Octstring
func k519SecondNtpAddress(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519SecondNtpPort k519_second_ntp_port int
func k519SecondNtpPort(value interface{}) error {
	//0-65535
	return nil
}

//k519TimeZone k519_time_zone int
func k519TimeZone(value interface{}) error {
	//0-25
	return nil
}

//k519SnmpEnable k519_snmp_enable int
func k519SnmpEnable(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k519SnmpServerIp k519_snmp_server_ip Octstring
func k519SnmpServerIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519SnmpServerPort k519_snmp_server_port int
func k519SnmpServerPort(value interface{}) error {
	//0-65535
	return nil
}

//k519SnmpDeviceZone k519_snmp_device_zone Octstring
func k519SnmpDeviceZone(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SnmpDeviceName k519_snmp_device_name Octstring
func k519SnmpDeviceName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069WorkMode k519_tr069_work_mode int
func k519Tr069WorkMode(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k519Tr069AcsServer k519_tr069_acs_server Octstring
func k519Tr069AcsServer(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069AcsAuthmode k519_tr069_acs_authmode int
func k519Tr069AcsAuthmode(value interface{}) error {
	//
	return nil
}

//k519Tr069AcsAccount k519_tr069_acs_account Octstring
func k519Tr069AcsAccount(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069AcsPassword k519_tr069_acs_password Octstring
func k519Tr069AcsPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069ConnectPeriod k519_tr069_connect_period int
func k519Tr069ConnectPeriod(value interface{}) error {
	//
	return nil
}

//k519Tr069CpeAuthmode k519_tr069_cpe_authmode int
func k519Tr069CpeAuthmode(value interface{}) error {
	//
	return nil
}

//k519Tr069CpeReqName k519_tr069_cpe_req_name Octstring
func k519Tr069CpeReqName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069CpeReqPassword k519_tr069_cpe_req_password Octstring
func k519Tr069CpeReqPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069InformTimeout k519_tr069_inform_timeout int
func k519Tr069InformTimeout(value interface{}) error {
	//
	return nil
}

//k519Tr069Manufacturer k519_tr069_manufacturer Octstring
func k519Tr069Manufacturer(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069Oui k519_tr069_oui Octstring
func k519Tr069Oui(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069ProductClass k519_tr069_product_class Octstring
func k519Tr069ProductClass(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069SerialNamber k519_tr069_serial_namber Octstring
func k519Tr069SerialNamber(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069HardwareVersion k519_tr069_hardware_version Octstring
func k519Tr069HardwareVersion(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069SoftwareVersion k519_tr069_software_version Octstring
func k519Tr069SoftwareVersion(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069ConreqPort k519_tr069_conreq_port int
func k519Tr069ConreqPort(value interface{}) error {
	//
	return nil
}

//k519Tr069ConregUri k519_tr069_conreg_uri Octstring
func k519Tr069ConregUri(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069ConreqUrl k519_tr069_conreq_url Octstring
func k519Tr069ConreqUrl(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069Dn0Number k519_tr069_dn0_number Octstring
func k519Tr069Dn0Number(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069Dn0Name k519_tr069_dn0_name Octstring
func k519Tr069Dn0Name(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519Tr069AutoProvision k519_tr069_auto_provision int
func k519Tr069AutoProvision(value interface{}) error {
	//
	return nil
}

//k519ControlOut1PutType k519_control_out1_put_type int
func k519ControlOut1PutType(value interface{}) error {
	//0-禁用 1-干接点 2-门禁
	return nil
}

//k519ControlOut1ConnectLev k519_control_out1_connect_lev int
func k519ControlOut1ConnectLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k519ControlOut1ConnectSec k519_control_out1_connect_sec int
func k519ControlOut1ConnectSec(value interface{}) error {
	//秒数
	return nil
}

//k519ControlOut1TriggerType k519_control_out1_trigger_type int
func k519ControlOut1TriggerType(value interface{}) error {
	//0-振铃 1-摘机 2-通话 3-振铃结束 4-通话结束
	return nil
}

//k519ControlOut1DoorLev k519_control_out1_door_lev int
func k519ControlOut1DoorLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k519ControlOut1DoorSec k519_control_out1_door_sec int
func k519ControlOut1DoorSec(value interface{}) error {
	//秒数
	return nil
}

//k519ControlOut1DoorNum k519_control_out1_door_num Octstring
func k519ControlOut1DoorNum(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k519ControlOut1DoorPsw k519_control_out1_door_psw Octstring
func k519ControlOut1DoorPsw(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k519ControlOut1DoorMaxTalkLen k519_control_out1_door_max_talk_len int
func k519ControlOut1DoorMaxTalkLen(value interface{}) error {
	//秒数
	return nil
}

//k519ControlOut2PutType k519_control_out2_put_type int
func k519ControlOut2PutType(value interface{}) error {
	//0-禁用 1-干接点 2-门禁
	return nil
}

//k519ControlOut2ConnectLev k519_control_out2_connect_lev int
func k519ControlOut2ConnectLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k519ControlOut2ConnectSec k519_control_out2_connect_sec int
func k519ControlOut2ConnectSec(value interface{}) error {
	//秒数
	return nil
}

//k519ControlOut2TriggerType k519_control_out2_trigger_type int
func k519ControlOut2TriggerType(value interface{}) error {
	//0-振铃 1-摘机 2-通话 3-振铃结束 4-通话结束
	return nil
}

//k519ControlOut2DoorLev k519_control_out2_door_lev int
func k519ControlOut2DoorLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k519ControlOut2DoorSec k519_control_out2_door_sec int
func k519ControlOut2DoorSec(value interface{}) error {
	//秒数
	return nil
}

//k519ControlOut2DoorNum k519_control_out2_door_num Octstring
func k519ControlOut2DoorNum(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k519ControlOut2DoorPsw k519_control_out2_door_psw Octstring
func k519ControlOut2DoorPsw(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k519ControlOut2DoorMaxTalkLen k519_control_out2_door_max_talk_len int
func k519ControlOut2DoorMaxTalkLen(value interface{}) error {
	//秒数
	return nil
}

//k519PnmSwitch k519_pnm_switch int
func k519PnmSwitch(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k519PnmServerIp k519_pnm_server_ip Octstring
func k519PnmServerIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519PnmServerPort k519_pnm_server_port int
func k519PnmServerPort(value interface{}) error {
	//0-65535
	return nil
}

//k519PnmAreaName k519_pnm_area_name Octstring
func k519PnmAreaName(value interface{}) error {
	//最大字符串长度：20
	return nil
}

//k519PnmDeviceName k519_pnm_device_name Octstring
func k519PnmDeviceName(value interface{}) error {
	//最大字符串长度：20
	return nil
}

//k519PnmSelfDetectTimer k519_pnm_self_detect_timer int
func k519PnmSelfDetectTimer(value interface{}) error {
	//秒数
	return nil
}

//k519PnmSelfPinControlList k519_pnm_self_pin_control_list Octstring
func k519PnmSelfPinControlList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519PnmSelfPinNameList k519_pnm_self_pin_name_list Octstring
func k519PnmSelfPinNameList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k519DtmfRelayMode k519_dtmf_relay_mode int
func k519DtmfRelayMode(value interface{}) error {
	//0-2833 1-sipinfo
	return nil
}

//k519CodecTypeList k519_codec_type_list Octstring
func k519CodecTypeList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519HotlineEnable k519_hotline_enable int
func k519HotlineEnable(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k519HotlineNumber k519_hotline_number Octstring
func k519HotlineNumber(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519HotlineAccount k519_hotline_account int
func k519HotlineAccount(value interface{}) error {
	//0-5
	return nil
}

//k519QuickAttrib k519_quick_attrib Octstring
func k519QuickAttrib(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519QuickNumberList k519_quick_number_list Octstring
func k519QuickNumberList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k519QuickAutoList k519_quick_auto_list Octstring
func k519QuickAutoList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519QuickAccountList k519_quick_account_list Octstring
func k519QuickAccountList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519BoardcastIpList k519_boardcast_ip_list Octstring
func k519BoardcastIpList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k519BoardcastPortList k519_boardcast_port_list Octstring
func k519BoardcastPortList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519RecvBoardcastVoicePriority k519_recv_boardcast_voice_priority int
func k519RecvBoardcastVoicePriority(value interface{}) error {
	//0-11
	return nil
}

//k519RecvBoardcastEnableList k519_recv_boardcast_enable_list Octstring
func k519RecvBoardcastEnableList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519RecvBoardcastPriorityList k519_recv_boardcast_priority_list Octstring
func k519RecvBoardcastPriorityList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519RecvBoardcastIpList k519_recv_boardcast_ip_list Octstring
func k519RecvBoardcastIpList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k519RecvBoardcastPortList k519_recv_boardcast_port_list Octstring
func k519RecvBoardcastPortList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519MicrophoneVolume k519_microphone_volume int
func k519MicrophoneVolume(value interface{}) error {
	//1-9
	return nil
}

//k519SpeakerVolume k519_speaker_volume int
func k519SpeakerVolume(value interface{}) error {
	//0-9
	return nil
}

//k519HookonWaitTime k519_hookon_wait_time int
func k519HookonWaitTime(value interface{}) error {
	//1-30
	return nil
}

//k519RingStyle k519_ring_style int
func k519RingStyle(value interface{}) error {
	//1-7
	return nil
}

//k519RingVolume k519_ring_volume int
func k519RingVolume(value interface{}) error {
	//0-9
	return nil
}

//k519VideoMode k519_video_mode int
func k519VideoMode(value interface{}) error {
	//0-720p 1-480p
	return nil
}

//k519VideoPaytype k519_video_paytype int
func k519VideoPaytype(value interface{}) error {
	//96-127
	return nil
}

//k519SipAccountEnable k519_sip_account_enable int
func k519SipAccountEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k519SipRegisterSwitch k519_sip_register_switch int
func k519SipRegisterSwitch(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k519SipUnregisterSwitch k519_sip_unregister_switch int
func k519SipUnregisterSwitch(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k519SipRegisterExpire k519_sip_register_expire int
func k519SipRegisterExpire(value interface{}) error {
	//秒数
	return nil
}

//k519SipDisplayName k519_sip_display_name Octstring
func k519SipDisplayName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SipAccountName k519_sip_account_name Octstring
func k519SipAccountName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SipAuthName k519_sip_auth_name Octstring
func k519SipAuthName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SipAccountPassword k519_sip_account_password Octstring
func k519SipAccountPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SipRegServer0 k519_sip_reg_server0 Octstring
func k519SipRegServer0(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519SipRegPort0 k519_sip_reg_port0 int
func k519SipRegPort0(value interface{}) error {
	//0-65535
	return nil
}

//k519SipRegDomain0 k519_sip_reg_domain0 Octstring
func k519SipRegDomain0(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SipRegServer1 k519_sip_reg_server1 Octstring
func k519SipRegServer1(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519SipRegPort1 k519_sip_reg_port1 int
func k519SipRegPort1(value interface{}) error {
	//0-65535
	return nil
}

//k519SipRegDomain1 k519_sip_reg_domain1 Octstring
func k519SipRegDomain1(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SipRegServer2 k519_sip_reg_server2 Octstring
func k519SipRegServer2(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k519SipRegPort2 k519_sip_reg_port2 int
func k519SipRegPort2(value interface{}) error {
	//0-65535
	return nil
}

//k519SipRegDomain2 k519_sip_reg_domain2 Octstring
func k519SipRegDomain2(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SipUserAgent k519_sip_user_agent Octstring
func k519SipUserAgent(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k519SipHeartBeatEnable k519_sip_heart_beat_enable int
func k519SipHeartBeatEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k519SipHeartBeatTime k519_sip_heart_beat_time int
func k519SipHeartBeatTime(value interface{}) error {
	//秒数
	return nil
}

//k519SipAutoAnswerEnable k519_sip_auto_answer_enable int
func k519SipAutoAnswerEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k519SipAutoAnswerTime k519_sip_auto_answer_time int
func k519SipAutoAnswerTime(value interface{}) error {
	//秒数
	return nil
}

//k519SipUserParamSwitch k519_sip_user_param_switch int
func k519SipUserParamSwitch(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k519DigitRule15 k519_digit_rule_1_5 Octstring
func k519DigitRule15(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule610 k519_digit_rule_6_10 Octstring
func k519DigitRule610(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule1115 k519_digit_rule_11_15 Octstring
func k519DigitRule1115(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule1620 k519_digit_rule_16_20 Octstring
func k519DigitRule1620(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule2125 k519_digit_rule_21_25 Octstring
func k519DigitRule2125(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule2630 k519_digit_rule_26_30 Octstring
func k519DigitRule2630(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule3135 k519_digit_rule_31_35 Octstring
func k519DigitRule3135(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule3640 k519_digit_rule_36_40 Octstring
func k519DigitRule3640(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule4145 k519_digit_rule_41_45 Octstring
func k519DigitRule4145(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k519DigitRule4650 k519_digit_rule_46_50 Octstring
func k519DigitRule4650(value interface{}) error {
	//最大字符串长度：1200
	return nil
}
