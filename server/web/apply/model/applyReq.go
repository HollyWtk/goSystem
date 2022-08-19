package model

type ApplyReq struct {
	MerchantId       string `validate:"required"`
	BaseInfo         *BaseInfo
	HkLicenseInfo    *HkLicenseInfo
	LegalInfo        *LegalInfo
	LicenseInfo      *LicenseInfo
	MerchantType     int
	ReceiveUse       int
	SettleInfo       *SettleInfo
	ShareholderInfos *[]ShareholderInfo
	ShopInfo         *ShopInfo
	Step             int
}

type BaseInfo struct {
	Address         string `json:"address"`
	BusinessScale   int    `json:"businessScale"`
	BusinessScope   string `json:"businessScope"`
	BusinessVolume  int    `json:"businessVolume"`
	BusinessWeb     string `json:"businessWeb"`
	Category        string `json:"category"`
	ContractEmail   string `json:"contractEmail"`
	ContractNumber  string `json:"contractNumber"`
	Industry        string `json:"industry"`
	Name            string `json:"name"`
	NameAbbr        string `json:"nameAbbr"`
	NameEn          string `json:"nameEn"`
	OfficeEnvImage  string `json:"officeEnvImage"`
	ReceptionImage  string `json:"receptionImage"`
	Region          string `json:"region"`
	StorehouseImage string `json:"storehouseImage"`
	StructureImage  string `json:"structureImage"`
}

type HkLicenseInfo struct {
	BusinessBrImage     string `json:"businessBrImage"`
	CeoHolderImage      string `json:"ceoHolderImage"`
	CompanyAddress      string `json:"companyAddress"`
	CompanyBusinessCode string `json:"companyBusinessCode"`
	CompanyCnName       string `json:"companyCnName"`
	CompanyEnName       string `json:"companyEnName"`
	CompanyRegisterCode string `json:"companyRegisterCode"`
	LegalsTableImage    string `json:"legalsTableImage"`
	RegisterCrImage     string `json:"registerCrImage"`
}

type LegalInfo struct {
	ExpiryDate  string  `json:"expiryDate"`
	IdCardEnd   string  `json:"idCardEnd"`
	IdCardFront string  `json:"idCardFront"`
	IdCardNo    string  `json:"idCardNo"`
	IdType      string  `json:"idType"`
	IssueDate   string  `json:"issueDate"`
	Mobile      string  `json:"mobile"`
	Name        string  `json:"name"`
	ReferenceNo string  `json:"referenceNo"`
	Region      string  `json:"region"`
	Shares      float32 `json:"shares"`
}
type LicenseInfo struct {
	Address          string `json:"address"`
	BusinessRange    string `json:"businessRange"`
	ExpiryDate       string `json:"expiryDate"`
	ImageUrl         string `json:"imageUrl"`
	IssueDate        string `json:"issueDate"`
	LegalHolderImage string `json:"legalHolderImage"`
	LegalName        string `json:"legalName"`
	LicenseName      string `json:"licenseName"`
	LicenseNumber    string `json:"licenseNumber"`
}
type SettleInfo struct {
	AccountMobile string `json:"accountMobile"`
	AccountNo     string `json:"accountNo"`
	BankArea      string `json:"bankArea"`
	BankCity      string `json:"bankCity"`
	BankCode      string `json:"bankCode"`
	BankOutlets   string `json:"bankOutlets"`
	BankProvince  string `json:"bankProvince"`
	CardImage     string `json:"cardImage"`
	Cardholder    string `json:"cardholder"`
	SettleType    int    `json:"settleType"`
}

type ShareholderInfo struct {
	IdCardEnd   string  `json:"idCardEnd"`
	IdCardFront string  `json:"idCardFront"`
	IdCardNo    string  `json:"idCardNo"`
	IdType      string  `json:"idType"`
	Mobile      string  `json:"mobile"`
	Name        string  `json:"name"`
	ReferenceNo string  `json:"referenceNo"`
	Region      string  `json:"region"`
	Shares      float32 `json:"shares"`
	Type        string  `json:"type"`
}

type ShopInfo struct {
	ReferenceNo     string `json:"referenceNo"`
	ShopAccount     string `json:"shopAccount"`
	ShopAccountName string `json:"shopAccountName"`
	ShopCategoryOne int    `json:"shopCategoryOne"`
	ShopCategoryTwo int    `json:"shopCategoryTwo"`
	ShopId          string `json:"shopId"`
	ShopName        string `json:"shopName"`
	ShopRegion      string `json:"shopRegion"`
	ShopType        int    `json:"shopType"`
	ShopWebUrl      string `json:"shopWebUrl"`
}
