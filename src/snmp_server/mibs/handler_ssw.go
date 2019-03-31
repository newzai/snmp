package mibs

//ksswNetMode kssw_net_mode int
func ksswNetMode(value interface{}) error {
	//0-静态 1-dhcp
	return nil
}

//ksswNetStaticIp kssw_net_static_ip Octstring
func ksswNetStaticIp(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//ksswNetStaticGateway kssw_net_static_gateway Octstring
func ksswNetStaticGateway(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//ksswNetStaticMask kssw_net_static_mask Octstring
func ksswNetStaticMask(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//ksswMasterDns kssw_master_dns Octstring
func ksswMasterDns(value interface{}) error {
	//最大字符串长度：40
	return nil
}

//ksswSlaveDns kssw_slave_dns Octstring
func ksswSlaveDns(value interface{}) error {
	//最大字符串长度：40
	return nil
}
