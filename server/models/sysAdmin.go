package models

import (
	"time"
)

type SysAdmin struct {
	Id         int64     `xorm:"'id' pk autoincr BIGINT(20)"`
	Username   string    `xorm:"'username' unique VARCHAR(64)"`
	Password   string    `xorm:"'password' VARCHAR(64)"`
	Icon       string    `xorm:"'icon' default '' VARCHAR(255)"`
	Mobile     string    `xorm:"'mobile' default '' comment('联系人号码') VARCHAR(11)"`
	Email      string    `xorm:"'email' default '' comment('邮箱') VARCHAR(100)"`
	NickName   string    `xorm:"'nick_name' default '' comment('昵称') VARCHAR(200)"`
	Note       string    `xorm:"'note' default '' comment('备注信息') VARCHAR(500)"`
	LoginTime  time.Time `xorm:"'login_time' default '1970-01-01 00:00:00' comment('最后登录时间') DATETIME"`
	Status     int       `xorm:"'status' default 1 comment('帐号启用状态：0->禁用；1->启用') INT(11)"`
	Creator    string    `xorm:"'creator' comment('创建人') VARCHAR(64)"`
	CreateTime time.Time `xorm:"'create_time' not null comment('创建时间') DATETIME"`
}

func (m *SysAdmin) TableName() string {
	return "t_sys_admin"
}
