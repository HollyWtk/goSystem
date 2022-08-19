package models

import (
	"time"
)

type TSysPermission struct {
	Id         int64     `xorm:"'id' pk autoincr BIGINT(20)"`
	Pid        int64     `xorm:"'pid' comment('父级权限id') BIGINT(20)"`
	Name       string    `xorm:"'name' comment('名称') VARCHAR(100)"`
	Value      string    `xorm:"'value' comment('权限值') VARCHAR(200)"`
	Icon       string    `xorm:"'icon' comment('图标') VARCHAR(500)"`
	Type       int       `xorm:"'type' comment('权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）') INT(11)"`
	Uri        string    `xorm:"'uri' comment('前端资源路径') VARCHAR(200)"`
	Status     int       `xorm:"'status' comment('启用状态；0->禁用；1->启用') INT(11)"`
	CreateTime time.Time `xorm:"'create_time' comment('创建时间') DATETIME"`
	Sort       int       `xorm:"'sort' comment('排序') INT(11)"`
	PageIndex  string    `xorm:"'page_index' VARCHAR(32)"`
}

func (m *TSysPermission) TableName() string {
	return "t_sys_permission"
}
