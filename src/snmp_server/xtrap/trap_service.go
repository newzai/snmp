package xtrap

import (
	"fmt"
	"net"
	"runtime/debug"

	"github.com/cihub/seelog"

	"github.com/soniah/gosnmp"
)

const trapID = ".1.3.6.1.6.3.1.1.4.1.0"

//OnTrap recv trap callback
type OnTrap func(packet *gosnmp.SnmpPacket, remote *net.UDPAddr)

var handlers map[string]OnTrap

//RegisterHandler register handler
func RegisterHandler(trapType string, handler OnTrap) {

	if _, ok := handlers[trapType]; ok {
		panic(fmt.Sprintf("trap(%s) handler duplicate", trapType))
	}

	handlers[trapType] = handler
}

func getTrapType(packet *gosnmp.SnmpPacket) string {
	trapType := ""
	for _, vb := range packet.Variables {
		seelog.Infof("vb [%s]", vb.Name)
		if vb.Name == trapID {
			trapType = vb.Value.(string)
			seelog.Infof("trapType is [%s] ", trapType)
			break
		}
	}
	return trapType
}

//OnTrapHandler OnTrap
func OnTrapHandler(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	defer func() {
		if err := recover(); err != nil {
			seelog.Error(err, " stack: ", string(debug.Stack()))
		}
	}()

	trapType := getTrapType(packet)
	if trapType == "" {
		seelog.Warnf("can't get trap type %s", packet)
		return
	}

	handler, ok := handlers[trapType]
	if ok {
		handler(packet, remote)
	} else {
		seelog.Warnf("can't find trap handler %s", trapType)
	}
}

func init() {

	handlers = make(map[string]OnTrap)

}
