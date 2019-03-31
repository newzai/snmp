package model

import (
	"encoding/json"
	"fmt"
	"snmp_server/allocateid"
	"strings"

	"github.com/cihub/seelog"
	"github.com/go-xorm/xorm"
)

//Zone zone info
type Zone struct {
	ID       int    `xorm:"pk 'id' "`
	Name     string `xorm:"notnull unique 'name' "`
	Path     string `xorm:"varchar(255) notnull unique 'path'"`
	Parent   int    `xorm:"'parent' "`
	ImageURL string `xorm:" 'image_url'"`
	X        int    `xorm:"'x'"`
	Y        int    `xorm:"'y'"`
}

func (r *Zone) String() string {
	jdata, _ := json.Marshal(r)
	return string(jdata)
}

//GetZoneByPath get by path
func GetZoneByPath(path string, engin *xorm.Engine) (*Zone, error) {
	var zone Zone
	_, err := engin.Where("path=?", path).Get(&zone)
	return &zone, err
}

//GetZoneByNamePath get by path
func GetZoneByNamePath(name string, path string, engin *xorm.Engine) (*Zone, error) {
	var zone Zone
	_, err := engin.Where("name=? and path=?", name, path).Get(&zone)
	return &zone, err
}

//GetZoneByID get by id
func GetZoneByID(id int, engin *xorm.Engine) (*Zone, error) {
	var zone Zone
	zone.ID = id
	exist, err := engin.Get(&zone)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("zone (%d) not exist", id)
	}
	return &zone, err
}

//GetZonesByParent get by parent
func GetZonesByParent(parent int, engin *xorm.Engine) ([]*Zone, error) {
	var zone Zone
	var zones []*Zone

	rows, err := engin.Where("parent=?", parent).Rows(&zone)
	if err != nil {
		return zones, err
	}
	defer rows.Close()
	for rows.Next() {
		tmpZone := new(Zone)
		rows.Scan(tmpZone)

		zones = append(zones, tmpZone)

	}

	return zones, nil

}

//CreatePath create zone path
func CreatePath(path string, engin *xorm.Engine) (int, error) {

	var zone Zone
	ok, err := engin.Where("path=?", path).Get(&zone)
	if ok {
		seelog.Infof("path hash exist:%s", path)
		return zone.ID, nil
	}
	seelog.Infof("get zone by path %s error:%v", path, err)

	paths := strings.Split(path, ".")
	var parent int
	if len(paths) > 1 {
		parentPaths := paths[:len(paths)-1]
		parenntPath := strings.Join(parentPaths, ".")
		parent, err = CreatePath(parenntPath, engin)
		if err != nil {
			return 0, err
		}
	}
	newZone := &Zone{
		ID:     allocateid.AllocateID(),
		Name:   paths[len(paths)-1],
		Path:   path,
		Parent: parent,
		X:      -1,
		Y:      -1,
	}

	_, err = engin.Insert(newZone)
	if err != nil {
		seelog.Infof("Inset Zone %s erorr", newZone)
		return 0, err
	}
	return newZone.ID, nil
}

//UpdateZone 更新区域信息
func UpdateZone(z *Zone, engine *xorm.Engine) error {
	_, err := engine.Id(z.ID).Cols("image_url", "x", "y").Update(z)
	return err
}
