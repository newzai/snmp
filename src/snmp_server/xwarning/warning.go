package xwarning

import (
	"fmt"
	"time"

	"github.com/cihub/seelog"

	"github.com/go-xorm/xorm"
)

//Warning 告警数据
type Warning struct {
	ID      int64     `xorm:" pk autoincr 'id' " json:"id"` //数据库自增
	TID     int       `xorm:"tid" json:"tid"`               //设备 Terminal ID
	NTID    string    `xorm:"ntid" json:"ntid"`             // 设备 NTID
	TName   string    `xorm:"tname" json:"tname"`           // 设备 Name
	Path    string    `xorm:"path" json:"path"`             //设备 Path
	WType   string    `xorm:"wtype" json:"wtype"`           //告警类型
	WValue  int       `xorm:"wvalue" json:"wvalue"`         //告警值
	WDemo   string    `xorm:"wdemo" json:"wdemo"`           //告警描述信息
	OTime   time.Time `xorm:"otime" json:"otime"`           //告警发生时间
	Confirm bool      `xorm:"confirm" json:"confirm"`       //告警是否确认
	CTime   time.Time `xorm:"ctime" json:"ctime"`           //告警确认时间
	CInfo   string    `xorm:"cinfo" json:"cinfo"`           //告警确认信息
}

//HistoryWarning 历史告警数据
type HistoryWarning struct {
	ID      int64     `xorm:" pk autoincr 'id' " json:"id"` //数据库自增
	TID     int       `xorm:"tid" json:"tid"`               //设备 Terminal ID
	NTID    string    `xorm:"ntid" json:"ntid"`             // 设备 NTID
	Path    string    `xorm:"path" json:"path"`             //设备 Path
	TName   string    `xorm:"tname" json:"tname"`           // 设备 Name
	WType   string    `xorm:"wtype" json:"wtype"`           //告警类型
	WValue  int       `xorm:"wvalue" json:"wvalue"`         //告警值
	WDemo   string    `xorm:"wdemo" json:"wdemo"`           //告警描述信息
	OTime   time.Time `xorm:"otime" json:"otime"`           //告警发生时间
	Confirm bool      `xorm:"confirm" json:"confirm"`       //告警是否确认
	CTime   time.Time `xorm:"ctime" json:"ctime"`           //告警确认时间
	CInfo   string    `xorm:"cinfo" json:"cinfo"`           //告警确认信息
	ETime   time.Time `xorm:"etime" json:"etime"`           //告警结束时间，（end，clear）
}

//GetWarning get warning by  ntid,wtype
func GetWarning(ntid string, wtype string, engine *xorm.Engine) (*Warning, bool) {
	tmpW := new(Warning)
	has, _ := engine.Where("ntid=? and wtype=?", ntid, wtype).Get(tmpW)
	return tmpW, has
}

//InsertWarning 插入告警
func InsertWarning(w *Warning, engine *xorm.Engine) error {

	tmpW := new(Warning)
	has, err := engine.Where("ntid=? and wtype=?", w.NTID, w.WType).Get(tmpW)
	if err != nil {
		return err
	}
	if has {
		//告警已经存在
		return nil
	}
	w.OTime = time.Now()

	_, err = engine.InsertOne(w)
	return err
}

//ConfirmWarning 告警确认
func ConfirmWarning(wid uint64, info string, engine *xorm.Engine) error {

	w := new(Warning)
	has, err := engine.Id(wid).Get(w)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("Warning(%d) not exist", wid)
	}
	w.Confirm = true
	w.CInfo = info
	w.CTime = time.Now()
	_, err = engine.Cols("confirm", "ctime", "cinfo").Update(w)
	return err
}

//ClearWarning clear warning
func ClearWarning(nitd string, wtype string, engine *xorm.Engine) {
	w, has := GetWarning(nitd, wtype, engine)
	if !has {
		return
	}
	seelog.Warnf("clear wanring %s-%s", nitd, wtype)
	engine.Delete(w)
	hw := HistoryWarning{
		TID:     w.TID,
		NTID:    w.NTID,
		TName:   w.TName,
		Path:    w.Path,
		WType:   w.WType,
		WValue:  w.WValue,
		OTime:   w.OTime,
		Confirm: w.Confirm,
		CInfo:   w.CInfo,
		CTime:   w.CTime,
		ETime:   time.Now(),
	}

	engine.InsertOne(hw)
}

//GetWarnings get warning  if id > wid
func GetWarnings(wid uint64, engine *xorm.Engine) ([]*Warning, error) {
	w := new(Warning)
	rows, err := engine.Where("id > ?", wid).OrderBy("id").Limit(2).Rows(w)
	if err != nil {
		return nil, err
	}

	warnings := make([]*Warning, 0, 2)
	for rows.Next() {
		ww := new(Warning)
		rows.Scan(ww)
		warnings = append(warnings, ww)
	}
	return warnings, nil
}

//GetWarningsByPath get warning  by path
func GetWarningsByPath(path string, engine *xorm.Engine) ([]*Warning, error) {
	w := new(Warning)
	rows, err := engine.Where("path like ?", path+"%").Rows(w)
	if err != nil {
		return nil, err
	}

	warnings := make([]*Warning, 0, 2)
	for rows.Next() {
		ww := new(Warning)
		rows.Scan(ww)
		warnings = append(warnings, ww)
	}
	return warnings, nil
}

//GetWarningsByTID get warning  by path
func GetWarningsByTID(tid int, engine *xorm.Engine) ([]*Warning, error) {
	w := new(Warning)
	rows, err := engine.Where("tid = ?", tid).Rows(w)
	if err != nil {
		return nil, err
	}

	warnings := make([]*Warning, 0, 2)
	for rows.Next() {
		ww := new(Warning)
		rows.Scan(ww)
		warnings = append(warnings, ww)
	}
	return warnings, nil
}

//InitDatabase init warning db
func InitDatabase(engine *xorm.Engine) {

	engine.Sync2(new(Warning))
}
