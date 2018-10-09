package mibs

import (
	"errors"
	"fmt"
	"reflect"
)

//k518NetMode k518_net_mode int
func k518NetMode(value interface{}) error {
	//0-静态 1-dhcp 2-pppoe
	switch v := value.(type) {
	case int:
		if !(v == 0 || v == 1 || v == 2) {
			return errors.New("k518_net_mode value not in range 0-static 1-dhcp 2-pppoe")
		}

	case float64:
		if !(int(v) == 0 || int(v) == 1 || int(v) == 2) {
			return errors.New("k518_net_mode value not in range 0-static 1-dhcp 2-pppoe")
		}

	default:
		return fmt.Errorf("k518_net_mode value type is not int , is %s", reflect.TypeOf(value).Name())
	}
	return nil
}

//k518NetStaticIp k518_net_static_ip Octstring
func k518NetStaticIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518NetStaticGateway k518_net_static_gateway Octstring
func k518NetStaticGateway(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518NetStaticMask k518_net_static_mask Octstring
func k518NetStaticMask(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518PppoeAccount k518_pppoe_account Octstring
func k518PppoeAccount(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518PppoePassword k518_pppoe_password Octstring
func k518PppoePassword(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518DnsMode k518_dns_mode int
func k518DnsMode(value interface{}) error {
	//0-关闭 1-开启
	return nil
}

//k518MasterDns k518_master_dns Octstring
func k518MasterDns(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518SlaveDns k518_slave_dns Octstring
func k518SlaveDns(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518Net0Adptive k518_net0_adptive int
func k518Net0Adptive(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k518Net0Rate k518_net0_rate int
func k518Net0Rate(value interface{}) error {
	//0-10M 2-100M
	return nil
}

//k518Net0Mode k518_net0_mode int
func k518Net0Mode(value interface{}) error {
	//0-半工 1-双工
	return nil
}

//k518Net1Adptive k518_net1_adptive int
func k518Net1Adptive(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k518Net1Rate k518_net1_rate int
func k518Net1Rate(value interface{}) error {
	//0-10M 2-100M
	return nil
}

//k518Net1Mode k518_net1_mode int
func k518Net1Mode(value interface{}) error {
	//0-半工 1-双工
	return nil
}

//k518AdminName k518_admin_name Octstring
func k518AdminName(value interface{}) error {
	//最大字符串长度：32
	return nil
}

//k518AdminPassword k518_admin_password Octstring
func k518AdminPassword(value interface{}) error {
	//最大字符串长度：32
	return nil
}

//k518LanguageMode k518_language_mode int
func k518LanguageMode(value interface{}) error {
	//0-中文 1-英文 2-韩文
	return nil
}

//k518HttpMode k518_http_mode int
func k518HttpMode(value interface{}) error {
	//0-http 1-https
	return nil
}

//k518HttpPort k518_http_port int
func k518HttpPort(value interface{}) error {
	//0-65535
	return nil
}

//k518TelnetPort k518_telnet_port int
func k518TelnetPort(value interface{}) error {
	//0-65535
	return nil
}

//k518SipLocalPort k518_sip_local_port int
func k518SipLocalPort(value interface{}) error {
	//0-65535
	return nil
}

//k518LogOuputMode k518_log_ouput_mode int
func k518LogOuputMode(value interface{}) error {
	//0-串口 1-telnet 2-网页
	return nil
}

//k518NtpMode k518_ntp_mode int
func k518NtpMode(value interface{}) error {
	//0-关闭 1-开启
	return nil
}

//k518PrimaryNtpAddress k518_primary_ntp_address Octstring
func k518PrimaryNtpAddress(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518PrimaryNtpPort k518_primary_ntp_port int
func k518PrimaryNtpPort(value interface{}) error {
	//0-65535
	return nil
}

//k518SecondNtpAddress k518_second_ntp_address Octstring
func k518SecondNtpAddress(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518SecondNtpPort k518_second_ntp_port int
func k518SecondNtpPort(value interface{}) error {
	//0-65535
	return nil
}

//k518TimeZone k518_time_zone int
func k518TimeZone(value interface{}) error {
	//0-25
	return nil
}

//k518SnmpEnable k518_snmp_enable int
func k518SnmpEnable(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k518SnmpServerIp k518_snmp_server_ip Octstring
func k518SnmpServerIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518SnmpServerPort k518_snmp_server_port int
func k518SnmpServerPort(value interface{}) error {
	//0-65535
	return nil
}

//k518SnmpDeviceZone k518_snmp_device_zone Octstring
func k518SnmpDeviceZone(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SnmpDeviceName k518_snmp_device_name Octstring
func k518SnmpDeviceName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069WorkMode k518_tr069_work_mode int
func k518Tr069WorkMode(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k518Tr069AcsServer k518_tr069_acs_server Octstring
func k518Tr069AcsServer(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069AcsAuthmode k518_tr069_acs_authmode int
func k518Tr069AcsAuthmode(value interface{}) error {
	//
	return nil
}

//k518Tr069AcsAccount k518_tr069_acs_account Octstring
func k518Tr069AcsAccount(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069AcsPassword k518_tr069_acs_password Octstring
func k518Tr069AcsPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069ConnectPeriod k518_tr069_connect_period int
func k518Tr069ConnectPeriod(value interface{}) error {
	//
	return nil
}

//k518Tr069CpeAuthmode k518_tr069_cpe_authmode int
func k518Tr069CpeAuthmode(value interface{}) error {
	//
	return nil
}

//k518Tr069CpeReqName k518_tr069_cpe_req_name Octstring
func k518Tr069CpeReqName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069CpeReqPassword k518_tr069_cpe_req_password Octstring
func k518Tr069CpeReqPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069InformTimeout k518_tr069_inform_timeout int
func k518Tr069InformTimeout(value interface{}) error {
	//
	return nil
}

//k518Tr069Manufacturer k518_tr069_manufacturer Octstring
func k518Tr069Manufacturer(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069Oui k518_tr069_oui Octstring
func k518Tr069Oui(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069ProductClass k518_tr069_product_class Octstring
func k518Tr069ProductClass(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069SerialNamber k518_tr069_serial_namber Octstring
func k518Tr069SerialNamber(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069HardwareVersion k518_tr069_hardware_version Octstring
func k518Tr069HardwareVersion(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069SoftwareVersion k518_tr069_software_version Octstring
func k518Tr069SoftwareVersion(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069ConreqPort k518_tr069_conreq_port int
func k518Tr069ConreqPort(value interface{}) error {
	//
	return nil
}

//k518Tr069ConregUri k518_tr069_conreg_uri Octstring
func k518Tr069ConregUri(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069ConreqUrl k518_tr069_conreq_url Octstring
func k518Tr069ConreqUrl(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069Dn0Number k518_tr069_dn0_number Octstring
func k518Tr069Dn0Number(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069Dn0Name k518_tr069_dn0_name Octstring
func k518Tr069Dn0Name(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518Tr069AutoProvision k518_tr069_auto_provision int
func k518Tr069AutoProvision(value interface{}) error {
	//
	return nil
}

//k518ControlOut1PutType k518_control_out1_put_type int
func k518ControlOut1PutType(value interface{}) error {
	//0-禁用 1-干接点 2-门禁
	return nil
}

//k518ControlOut1ConnectLev k518_control_out1_connect_lev int
func k518ControlOut1ConnectLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k518ControlOut1ConnectSec k518_control_out1_connect_sec int
func k518ControlOut1ConnectSec(value interface{}) error {
	//秒数
	return nil
}

//k518ControlOut1TriggerType k518_control_out1_trigger_type int
func k518ControlOut1TriggerType(value interface{}) error {
	//0-振铃 1-摘机 2-通话 3-振铃结束 4-通话结束
	return nil
}

//k518ControlOut1DoorLev k518_control_out1_door_lev int
func k518ControlOut1DoorLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k518ControlOut1DoorSec k518_control_out1_door_sec int
func k518ControlOut1DoorSec(value interface{}) error {
	//秒数
	return nil
}

//k518ControlOut1DoorNum k518_control_out1_door_num Octstring
func k518ControlOut1DoorNum(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k518ControlOut1DoorPsw k518_control_out1_door_psw Octstring
func k518ControlOut1DoorPsw(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k518ControlOut1DoorMaxTalkLen k518_control_out1_door_max_talk_len int
func k518ControlOut1DoorMaxTalkLen(value interface{}) error {
	//秒数
	return nil
}

//k518ControlOut2PutType k518_control_out2_put_type int
func k518ControlOut2PutType(value interface{}) error {
	//0-禁用 1-干接点 2-门禁
	return nil
}

//k518ControlOut2ConnectLev k518_control_out2_connect_lev int
func k518ControlOut2ConnectLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k518ControlOut2ConnectSec k518_control_out2_connect_sec int
func k518ControlOut2ConnectSec(value interface{}) error {
	//秒数
	return nil
}

//k518ControlOut2TriggerType k518_control_out2_trigger_type int
func k518ControlOut2TriggerType(value interface{}) error {
	//0-振铃 1-摘机 2-通话 3-振铃结束 4-通话结束
	return nil
}

//k518ControlOut2DoorLev k518_control_out2_door_lev int
func k518ControlOut2DoorLev(value interface{}) error {
	//0-低电平 1-高电平
	return nil
}

//k518ControlOut2DoorSec k518_control_out2_door_sec int
func k518ControlOut2DoorSec(value interface{}) error {
	//秒数
	return nil
}

//k518ControlOut2DoorNum k518_control_out2_door_num Octstring
func k518ControlOut2DoorNum(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k518ControlOut2DoorPsw k518_control_out2_door_psw Octstring
func k518ControlOut2DoorPsw(value interface{}) error {
	//最大字符串长度：6
	return nil
}

//k518ControlOut2DoorMaxTalkLen k518_control_out2_door_max_talk_len int
func k518ControlOut2DoorMaxTalkLen(value interface{}) error {
	//秒数
	return nil
}

//k518PnmSwitch k518_pnm_switch int
func k518PnmSwitch(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k518PnmServerIp k518_pnm_server_ip Octstring
func k518PnmServerIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518PnmServerPort k518_pnm_server_port int
func k518PnmServerPort(value interface{}) error {
	//0-65535
	return nil
}

//k518PnmAreaName k518_pnm_area_name Octstring
func k518PnmAreaName(value interface{}) error {
	//最大字符串长度：20
	return nil
}

//k518PnmDeviceName k518_pnm_device_name Octstring
func k518PnmDeviceName(value interface{}) error {
	//最大字符串长度：20
	return nil
}

//k518PnmSelfDetectTimer k518_pnm_self_detect_timer int
func k518PnmSelfDetectTimer(value interface{}) error {
	//秒数
	return nil
}

//k518PnmSelfPinControlList k518_pnm_self_pin_control_list Octstring
func k518PnmSelfPinControlList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518PnmSelfPinNameList k518_pnm_self_pin_name_list Octstring
func k518PnmSelfPinNameList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k518DtmfRelayMode k518_dtmf_relay_mode int
func k518DtmfRelayMode(value interface{}) error {
	//0-2833 1-sipinfo
	return nil
}

//k518CodecTypeList k518_codec_type_list Octstring
func k518CodecTypeList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518HotlineEnable k518_hotline_enable int
func k518HotlineEnable(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k518HotlineNumber k518_hotline_number Octstring
func k518HotlineNumber(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518HotlineAccount k518_hotline_account int
func k518HotlineAccount(value interface{}) error {
	//0-5
	return nil
}

//k518QuickAttrib k518_quick_attrib Octstring
func k518QuickAttrib(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518QuickNumberList k518_quick_number_list Octstring
func k518QuickNumberList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k518QuickAutoList k518_quick_auto_list Octstring
func k518QuickAutoList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518QuickAccountList k518_quick_account_list Octstring
func k518QuickAccountList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518BoardcastIpList k518_boardcast_ip_list Octstring
func k518BoardcastIpList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k518BoardcastPortList k518_boardcast_port_list Octstring
func k518BoardcastPortList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518RecvBoardcastVoicePriority k518_recv_boardcast_voice_priority int
func k518RecvBoardcastVoicePriority(value interface{}) error {
	//0-11
	return nil
}

//k518RecvBoardcastEnableList k518_recv_boardcast_enable_list Octstring
func k518RecvBoardcastEnableList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518RecvBoardcastPriorityList k518_recv_boardcast_priority_list Octstring
func k518RecvBoardcastPriorityList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518RecvBoardcastIpList k518_recv_boardcast_ip_list Octstring
func k518RecvBoardcastIpList(value interface{}) error {
	//最大字符串长度：200
	return nil
}

//k518RecvBoardcastPortList k518_recv_boardcast_port_list Octstring
func k518RecvBoardcastPortList(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518MicrophoneVolume k518_microphone_volume int
func k518MicrophoneVolume(value interface{}) error {
	//1-9
	return nil
}

//k518SpeakerVolume k518_speaker_volume int
func k518SpeakerVolume(value interface{}) error {
	//0-9
	return nil
}

//k518HookonWaitTime k518_hookon_wait_time int
func k518HookonWaitTime(value interface{}) error {
	//1-30
	return nil
}

//k518RingStyle k518_ring_style int
func k518RingStyle(value interface{}) error {
	//1-7
	return nil
}

//k518RingVolume k518_ring_volume int
func k518RingVolume(value interface{}) error {
	//0-9
	return nil
}

//k518SipAccountEnable k518_sip_account_enable int
func k518SipAccountEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k518SipRegisterSwitch k518_sip_register_switch int
func k518SipRegisterSwitch(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k518SipUnregisterSwitch k518_sip_unregister_switch int
func k518SipUnregisterSwitch(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k518SipRegisterExpire k518_sip_register_expire int
func k518SipRegisterExpire(value interface{}) error {
	//秒数
	return nil
}

//k518SipDisplayName k518_sip_display_name Octstring
func k518SipDisplayName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SipAccountName k518_sip_account_name Octstring
func k518SipAccountName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SipAuthName k518_sip_auth_name Octstring
func k518SipAuthName(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SipAccountPassword k518_sip_account_password Octstring
func k518SipAccountPassword(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SipRegServer0 k518_sip_reg_server0 Octstring
func k518SipRegServer0(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518SipRegPort0 k518_sip_reg_port0 int
func k518SipRegPort0(value interface{}) error {
	//0-65535
	return nil
}

//k518SipRegDomain0 k518_sip_reg_domain0 Octstring
func k518SipRegDomain0(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SipRegServer1 k518_sip_reg_server1 Octstring
func k518SipRegServer1(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518SipRegPort1 k518_sip_reg_port1 int
func k518SipRegPort1(value interface{}) error {
	//0-65535
	return nil
}

//k518SipRegDomain1 k518_sip_reg_domain1 Octstring
func k518SipRegDomain1(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SipRegServer2 k518_sip_reg_server2 Octstring
func k518SipRegServer2(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//k518SipRegPort2 k518_sip_reg_port2 int
func k518SipRegPort2(value interface{}) error {
	//0-65535
	return nil
}

//k518SipRegDomain2 k518_sip_reg_domain2 Octstring
func k518SipRegDomain2(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SipUserAgent k518_sip_user_agent Octstring
func k518SipUserAgent(value interface{}) error {
	//最大字符串长度：100
	return nil
}

//k518SipHeartBeatEnable k518_sip_heart_beat_enable int
func k518SipHeartBeatEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k518SipHeartBeatTime k518_sip_heart_beat_time int
func k518SipHeartBeatTime(value interface{}) error {
	//秒数
	return nil
}

//k518SipAutoAnswerEnable k518_sip_auto_answer_enable int
func k518SipAutoAnswerEnable(value interface{}) error {
	//0-启用  1-禁止
	return nil
}

//k518SipAutoAnswerTime k518_sip_auto_answer_time int
func k518SipAutoAnswerTime(value interface{}) error {
	//秒数
	return nil
}

//k518SipUserParamSwitch k518_sip_user_param_switch int
func k518SipUserParamSwitch(value interface{}) error {
	//0-禁止  1-启用
	return nil
}

//k518DigitRule15 k518_digit_rule_1_5 Octstring
func k518DigitRule15(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule610 k518_digit_rule_6_10 Octstring
func k518DigitRule610(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule1115 k518_digit_rule_11_15 Octstring
func k518DigitRule1115(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule1620 k518_digit_rule_16_20 Octstring
func k518DigitRule1620(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule2125 k518_digit_rule_21_25 Octstring
func k518DigitRule2125(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule2630 k518_digit_rule_26_30 Octstring
func k518DigitRule2630(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule3135 k518_digit_rule_31_35 Octstring
func k518DigitRule3135(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule3640 k518_digit_rule_36_40 Octstring
func k518DigitRule3640(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule4145 k518_digit_rule_41_45 Octstring
func k518DigitRule4145(value interface{}) error {
	//最大字符串长度：1200
	return nil
}

//k518DigitRule4650 k518_digit_rule_46_50 Octstring
func k518DigitRule4650(value interface{}) error {
	//最大字符串长度：1200
	return nil
}
