package models

import (
	"time"
)

type TSysRole struct {
	Id          int64     `xorm:"'id' pk autoincr BIGINT(20)"`
	Name        string    `xorm:"'name' comment('名称') VARCHAR(100)"`
	Description string    `xorm:"'description' comment('描述') VARCHAR(500)"`
	AdminCount  int       `xorm:"'admin_count' comment('后台用户数量') INT(11)"`
	CreateTime  time.Time `xorm:"'create_time' comment('创建时间') DATETIME"`
	Creator     string    `xorm:"'creator' comment('创建人') VARCHAR(255)"`
	Status      int       `xorm:"'status' default 1 comment('启用状态：0->禁用；1->启用') INT(11)"`
	Sort        int       `xorm:"'sort' default 0 INT(11)"`
}

func (m *TSysRole) TableName() string {
	return "t_sys_role"
}
