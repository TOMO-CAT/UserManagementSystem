package config

import (
	"github.com/BurntSushi/toml"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util"
	"github.com/TOMO-CAT/UserManagementSystem/pkg/util/logger"
)

var (
	GlobalUmsConfig *umsConfig = &umsConfig{}
)

func ParseConfig(configPath string) error {
	if _, err := toml.DecodeFile(configPath, GlobalUmsConfig); err != nil {
		logger.Error("parse ums config [%s] fail with err [%v]", configPath, err)
		return err
	}

	logger.Info("init ums config by [%s] successfully!", configPath)
	logger.Info("ums config [%s]", util.ToString(GlobalUmsConfig))
	return nil
}
