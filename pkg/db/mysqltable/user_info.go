package mysqltable

import "gorm.io/gorm"

const (
	TableNameUserInfo = "user_infos"
)

type UserInfo struct {
	gorm.Model
	Uid       uint64
	Username  string
	Nickname  string
	AvatarUrl string
}
