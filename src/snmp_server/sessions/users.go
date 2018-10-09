package sessions

import (
	"snmp_server/model"
	"sync"

	uuid "github.com/satori/go.uuid"
)

var userSessions sync.Map

//GetUserSession get session by token
func GetUserSession(token string) (*model.User, bool) {

	value, ok := userSessions.Load(token)
	if ok {
		return value.(*model.User), true
	}
	return nil, false
}

//AllocateSession return token for user
func AllocateSession(user *model.User) string {
	id, _ := uuid.NewV4()

	userSessions.Store(id.String(), user)
	return id.String()

}
