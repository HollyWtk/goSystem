package models

type SysAdminPermissionRelation struct {
	Id           int64 `xorm:"'id' pk autoincr BIGINT(20)"`
	AdminId      int64 `xorm:"'admin_id' BIGINT(20)"`
	PermissionId int64 `xorm:"'permission_id' BIGINT(20)"`
	Type         int   `xorm:"'type' INT(11)"`
}

func (m *SysAdminPermissionRelation) TableName() string {
	return "t_sys_admin_permission_relation"
}
