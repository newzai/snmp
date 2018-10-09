package model

import (
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func Test_Zone(t *testing.T) {

	engin, err := xorm.NewEngine("sqlite3", "./model.db")
	if err != nil {
		t.Errorf("NewEngine:%s", err)
		return
	}

	err = engin.Sync2(new(Zone))
	if err != nil {
		t.Errorf("Sync2:%s", err)
		return
	}

	_, err = CreatePath("root.nanshang.kejiy", engin)
	if err != nil {
		t.Errorf("CreatePath :%s", err)
	}

	_, err = CreatePath("root.futian.kejiy", engin)
	if err != nil {
		t.Errorf("CreatePath :%s", err)
	}

	_, err = CreatePath("root.luohu.kejiy", engin)
	if err != nil {
		t.Errorf("CreatePath :%s", err)
	}

	zone, err := GetZoneByID(1, engin)
	if err != nil {
		t.Errorf("GetZone by id error %s", err)
	} else if zone.ID != 1 {
		t.Error("zone id error")
	} else if zone.Name != "root" {
		t.Error("zone name error")
	}

	zones, err := GetZonesByParent(1, engin)
	t.Error(zones, err)
}
