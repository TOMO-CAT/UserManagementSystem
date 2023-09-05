package mysqltable

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Uid       uint64
	Username  string
	Nickname  string
	AvatarUrl string
}
