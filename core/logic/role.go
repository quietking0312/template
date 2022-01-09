package logic

import (
	"go.uber.org/zap"
	"server/common/idprocess"
	"server/common/log"
	"server/core/dao"
)

type RoleLogic struct {
}

// 角色id 生成器
var ridProcess = idprocess.New(99)

func (r RoleLogic) GetRoleList(page, limit int) ([]dao.MRoleTable, error) {
	var dest []dao.MRoleTable
	roleModel := new(dao.RoleModel)
	if err := roleModel.SelectRoleList((page-1)*limit, limit, &dest); err != nil {
		log.Error("", zap.Error(err))
		return nil, err
	}
	return dest, nil
}

func (r RoleLogic) GetRoleTotal() (int, error) {
	var total int
	roleModel := new(dao.RoleModel)
	if err := roleModel.SelectRoleTotal(&total); err != nil {
		log.Error("", zap.Error(err))
		return 0, err
	}
	return total, nil
}

func (r RoleLogic) AddRole(roleName string) error {
	roleModel := new(dao.RoleModel)
	var roleTable = dao.MRoleTable{
		Rid:      int64(ridProcess.Id()),
		RoleName: roleName,
	}
	if err := roleModel.InsertOne(roleTable); err != nil {
		return err
	}
	return nil
}
