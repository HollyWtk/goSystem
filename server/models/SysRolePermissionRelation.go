package models

type TSysRolePermissionRelation struct {
	Id           int64 `xorm:"'id' pk autoincr BIGINT(20)"`
	RoleId       int64 `xorm:"'role_id' BIGINT(20)"`
	PermissionId int64 `xorm:"'permission_id' BIGINT(20)"`
}

func (m *TSysRolePermissionRelation) TableName() string {
	return "t_sys_role_permission_relation"
}
