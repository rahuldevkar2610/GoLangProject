package auth

import (
	"HotelAggregatorService/consts"
	"errors"
)

type user struct {
	user     string
	password string
	role     string
}
type AuthSrvMemory struct {
	userdb map[string]user
}

func NewAuthServiceMemory() AuthSrvMemory {

	users := make(map[string]user)

	users["abhijit"] = user{"abhijit", "abhijit123", consts.ROLEADMIN}
	users["amit"] = user{"amit", "amit@123", consts.ROLEUSER}
	users["somnath"] = user{"somnath", "somnath123", consts.ROLEUSER}
	users["sheetal"] = user{"sheetal", "sheetal123", consts.ROLEADMIN}

	return AuthSrvMemory{
		userdb: users,
	}
}

func (authsrv *AuthSrvMemory) Authenticate(user, password string) (*AuthInfo, error) {
	userrec, ok := authsrv.userdb[user]

	if !ok {
		return nil, errors.New("user not found")
	}

	if userrec.user == user && userrec.password == password {
		return &AuthInfo{User: userrec.user, Role: userrec.role}, nil
	} else {
		return nil, errors.New("unauthorised")
	}
}
