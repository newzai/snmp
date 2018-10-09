package mibs

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//OIDAttr oid attribute
type OIDAttr struct {
	Name        string
	OID         string
	Type        gosnmp.Asn1BER
	ReadOnly    bool
	ValidHander func(interface{}) error
}

func defaultValidHander(value interface{}) error {
	return nil
}

const oidPrefix = ".1.3.6.1.4.1."

func checkoutFtpParam(jPdu map[string]interface{}) error {
	//usl_ftp_server_ip
	//usl_ftp_server_port
	//usl_ftp_user_name
	//usl_ftp_user_passwd
	//usl_ftp_file_size
	if _, ok := jPdu["usl_ftp_server_ip"]; !ok {
		return errors.New("missing usl_ftp_server_ip")
	}
	if _, ok := jPdu["usl_ftp_server_port"]; !ok {
		return errors.New("missing usl_ftp_server_port")
	}
	if _, ok := jPdu["usl_ftp_user_name"]; !ok {
		return errors.New("missing usl_ftp_user_name")
	}
	if _, ok := jPdu["usl_ftp_user_passwd"]; !ok {
		return errors.New("missing usl_ftp_user_passwd")
	}
	return nil
}

//CheckSetPDUs checkout value is valid
func CheckSetPDUs(jPdu map[string]interface{}) error {
	for key, value := range jPdu {
		attr, ok := jsonKeyAttr[key]
		if ok {
			switch attr.Type {
			case gosnmp.Integer:
				switch value.(type) {
				case int:
				case float32:
				case float64:
				case uint:
				default:
					return fmt.Errorf("oid[%s] type must int", key)

				}
			case gosnmp.OctetString:
				fallthrough
			case gosnmp.ObjectIdentifier:
				switch value.(type) {
				case string:
				default:
					return fmt.Errorf("oid[%s] type must string", key)
				}
			}
			err := attr.ValidHander(value)
			if err != nil {
				return fmt.Errorf("%s|%v", key, err)
			}
		} else {
			return fmt.Errorf("snmp oid[%s] not exist", key)
		}
	}

	//usl_ftp_soft_file_name   usl_ftp_save_cfg_file_name   usl_ftp_restore_cfg_file_name checkout
	if _, ok := jPdu["usl_ftp_soft_file_name"]; ok {
		err := checkoutFtpParam(jPdu)
		if err != nil {
			return err
		}
	}
	if _, ok := jPdu["usl_ftp_save_cfg_file_name"]; ok {
		err := checkoutFtpParam(jPdu)
		if err != nil {
			return err
		}
	}
	if _, ok := jPdu["usl_ftp_restore_cfg_file_name"]; ok {
		err := checkoutFtpParam(jPdu)
		if err != nil {
			return err
		}
	}
	return nil
}

//JSON2PDU from http json snmp json key/value to  SnmpPDUs
func JSON2PDU(jPdu map[string]interface{}, index int, pdutype gosnmp.PDUType) []gosnmp.SnmpPDU {

	inst := fmt.Sprintf(".%d", index)
	pdus := make([]gosnmp.SnmpPDU, 0, len(jPdu))
	for name, value := range jPdu {
		attr, ok := jsonKeyAttr[name]
		if ok {
			if pdutype == gosnmp.GetRequest || pdutype == gosnmp.GetNextRequest || pdutype == gosnmp.GetBulkRequest {
				pdu := gosnmp.SnmpPDU{
					Name:  attr.OID + inst,
					Type:  gosnmp.Null,
					Value: nil,
				}
				seelog.Infof("add get pdu: %s-->%s ", name, pdu.Name)
				pdus = append(pdus, pdu)
			} else {
				//string to json解析后 整数的类型被修改为 float64
				var pduValue interface{}
				switch value.(type) {
				case float64:
					if attr.Type == gosnmp.Integer {
						pduValue = int(value.(float64))
					} else if attr.Type == gosnmp.Uinteger32 {
						pduValue = uint32(value.(float64))
					} else {
						pduValue = uint32(value.(float64))
					}
				default:
					pduValue = value
				}

				//for set
				pdu := gosnmp.SnmpPDU{
					Name:  attr.OID + inst,
					Type:  attr.Type,
					Value: pduValue,
				}
				if attr.ReadOnly {
					seelog.Infof("ignore set pdu %s-->%s", name, pdu.Name)
					continue
				}
				seelog.Infof("add set pdu: %s-->%s |%v  pdutype(%s) value type(%s)", name, pdu.Name, pduValue, pdu.Type, reflect.TypeOf(pdu.Value).Name())
				pdus = append(pdus, pdu)

			}
		} else {
			seelog.Warnf("[JSON2PDU] can't find attr by key:%s", name)
		}
	}

	seelog.Info("rsp pdus:", pdus)

	return pdus
}

//PDU2JSON pdu value to json
func PDU2JSON(pdus []gosnmp.SnmpPDU) map[string]interface{} {
	jPdus := make(map[string]interface{})

	for _, pdu := range pdus {
		idx := strings.LastIndex(pdu.Name, ".")
		name := pdu.Name[:idx]
		seelog.Info("PDU2JSON name ", pdu.Name, "-->", name)
		attr, ok := oidKeyAttr[name]
		if ok {
			seelog.Infof("pdu %s value %v", attr.Name, pdu.Value)
			if pdu.Type == gosnmp.OctetString {
				jPdus[attr.Name] = string(pdu.Value.([]byte))
			} else {
				jPdus[attr.Name] = pdu.Value
			}
		} else {
			seelog.Warnf("[PDU2JSON] can't find attr by oid: %s,%s", pdu.Name, name)
		}

	}
	return jPdus
}

var jsonKeyAttr map[string]*OIDAttr
var oidKeyAttr map[string]*OIDAttr

func init() {

	jsonKeyAttr = make(map[string]*OIDAttr)
	oidKeyAttr = make(map[string]*OIDAttr)
}
