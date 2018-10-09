package model

import (
	"snmp_server/allocateid"

	"github.com/go-xorm/xorm"
)

//User user info
type User struct {
	ID       int    `xorm:"pk 'id' "`
	Username string `xorm:"varchar(255) notnull unique 'username'"`
	Password string `xorm:"'password'"`
	Type     int    `xorm:"'type'"`
	Parent   int    `xorm:"'parent'"`
}

//GetUsersByParent get users by parent
func GetUsersByParent(parent int, engine *xorm.Engine) ([]*User, error) {
	var user User
	rows, err := engine.Where("parent=?", parent).Rows(user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		tmp := new(User)
		rows.Scan(tmp)
		users = append(users, tmp)
	}
	return users, nil
}

//GetUserByID get user by id
func GetUserByID(id int, engine *xorm.Engine) (*User, error) {

	user := new(User)
	ok, err := engine.Where("id=?", id).Get(user)
	if ok {
		return user, nil
	}
	return nil, err
}

//GetUserByName get user by name
func GetUserByName(name string, engine *xorm.Engine) (*User, error) {

	user := new(User)
	ok, err := engine.Where("username=?", name).Get(user)
	if ok {
		return user, nil
	}
	return nil, err
}

//CreateUser create user
func CreateUser(u *User, engine *xorm.Engine) (int, error) {

	u.ID = allocateid.AllocateID()
	_, err := engine.InsertOne(u)
	return u.ID, err
}

//UpdateUser update
func UpdateUser(u *User, engine *xorm.Engine) error {
	_, err := engine.Id(u.ID).Update(u)
	return err
}

//RemoteUser remote user
func RemoteUser(id int, engine *xorm.Engine) error {
	var user User
	user.ID = id

	_, err := engine.Delete(&user)
	return err
}
