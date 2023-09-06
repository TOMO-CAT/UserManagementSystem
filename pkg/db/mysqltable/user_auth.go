package mysqltable

import "gorm.io/gorm"

const (
	TableNameUserAuth = "user_auths"
)

type UserAuth struct {
	gorm.Model
	Uid          uint64
	Username     string
	PasswordHash string
}
