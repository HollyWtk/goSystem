package models

type SysAdminRoleRelation struct {
	Id      int64 `xorm:"'id' pk autoincr BIGINT(20)"`
	AdminId int64 `xorm:"'admin_id' BIGINT(20)"`
	RoleId  int64 `xorm:"'role_id' BIGINT(20)"`
}

func (m *SysAdminRoleRelation) TableName() string {
	return "t_sys_admin_role_relation"
}
