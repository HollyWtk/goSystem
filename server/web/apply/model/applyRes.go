package model

type ApplyRes struct {
	Id                    int              `json:"id"`
	AccountOpenFormImage  string           `json:"accountOpenFormImage"`
	AgentClass            int              `json:"agentClass"`
	AgentId1G             string           `json:"agentId1g"`
	AgentIdSuper          string           `json:"agentIdSuper"`
	ChannelAccount        string           `json:"channelAccount"`
	ChannelStatus         int              `json:"channelStatus"`
	MerchantId            string           `json:"merchantId"`
	CheckStatus           int              `json:"checkStatus"`
	CheckTime             string           `json:"checkTime"`
	Comment               string           `json:"comment"`
	CreateTime            string           `json:"createTime"`
	Creator               string           `json:"creator"`
	DefaultComment        string           `json:"defaultComment"`
	ForeignReceiveAccount string           `json:"foreignReceiveAccount"`
	MerchantType          int              `json:"merchantType"`
	ReceiveUse            int              `json:"receiveUse"`
	Step                  int              `json:"step"`
	Level                 int              `json:"level"`
	MerchantCheckStatus   int              `json:"merchantCheckStatus"`
	RegisterMobile        string           `json:"registerMobile"`
	RegisterTime          string           `json:"registerTime"`
	Status                int              `json:"status"`
	UpdateTime            string           `json:"updateTime"`
	Updater               string           `json:"updater"`
	BaseInfo              *BaseInfo        `json:"baseInfo"`
	HkLicenseInfo         *HkLicenseInfo   `json:"hkLicenseInfo"`
	LegalInfo             *LegalInfo       `json:"legalInfo"`
	LicenseInfo           *LicenseInfo     `json:"licenseInfo"`
	SettleInfo            *SettleInfo      `json:"settleInfo"`
	Shareholder1          *ShareholderInfo `json:"shareholder1"`
	Shareholder2          *ShareholderInfo `json:"shareholder2"`
	Shareholder3          *ShareholderInfo `json:"shareholder3"`
	Shareholder4          *ShareholderInfo `json:"shareholder4"`
	ShopInfo              *ShopInfo        `json:"shopInfo"`
	LegalOcrEnd           *LegalOcrEnd     `json:"LegalOcrEnd"`
	LegalOcrFront         *LegalOcrFront   `json:"legalOcrFront"`
	LicenseOcr            *LicenseOcr      `json:"licenseOcr"`
	Records               *[]Record        `json:"records"`
	SettleOcrInfo         *SettleOcrInfo   `json:"settleOcrInfo"`
}
type LegalOcrEnd struct {
	Authority   string `json:"authority"`
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	ValidDate   string `json:"validDate"`
	WarningCode string `json:"warningCode"`
	WarningMsg  string `json:"warningMsg"`
}
type LegalOcrFront struct {
	Address     string `json:"address"`
	Birth       string `json:"birth"`
	Code        string `json:"code"`
	Idcard      string `json:"idcard"`
	Msg         string `json:"msg"`
	Name        string `json:"name"`
	Nation      string `json:"nation"`
	Sex         string `json:"sex"`
	WarningCode string `json:"warningCode"`
	WarningMsg  string `json:"warningMsg"`
}

type LicenseOcr struct {
	Address   string `json:"address"`
	Capital   string `json:"capital"`
	Date      string `json:"date"`
	Issue     string `json:"issue"`
	LegalName string `json:"legalName"`
	LicenseNo string `json:"licenseNo"`
	Name      string `json:"name"`
	Period    string `json:"period"`
	Type      string `json:"type"`
}
type Record struct {
	CheckOperator string `json:"checkOperator"`
	CheckTime     string `json:"checkTime"`
	Comment       string `json:"comment"`
	Id            int    `json:"id"`
	MerchantId    string `json:"merchantId"`
	Result        int    `json:"result"`
	Type          int    `json:"type"`
}
type SettleOcrInfo struct {
	BankcardNo        string `json:"bankcardNo"`
	BankcardValidDate string `json:"bankcardValidDate"`
}
