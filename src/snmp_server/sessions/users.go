package sessions

import (
	"fmt"
	"snmp_server/model"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

var userSessions sync.Map

//UserSession user for session
type UserSession struct {
	user     *model.User
	activeCh chan bool
	ttl      int
	token    string
	lastTime time.Time
}

func (r *UserSession) keepalive() {

	defer userSessions.Delete(r.token)
	timer := time.NewTimer(time.Duration(r.ttl) * time.Second)
	r.lastTime = time.Now()
	for {
		select {
		case active := <-r.activeCh:
			if active {
				r.lastTime = time.Now()
				timer.Reset(time.Duration(r.ttl) * time.Second)
			} else {
				model.UserLogoutLog(r.user.Username, fmt.Sprintf("last:%s, token:%s", r.lastTime.Format("2006-01-02 15:04:05"), r.token), false)
				return
			}

		case <-timer.C:
			model.UserLogoutLog(r.user.Username, fmt.Sprintf("last:%s, token:%s", r.lastTime.Format("2006-01-02 15:04:05"), r.token), true)
			return
		}
	}
}

//SetUserSessionKeepalive set keepalive
func SetUserSessionKeepalive(token string) {

	value, ok := userSessions.Load(token)
	if ok {
		us := value.(*UserSession)
		us.activeCh <- true

	}
}

//Logout 用户主动注销
func Logout(token string) {
	value, ok := userSessions.Load(token)
	if ok {
		us := value.(*UserSession)
		us.activeCh <- false

	}
}

//GetUserSession get session by token
func GetUserSession(token string) (*model.User, bool) {

	value, ok := userSessions.Load(token)
	if ok {
		return value.(*UserSession).user, true
	}
	return nil, false
}

//AllocateSession return token for user
func AllocateSession(user *model.User) string {
	id, _ := uuid.NewV4()

	u := &UserSession{
		user:     user,
		activeCh: make(chan bool, 1),
		token:    id.String(),
		ttl:      600,
	}
	userSessions.Store(id.String(), u)
	go u.keepalive()
	return id.String()

}
