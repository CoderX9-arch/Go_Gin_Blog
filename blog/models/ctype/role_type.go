package ctype

import (
	"encoding/json"
)

type Role int

const (
	Permissionadmin       Role = 1
	PermissionUser        Role = 2
	PermissionVisiter     Role = 3
	PermissionDisableUser Role = 4
)

type Status int

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case Permissionadmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisiter:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁止的用户"
	default:
		str = "被禁止的用户"
	}
	return str
}
