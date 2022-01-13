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

type RolePidItem struct {
	Rid           int64    `json:"rid"`
	RoleName      string   `json:"role_name"`
	PermissionIds []uint32 `json:"permission_ids"`
}

func (r RoleLogic) GetRoleList(page, limit int) ([]RolePidItem, error) {
	var dest []dao.MRoleTable
	roleModel := new(dao.RoleModel)
	if err := roleModel.SelectRoleList((page-1)*limit, limit, &dest); err != nil {
		log.Error("", zap.Error(err))
		return nil, err
	}
	var rolePidList []RolePidItem
	rolePermissionModel := new(dao.RolePermissionModel)
	for _, roleItem := range dest {
		var rolePid = RolePidItem{
			Rid:      roleItem.Rid,
			RoleName: roleItem.RoleName,
		}
		var pIds []uint32
		if err := rolePermissionModel.SelectListByRid(roleItem.Rid, &pIds); err != nil {
			return nil, err
		}
		rolePid.PermissionIds = pIds
		rolePidList = append(rolePidList, rolePid)
	}
	return rolePidList, nil
}

func (r RoleLogic) GetRoleAll() ([]dao.MRoleTable, error) {
	var dest []dao.MRoleTable
	roleModel := new(dao.RoleModel)
	if err := roleModel.SelectRoleList(0, 100, &dest); err != nil {
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

func (r RoleLogic) UpdateRole(rid int64, roleName string) error {
	roleModel := new(dao.RoleModel)
	var roleTable = dao.MRoleTable{
		Rid:      rid,
		RoleName: roleName,
	}
	if err := roleModel.UpdateRoleOne(roleTable); err != nil {
		return err
	}
	return nil
}

func (r RoleLogic) UpdatePermission(rid int64, pidS []uint32) error {
	rolePermissionModel := new(dao.RolePermissionModel)
	if err := rolePermissionModel.Insert(rid, pidS); err != nil {
		return err
	}
	return nil
}
