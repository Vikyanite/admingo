package app

import "server/service/user/dao"

type User struct {
	dao.User
}

func (u *User) GetUser() {

}
