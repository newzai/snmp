package model

import (
	"fmt"
	"net"
	"snmp_server/allocateid"
	"time"

	"github.com/cihub/seelog"

	"github.com/go-xorm/xorm"
)

var terminalUpdateCols = []string{"name", "path", "parent", "type", "version", "ip", "port", "keepalive", "service_status"}
var terminalUpdateKeepaliveCols = []string{"ip", "port", "keepalive", "service_status"}

//Terminal 设备信息
type Terminal struct {
	ID            int       `xorm:"pk 'id' "`
	Name          string    `xorm:"name"`
	Path          string    `xorm:"path"`
	Parent        int       `xorm:"parent"`
	NTID          string    `xorm:"varchar(255) notnull unique 'ntid'"`
	Type          string    `xorm:"type"`
	Version       string    `xorm:"ver"`
	IP            string    `xorm:"ip"`
	Port          int       `xorm:"port"`
	LastKeepalive time.Time `xorm:"keepalive"`
	X             int       `xorm:"'x'"`
	Y             int       `xorm:"'y'"`
	ServiceStatus int       `xorm:" 'service_status' "`
}

//IsOnline get set
func (r *Terminal) IsOnline() bool {
	now := time.Now()
	duration := now.Sub(r.LastKeepalive)
	seelog.Infof("%s duration: %d,  last %s, now %s", r.NTID, duration, r.LastKeepalive.String(), now.String())
	if duration < time.Second*60 {
		return true
	}
	return false
}

//Remote get udp address
func (r *Terminal) Remote() *net.UDPAddr {
	addr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", r.IP, r.Port))
	if err != nil {
		panic(err)
	}
	return addr
}

//GetAllTerminals 返回所有设备
func GetAllTerminals(engine *xorm.Engine) ([]*Terminal, error) {
	var t Terminal
	rows, err := engine.Rows(t)
	if err != nil {
		return nil, err
	}

	var ts []*Terminal
	defer rows.Close()
	for rows.Next() {
		tmp := new(Terminal)
		rows.Scan(tmp)
		ts = append(ts, tmp)
	}
	return ts, nil
}

//GetTerminalByParent get by parent..
func GetTerminalByParent(parent int, engin *xorm.Engine) ([]*Terminal, error) {
	var t Terminal
	rows, err := engin.Where("parent=?", parent).Rows(t)
	if err != nil {
		return nil, err
	}

	var ts []*Terminal
	defer rows.Close()
	for rows.Next() {
		tmp := new(Terminal)
		rows.Scan(tmp)
		ts = append(ts, tmp)
	}
	return ts, nil
}

//GetTerminalByPathAndType get by path like
func GetTerminalByPathAndType(path string, devType string, engine *xorm.Engine) ([]*Terminal, error) {
	var t Terminal
	rows, err := engine.Where("path like ? and type = ? ", path+"%", devType).Rows(t)
	if err != nil {
		return nil, err
	}

	var ts []*Terminal
	defer rows.Close()
	for rows.Next() {
		tmp := new(Terminal)
		rows.Scan(tmp)
		ts = append(ts, tmp)
	}
	return ts, nil
}

//GetTerminalByID get by terminal id
func GetTerminalByID(id int, engin *xorm.Engine) (*Terminal, error) {

	var t Terminal
	t.ID = id
	has, err := engin.Get(&t)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("terminal (%d) not exist", id)
	}
	return &t, err
}

//GetTerminalByNTID get by ntid
func GetTerminalByNTID(ntid string, engin *xorm.Engine) (*Terminal, error) {
	t := new(Terminal)
	_, err := engin.Where("ntid=?", ntid).Limit(1).Get(t)
	if t.ID > 0 {
		return t, nil
	}
	return nil, err
}

//CreateTerminal create terminal
func CreateTerminal(t *Terminal, engin *xorm.Engine) error {
	t.ID = allocateid.AllocateID()
	parent, err := CreatePath(t.Path, engin)
	if err != nil {
		return err
	}
	t.Parent = parent
	_, err = engin.InsertOne(t)
	return err
}

//UpdateTerminal update
func UpdateTerminal(t *Terminal, keepalive bool, engine *xorm.Engine) error {
	if keepalive {
		affected, err := engine.Id(t.ID).Cols(terminalUpdateKeepaliveCols...).Update(t)
		seelog.Info("update  keepalive terminal affected :", affected, " error:", err)
		return err
	}
	affected, err := engine.Id(t.ID).Cols(terminalUpdateCols...).Update(t)
	seelog.Info("update terminal affected:", affected, " error:", err)
	return err
}

//UpdateTerminalXY 更新x，y坐标轴
func UpdateTerminalXY(t *Terminal, engine *xorm.Engine) error {
	_, err := engine.Id(t.ID).Cols("x", "y").Update(t)
	return err
}
