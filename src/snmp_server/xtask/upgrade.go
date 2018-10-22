package xtask

import (
	"errors"
	"time"

	"github.com/go-xorm/xorm"
)

//UpgradeResult 升级结果
type UpgradeResult int

const (
	//Downing 正在升级
	Downing UpgradeResult = 0
	//Upgrading Upgrading
	Upgrading UpgradeResult = 1
	//OK ok
	OK UpgradeResult = 2
	//ERROR error
	ERROR UpgradeResult = 3
	//Timeout timeout
	Timeout UpgradeResult = 4
)

func (r UpgradeResult) String() string {
	switch r {
	case Downing:
		return "Downing"
	case Upgrading:
		return "Upgrading"
	case OK:
		return "OK"
	case ERROR:
		return "ERROR"
	case Timeout:
		return "Timeout"
	default:
		return "N/A"
	}
}

//Upgrade ftp 软件升级任务
type Upgrade struct {
	ID        int64     `xorm:" pk autoincr 'id' " json:"-"`
	TID       int       `xorm:" unique 'tid' " json:"itemid"`                   // from model.Terminal.ID
	NTID      string    `xorm:"varchar(255) notnull unique 'ntid'" json:"ntid"` // from model.Terminal.NTID
	Name      string    `xorm:" notnull 'name'" json:"itemname"`                // from model.Terminal.Name
	Path      string    `xorm:" notnull 'path'" json:"itempath"`                // from model.Terminal.Path
	Type      string    `xorm:" notnull 'type'" json:"dev_type"`                // from model.Terminal.Type
	Completed bool      `xorm:" notnull 'completed'" json:"completed"`          // true  ftp upgrade is completed
	Result    int       `xorm:" notnull 'result'" json:"result"`                // Upgrading, OK, Error,Timeout
	Progress  int       `xorm:" notnull 'progress'" json:"progress"`
	Reason    string    `xorm:" notnull 'reason'" json:"reason"` // for Result is  Error,
	Last      time.Time `xorm:" notnull 'last'" json:"last"`     // 最后一次操作时间
}

//Insert insert to db
func (r *Upgrade) Insert(engine *xorm.Engine) error {
	r.Last = time.Now()
	_, err := engine.InsertOne(r)
	return err
}

//Update update
func (r *Upgrade) Update(engine *xorm.Engine) error {
	r.Last = time.Now()
	_, err := engine.Id(r.ID).Cols("completed", "result", "progress", "reason", "last").Update(r)
	return err
}

//Delete delete from db
func (r *Upgrade) Delete(engine *xorm.Engine) error {
	_, err := engine.Delete(r)
	return err
}

//Get get by Ntid
func (r *Upgrade) Get(engine *xorm.Engine) (bool, error) {

	if len(r.NTID) > 0 {
		ntid := r.NTID
		r.NTID = ""
		has, err := engine.Where("ntid=?", ntid).Get(r)
		return has, err
	} else if r.TID > 0 {
		tid := r.TID
		r.TID = 0
		has, err := engine.Where("tid=?", tid).Get(r)
		return has, err
	} else {
		return false, errors.New("missing TID or NTID")
	}
}

//Find find by upgrade
func (r *Upgrade) Find(engine *xorm.Engine) ([]*Upgrade, error) {
	if r.Path == "" || r.Type == "" {
		return nil, errors.New("missing Path or Type")
	}

	upgrades := make([]*Upgrade, 0)
	err := engine.Where("path like ? and type=?", r.Path+"%", r.Type).Find(&upgrades)
	return upgrades, err
}

//InitDatabase init
func InitDatabase(engine *xorm.Engine) {

	engine.Sync2(new(Upgrade))
}
