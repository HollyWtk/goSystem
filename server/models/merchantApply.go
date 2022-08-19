package models

import (
	"time"
)

type MerchantApply struct {
	Id                    int       `xorm:"F_id pk autoincr"`
	MerchantId            string    `xorm:"F_merchant_id"`
	UserNo                string    `xorm:"F_user_no"`
	Level                 int       `xorm:"F_level"`
	RegisterTime          time.Time `xorm:"F_register_time"`
	RegisterMobile        string    `xorm:"F_register_mobile"`
	AgentId1g             string    `xorm:"F_agent_id_1g"`
	AgentIdSuper          string    `xorm:"F_agent_id_super"`
	AgentClass            int       `xorm:"F_agent_class"`
	Status                int       `xorm:"F_status default 0"`
	CheckTime             time.Time `xorm:"F_check_time"`
	CheckStatus           int       `xorm:"F_check_status"`
	MerchantCheckStatus   int       `xorm:"F_merchant_check_status"`
	ChannelStatus         int       `xorm:"F_channel_status"`
	BaseInfo              string    `xorm:"F_base_info"`
	SettleInfo            string    `xorm:"F_settle_info"`
	SettleOcrInfo         string    `xorm:"F_settle_ocr_info"`
	LicenseInfo           string    `xorm:"F_license_info"`
	LicenseOcr            string    `xorm:"F_license_ocr"`
	LicenseAuth           string    `xorm:"F_license_auth"`
	LicenseNumber         string    `xorm:"F_license_number"`
	LegalInfo             string    `xorm:"F_legal_Info"`
	LegalOcrFront         string    `xorm:"F_legal_ocr_front"`
	LegalOcrEnd           string    `xorm:"F_legal_ocr_end"`
	Shareholder1          string    `xorm:"F_shareholder_1"`
	Shareholder2          string    `xorm:"F_shareholder_2"`
	Shareholder3          string    `xorm:"F_shareholder_3"`
	Shareholder4          string    `xorm:"F_shareholder_4"`
	LegalIdcardNo         string    `xorm:"F_legal_idcard_no"`
	ShopInfo              string    `xorm:"F_shop_Infos"`
	ReceiveUse            int       `xorm:"F_receive_use"`
	Step                  int       `xorm:"F_step"`
	Creator               string    `xorm:"F_creator"`
	CreateTime            time.Time `xorm:"F_create_time"`
	Updater               string    `xorm:"F_updater"`
	UpdateTime            time.Time `xorm:"F_update_time"`
	ChannelAccount        string    `xorm:"F_channel_account"`
	ForeignReceiveAccount string    `xorm:"F_foreign_receive_account"`
	Comment               string    `xorm:"F_comment"`
	MerchantType          int       `xorm:"F_merchant_type"`
	AccountOpenFormImage  string    `xorm:"F_account_open_form_image"`
	HkLicenseInfo         string    `xorm:"F_hk_license_info"`
}

func (*MerchantApply) TableName() string {
	return "t_receive_payment_merchant_apply"
}
