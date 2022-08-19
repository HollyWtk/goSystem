package models

import "time"

type UserInfo struct {
	Id             int16     `xorm:"F_id pk autoincr"`
	UserNo         string    `xorm:"F_user_no"`
	Tel            string    `xorm:"F_tel"`
	Email          string    `xorm:"F_email"`
	Address        string    `xorm:"F_address"`
	State          string    `xorm:"F_state"`
	CreateTime     time.Time `xorm:"F_create_time"`
	CreateBy       string    `xorm:"F_create_by"`
	UpdateTime     time.Time `xorm:"F_update_time"`
	UpdateBy       string    `xorm:"F_update_by"`
	MerchantId     string    `xorm:"F_merchant_id"`
	ChannelAgentNo string    `xorm:"F_channel_agent_no"`
	UserType       int       `xorm:"F_user_type"`
	AgentId        string    `xorm:"F_agent_id"`
}

func (*UserInfo) TableName() string {
	return "t_user_info"
}
