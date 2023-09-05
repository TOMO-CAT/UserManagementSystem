package mysqltable

import "gorm.io/gorm"

type UserAuth struct {
	gorm.Model
	Uid          uint64
	Username     string
	PasswordHash string
}
