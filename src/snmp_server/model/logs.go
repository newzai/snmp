package model

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-xorm/xorm"
)

var engines sync.Map

//ShowSQL show sql for log
var ShowSQL = false

//LogInfo 日志信息
type LogInfo struct {
	ID       int64     `xorm:"pk autoincr 'id' " json:"id"`
	User     string    `xorm:"user" json:"username"`
	NTID     string    `xorm:"ntid" json:"ntid"`
	Event    string    `xorm:"event" json:"event"`
	SubEvent string    `xorm:"sub_event" json:"sub_event"`
	Time     time.Time `xorm:"ts" json:"timestamp"`
	Info     string    `xorm:"info" json:"info"`
}

//Insert insert into db
func (r *LogInfo) Insert() error {
	r.Time = time.Now()
	engine, err := getEngine(r.Time)
	if err != nil {
		return err
	}
	_, err = engine.InsertOne(r)
	return err
}

//SystemStartLog 系统启动日志
func SystemStartLog() {
	log := LogInfo{
		User:     "system",
		NTID:     "NA",
		Event:    "start",
		SubEvent: "start",
		Info:     "system start",
	}
	log.Insert()
}

//SystemStopLog 系统停止日志
func SystemStopLog(info string) {
	log := LogInfo{
		User:     "system",
		NTID:     "NA",
		Event:    "stop",
		SubEvent: "stop",
		Info:     info,
	}
	log.Insert()
}

//UserLoginLog 用户登录记录
func UserLoginLog(user string, info string) {
	log := LogInfo{
		User:     user,
		NTID:     "NA",
		Event:    "login",
		SubEvent: "login",
		Info:     info,
	}
	log.Insert()
}

//UserLogoutLog 用户登录记录
func UserLogoutLog(user string, info string) {
	log := LogInfo{
		User:     user,
		NTID:     "NA",
		Event:    "logout",
		SubEvent: "timeout",
		Info:     info,
	}
	log.Insert()
}

func dbFileIsExist(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//getEngine 根据时间获取到 db engine
func getEngine(t time.Time) (*xorm.Engine, error) {
	key := t.Format("2006-01")
	value, ok := engines.Load(key)
	if ok {
		return value.(*xorm.Engine), nil
	}

	dbfile := fmt.Sprintf("./log_%s.db", key)
	hasFile := dbFileIsExist(dbfile)
	engine, err := xorm.NewEngine("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	if !hasFile {
		engine.Sync2(new(LogInfo))
	}
	engine.DatabaseTZ = time.Local
	engine.TZLocation = time.Local
	engine.ShowSQL(ShowSQL)
	engines.Store(key, engine)
	return engine, nil
}

//GetLogEngine get engine by key
func GetLogEngine(key string) (*xorm.Engine, error) {
	value, ok := engines.Load(key)
	if ok {
		return value.(*xorm.Engine), nil
	}
	dbfile := fmt.Sprintf("./log_%s.db", key)
	if dbFileIsExist(dbfile) {
		engine, err := xorm.NewEngine("sqlite3", dbfile)
		if err != nil {
			return nil, err
		}
		engine.ShowSQL(ShowSQL)
		engines.Store(key, engine)
		engine.DatabaseTZ = time.Local
		engine.TZLocation = time.Local
		return engine, nil
	}
	return nil, errors.New("DB file not exist")

}
