package logic

import (
	"server/core/config"
)

var Common *CommonLogic

type CommonLogic struct {
	adminExists bool
	register    bool
	init        bool // 是否初始化
}

func init() {
	Common = &CommonLogic{
		adminExists: false,
		init:        false,
	}
}

func (comm *CommonLogic) Init() error {
	if comm.init {
		return nil
	}
	defer func() {
		comm.init = true
	}()
	userLogic := new(UserLogic)
	exists, err := userLogic.adminExists()
	if err != nil {
		return err
	}
	comm.adminExists = exists
	return nil
}

func (comm *CommonLogic) SetAdminExists(exists bool) {
	comm.adminExists = exists
}

func (comm *CommonLogic) AdminExists() bool {
	return comm.adminExists
}

// Register 注册功能是否开放
func (comm *CommonLogic) Register() bool {
	if config.GetConfig().Server.Register {
		return true
	} else {
		return !comm.AdminExists()
	}
}
