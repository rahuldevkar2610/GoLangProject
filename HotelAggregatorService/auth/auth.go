package auth

import "HotelAggregatorService/consts"

type Auth interface {
	Authenticate(string, string) (*AuthInfo, error)
}

type AuthInfo struct {
	User string
	Role string
}

func (ai AuthInfo) IsAdmin() bool {
	return ai.Role == consts.ROLEADMIN
}
