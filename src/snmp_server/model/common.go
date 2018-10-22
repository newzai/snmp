package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/go-xorm/xorm"

	"github.com/cihub/seelog"
)

//ItemType item type
type ItemType int

const (
	//ZONE is zone
	ZONE ItemType = 1
	//TERMINAL is terminal
	TERMINAL ItemType = 2
	//USER is user
	USER ItemType = 3
)

//InitDatabase db init
func InitDatabase(engine *xorm.Engine) {

	engine.Sync2(new(User), new(Zone), new(Terminal))
	rootid, _ := CreatePath("root", engine)
	admin, err := GetUserByName("admin", engine)
	if err != nil {
		seelog.Warnf("get admin user err:%s", err)
		return
	}
	if admin == nil {
		admin = new(User)
		admin.Username = "admin"
		h := md5.New()
		h.Write([]byte("123456"))
		admin.Password = hex.EncodeToString(h.Sum(nil))
		admin.Type = 1
		admin.Parent = rootid
		CreateUser(admin, engine)
	} else {

	}

}

//MD5 MD5
func MD5(raw string) string {
	h := md5.New()
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}
