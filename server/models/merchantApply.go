package models

import (
	"time"
)

type MerchantApply struct {
	Id                    int64     `xorm:"'F_id' not null pk autoincr BIGINT(20)"`
	MerchantId            string    `xorm:"'F_merchant_id' not null comment('商户编号') index unique VARCHAR(20)"`
	UserNo                string    `xorm:"'F_user_no' not null comment('商户对应用户编号') VARCHAR(128)"`
	Level                 int       `xorm:"'F_level' default 0 comment('商户等级') TINYINT(4)"`
	RegisterTime          time.Time `xorm:"'F_register_time' not null comment('注册时间') DATETIME"`
	RegisterMobile        string    `xorm:"'F_register_mobile' not null comment('注册手机号') VARCHAR(32)"`
	AgentId1g             string    `xorm:"'F_agent_id_1g' default '' comment('一级代理商编号') index VARCHAR(32)"`
	AgentIdSuper          string    `xorm:"'F_agent_id_super' default '' comment('直属代理商编号') VARCHAR(32)"`
	AgentClass            int       `xorm:"'F_agent_class' default -1 comment('代理商类型') TINYINT(4)"`
	Status                int       `xorm:"'F_status' default 1 comment('商户状态 1| 正常 2|注销') TINYINT(4)"`
	CheckStatus           int       `xorm:"'F_check_status' default 1 comment('1|待审核 2| 审核通过 3|审核不通过') TINYINT(4)"`
	ChannelStatus         string    `xorm:"'F_channel_status' default '' comment('通道审核状态') VARCHAR(64)"`
	MerchantCheckStatus   int       `xorm:"'F_merchant_check_status' default -1 comment('商户审核状态') TINYINT(4)"`
	CheckTime             time.Time `xorm:"'F_check_time' comment('审核通过时间') index DATETIME"`
	BaseInfo              string    `xorm:"'F_base_info' default '' comment('企业基础信息') VARCHAR(1000)"`
	SettleInfo            string    `xorm:"'F_settle_info' default '' comment('结算信息') VARCHAR(1000)"`
	LicenseInfo           string    `xorm:"'F_license_info' default '' comment('营业执照信息') VARCHAR(1000)"`
	LicenseNumber         string    `xorm:"'F_license_number' default '' comment('营业执照编号') VARCHAR(128)"`
	LegalInfo             string    `xorm:"'F_legal_info' default '' comment('法人信息') VARCHAR(1000)"`
	LegalIdcardNo         string    `xorm:"'F_legal_idcard_no' default '' comment('法人身份证') VARCHAR(128)"`
	Shareholders          string    `xorm:"'F_shareholders' default '' comment('股东信息') VARCHAR(1000)"`
	ShopInfos             string    `xorm:"'F_shop_Infos' default '' comment('店铺信息') VARCHAR(1000)"`
	ReceiveUse            int       `xorm:"'F_receive_use' default -1 comment('收款用途') TINYINT(4)"`
	Step                  int       `xorm:"'F_step' default 1 comment('提交开户资料步骤') TINYINT(4)"`
	Creator               string    `xorm:"'F_creator' not null comment('创建人') VARCHAR(64)"`
	CreateTime            time.Time `xorm:"'F_create_time' not null comment('创建时间') DATETIME"`
	Updater               string    `xorm:"'F_updater' default '' comment('最后更新人') VARCHAR(64)"`
	UpdateTime            time.Time `xorm:"'F_update_time' default '1970-01-01 00:00:00' comment('最后更新时间') DATETIME"`
	ChannelAccount        string    `xorm:"'F_channel_account' default '' comment('hkpu账户ID') VARCHAR(128)"`
	ForeignReceiveAccount string    `xorm:"'F_foreign_receive_account' default '' comment('国外收款账户') VARCHAR(128)"`
	Comment               string    `xorm:"'F_comment' default '' comment('审核备注') VARCHAR(255)"`
	Shareholder1          string    `xorm:"'F_shareholder_1' default '' comment('股东信息1') VARCHAR(1000)"`
	Shareholder2          string    `xorm:"'F_shareholder_2' default '' comment('股东信息2') VARCHAR(1000)"`
	Shareholder3          string    `xorm:"'F_shareholder_3' default '' comment('股东信息3') VARCHAR(1000)"`
	Shareholder4          string    `xorm:"'F_shareholder_4' default '' comment('股东信息4') VARCHAR(1000)"`
	LicenseOcr            string    `xorm:"'F_license_ocr' default '' comment('营业执照OCR信息') VARCHAR(1000)"`
	LicenseAuth           string    `xorm:"'F_license_auth' comment('营业执照工商鉴权信息') VARCHAR(1000)"`
	LegalOcrFront         string    `xorm:"'F_legal_ocr_front' default '' comment('法人身份证正面OCR') VARCHAR(255)"`
	LegalOcrEnd           string    `xorm:"'F_legal_ocr_end' default '' comment('法人身份证反面OCR') VARCHAR(255)"`
	SettleOcrInfo         string    `xorm:"'F_settle_ocr_info' default '' comment('结算卡OCR信息') VARCHAR(255)"`
	MerchantType          int       `xorm:"'F_merchant_type' not null default 1 comment('1.个人 | 2 个体 | 3 香港企业 | 4 中国企业') INT(11)"`
	AccountOpenFormImage  string    `xorm:"'F_account_open_form_image' default '' comment('开户申请表(个人商户不需要)') VARCHAR(255)"`
	HkLicenseInfo         string    `xorm:"'F_hk_license_info' default '' comment('香港企业营业制造信息') VARCHAR(1000)"`
}

func (m *MerchantApply) TableName() string {
	return "t_receive_payment_merchant_apply"
}
