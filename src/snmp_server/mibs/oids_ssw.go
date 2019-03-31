package mibs

import "github.com/soniah/gosnmp"

var oidSSW = []OIDAttr{
	OIDAttr{Name: "kssw_software_version", OID: "1800.53.10.1", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "kssw_phone_style", OID: "1800.53.10.2", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "kssw_current_time", OID: "1800.53.10.3", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "kssw_current_net_ip", OID: "1800.53.10.5", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "kssw_current_net_mac", OID: "1800.53.10.6", Type: gosnmp.OctetString, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "kssw_net0_link_status", OID: "1800.53.10.7", Type: gosnmp.Integer, ReadOnly: true, ValidHander: defaultValidHander},
	OIDAttr{Name: "kssw_net_mode", OID: "1800.53.11.1", Type: gosnmp.Integer, ReadOnly: false, ValidHander: ksswNetMode},
	OIDAttr{Name: "kssw_net_static_ip", OID: "1800.53.11.2", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: ksswNetStaticIp},
	OIDAttr{Name: "kssw_net_static_gateway", OID: "1800.53.11.3", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: ksswNetStaticGateway},
	OIDAttr{Name: "kssw_net_static_mask", OID: "1800.53.11.4", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: ksswNetStaticMask},
	OIDAttr{Name: "kssw_master_dns", OID: "1800.53.11.8", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: ksswMasterDns},
	OIDAttr{Name: "kssw_slave_dns", OID: "1800.53.11.9", Type: gosnmp.OctetString, ReadOnly: false, ValidHander: ksswSlaveDns},
}

func init() {
	for _, attr := range oidSSW {
		cAttr := attr
		cAttr.OID = oidPrefix + cAttr.OID

		jsonKeyAttr[cAttr.Name] = &cAttr
		oidKeyAttr[cAttr.OID] = &cAttr
	}
}
