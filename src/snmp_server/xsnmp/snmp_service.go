package xsnmp

import (
	"errors"
	"fmt"
	"net"
	"runtime/debug"
	"snmp_server/mibs"
	"sync"
	"time"

	"github.com/cihub/seelog"

	"github.com/soniah/gosnmp"
)

//OnTrap recv trap callback
type OnTrap func(packet *gosnmp.SnmpPacket, remote *net.UDPAddr)

//OnResponse for resposne
type OnResponse func(packet *gosnmp.SnmpPacket, remote *net.UDPAddr)

type Service struct {
	conn     *net.UDPConn
	onTrap   OnTrap
	requests sync.Map
	sequence uint32
	mutex    sync.Mutex
}

func (r *Service) makeSnmpPacket(pdutype gosnmp.PDUType, pdus []gosnmp.SnmpPDU) *gosnmp.SnmpPacket {
	return &gosnmp.SnmpPacket{
		Version:    gosnmp.Version2c,
		Community:  "public",
		Error:      0,
		ErrorIndex: 0,
		PDUType:    pdutype,
		RequestID:  r.GetSequence(),
		Variables:  pdus,
	}
}

//Get get snmp
func (r *Service) Get(input map[string]interface{}, index int, udpAddr *net.UDPAddr) (map[string]interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err)
			seelog.Error(string(debug.Stack()))
		}
	}()
	var result map[string]interface{}

	pdus := mibs.JSON2PDU(input, index, gosnmp.GetRequest)
	if len(pdus) == 0 {
		seelog.Warn("pdu is empty for get")
		return nil, errors.New("pdu is empty")
	}
	packet := r.makeSnmpPacket(gosnmp.GetRequest, pdus)
	data, err := packet.MarshalMsg()
	if err != nil {
		seelog.Error(err)
		return nil, err
	}

	resultChan := make(chan bool)
	defer close(resultChan)
	onRsp := func(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
		result = mibs.PDU2JSON(packet.Variables)
		resultChan <- true
	}
	r.requests.Store(packet.RequestID, onRsp) //add response hander for SNMP service ,when receive response callback
	for retry := 0; retry < 3; retry++ {
		_, err = r.conn.WriteToUDP(data, udpAddr)
		if err != nil {
			return nil, err
		}
		seelog.Infof("send snmp get retry(%d) to %s ,wait for response", retry, udpAddr.String())
		select {
		case <-resultChan:
			seelog.Infof("snmp get response ok retry(%d)", retry)
			return result, nil
		case <-time.After(time.Second * 3):
			seelog.Infof("snmp get response timeout retry(%d)", retry)
		}

	}
	r.requests.Delete(packet.RequestID)
	return nil, errors.New("get timeout")
}

//Set get snmp
func (r *Service) Set(input map[string]interface{}, index int, udpAddr *net.UDPAddr) (map[string]interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err)
			seelog.Error(string(debug.Stack()))
		}
	}()
	err := mibs.CheckSetPDUs(input)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}

	pdus := mibs.JSON2PDU(input, index, gosnmp.SetRequest)
	if len(pdus) == 0 {
		seelog.Warn("pdu is empty for set")
		return nil, errors.New("pdu is empty")
	}
	packet := r.makeSnmpPacket(gosnmp.SetRequest, pdus)
	data, err := packet.MarshalMsg()
	if err != nil {
		seelog.Error(err)
		return nil, err
	}

	resultChan := make(chan bool)
	defer close(resultChan)
	onRsp := func(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
		result = mibs.PDU2JSON(packet.Variables)
		resultChan <- true
	}
	r.requests.Store(packet.RequestID, onRsp) //add response hander for SNMP service ,when receive response callback

	for retry := 0; retry < 3; retry++ {
		_, err = r.conn.WriteToUDP(data, udpAddr)
		if err != nil {
			return nil, err
		}
		seelog.Infof("send snmp set retry(%d) to %s ,wait for response", retry, udpAddr.String())
		select {
		case <-resultChan:
			seelog.Infof("snmp set response ok retry(%d)", retry)
			return result, nil
		case <-time.After(time.Second * 3):
			seelog.Infof("snmp set response timeout retry(%d)", retry)
		}

	}
	r.requests.Delete(packet.RequestID)
	return nil, errors.New("set timeout")

}

//Start start
func (r *Service) Start(ip string, port uint16, onTrap OnTrap) {

	r.onTrap = onTrap

	r.listen(ip, port)
}

func (r *Service) unmarshalTrap(trap []byte) (result *gosnmp.SnmpPacket) {
	defer func() {
		if err := recover(); err != nil {

			seelog.Errorf(" something wrong, %s \n %s ", err, string(debug.Stack()))
			seelog.Flush()

		}
	}()

	return gosnmp.Default.UnmarshalTrap(trap)

}

//Listen listen
func (r *Service) listen(ip string, port uint16) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	seelog.Warnf("start service at %s:%d ok ", ip, port)
	r.conn = conn

	for {
		var buf [4096]byte
		rlen, remote, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			panic(err)
		}
		msg := buf[:rlen]
		snmpPacket := r.unmarshalTrap(msg)
		if snmpPacket != nil {
			switch snmpPacket.PDUType {
			case gosnmp.Trap:
				fallthrough
			case gosnmp.Report:
				fallthrough
			case gosnmp.SNMPv2Trap:
				go r.onTrap(snmpPacket, remote)
			default:
				//is resposne
				value, ok := r.requests.Load(snmpPacket.RequestID)
				if ok {
					r.requests.Delete(snmpPacket.RequestID)
					handler := value.(func(*gosnmp.SnmpPacket, *net.UDPAddr))
					if handler != nil {
						go handler(snmpPacket, remote)
					}
				}
			}
		}
	}

}

//GetSequence get sequence
func (r *Service) GetSequence() uint32 {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.sequence++
	return r.sequence
}

//Default Default
var Default = &Service{}
