package logic

import (
	"encoding/json"
	"log"
	"merchantService/constant"
	"merchantService/db"
	"merchantService/server/common"
	"merchantService/server/models"
	model2 "merchantService/server/web/apply/model"
	"merchantService/utils"
	"time"
)

var DefaultApplyLogic = &ApplyLogic{}

type ApplyLogic struct {
}

func (a *ApplyLogic) DoApply(req *model2.ApplyReq) error {
	err, done := doStep(req)
	if done {
		return err
	}
	return nil
}

func doStep(req *model2.ApplyReq) (error, bool) {
	step := req.Step
	var err error
	switch step {
	case 1:
		err = firstStep(req)
		if err != nil {
			return err, true
		}
		break
	case 2:
		secondStep(req)
		break
	case 3:
		thirdStep(req)
		break
	default:
		return common.New(constant.InvalidParam, "参数错误"), true
	}
	return nil, false
}

func firstStep(req *model2.ApplyReq) error {
	merchantId := req.MerchantId
	merchantApply := &models.MerchantApply{}
	ok, err := db.Engine.Table(merchantApply).Where("F_merchant_id=?", merchantId).Get(merchantApply)
	if err != nil {
		log.Println("商户申请信息查询失败", err)
		return common.New(constant.DBError, "数据库异常")
	}
	if ok {

	} else {
		log.Println("商户申请记录不存在,生成新的申请记录")
		err2, done := saveNewApply(req, merchantId, merchantApply)
		if done {
			return err2
		}
	}
	return nil
}

func saveNewApply(req *model2.ApplyReq, merchantId string, merchantApply *models.MerchantApply) (error, bool) {
	userInfo := &models.UserInfo{}
	ok, err := db.Engine.Table(userInfo).Where("F_merchant_id=?", merchantId).Get(userInfo)
	if err != nil {
		log.Println("用户信息查询失败", err)
		return common.New(constant.DBError, "数据库异常"), true
	}
	if !ok {
		log.Println("用户不存在", err)
		return common.New(constant.DBError, "用户不存在"), true
	}
	merchantApply = &models.MerchantApply{
		MerchantId:          merchantId,
		ChannelStatus:       string(rune(constant.AUTH_WAIT)),
		CheckStatus:         constant.WAIT_COMPLETE,
		MerchantCheckStatus: constant.WAIT_COMPLETE,
		AgentIdSuper:        userInfo.ChannelAgentNo,
		UserNo:              userInfo.UserNo,
		RegisterTime:        userInfo.CreateTime,
		RegisterMobile:      userInfo.Tel,
		Creator:             merchantId,
		Updater:             merchantId,
		CreateTime:          time.Now(),
		UpdateTime:          time.Now(),
		Status:              1,
	}
	handlerMerchantType(req, merchantApply)
	_, err = db.Engine.Insert(merchantApply)
	if err != nil {
		log.Println("商户申请记录插入失败", err)
		return common.New(constant.DBError, "数据库异常"), true
	}
	return nil, false
}

func handlerMerchantType(req *model2.ApplyReq, apply *models.MerchantApply) {
	merchantType := req.MerchantType
	baseInfo := req.BaseInfo
	legalInfo := req.LegalInfo
	hkLicenseInfo := req.HkLicenseInfo
	licenseInfo := req.LicenseInfo
	switch merchantType {
	//个人商户
	case 1:
		personBaseInfo := &model2.BaseInfo{
			ContractNumber: apply.RegisterMobile,
			Address:        baseInfo.Address,
			ContractEmail:  baseInfo.ContractEmail,
			Industry:       baseInfo.Industry,
			Name:           "商户_" + legalInfo.Name,
			NameAbbr:       "商户_" + legalInfo.Name,
		}
		baseInfoStr, _ := json.Marshal(personBaseInfo)
		apply.BaseInfo = string(baseInfoStr)
		legalInfoStr, _ := json.Marshal(legalInfo)
		apply.LegalInfo = string(legalInfoStr)
		apply.LegalIdcardNo = legalInfo.IdCardNo
		break
		//个体,中国企业
	case 2, 4:
		baseInfo.Name = licenseInfo.LicenseName
		baseInfo.NameAbbr = licenseInfo.LicenseName
		licenseStr, _ := json.Marshal(licenseInfo)
		apply.LicenseInfo = string(licenseStr)
		apply.LicenseNumber = licenseInfo.LicenseNumber
		baseInfoStr, _ := json.Marshal(baseInfo)
		apply.BaseInfo = string(baseInfoStr)
		break
		//香港企业
	case 3:
		hkLicenseInfoStr, _ := json.Marshal(hkLicenseInfo)
		apply.HkLicenseInfo = string(hkLicenseInfoStr)
		apply.LicenseNumber = hkLicenseInfo.CompanyRegisterCode
		baseInfo.NameEn = hkLicenseInfo.CompanyEnName
		baseInfo.NameAbbr = baseInfo.Name
		baseInfo.Name = hkLicenseInfo.CompanyCnName
		baseInfoStr, _ := json.Marshal(baseInfo)
		apply.BaseInfo = string(baseInfoStr)
		break
	}
	println(apply)

}

func secondStep(req *model2.ApplyReq) {

}

func thirdStep(req *model2.ApplyReq) {

}

func (a *ApplyLogic) QueryApply(merchantId string) (*model2.ApplyRes, error) {
	merchantApply := &models.MerchantApply{}
	applyRes := &model2.ApplyRes{}
	ok, err := db.Engine.Table(merchantApply).Where("F_merchant_id=?", merchantId).Get(merchantApply)
	if err != nil {
		log.Println("商户申请信息查询失败", err)
		return nil, common.New(constant.DBError, "数据库异常")
	}
	if ok {
		utils.StructCopy(applyRes, merchantApply)
		convert(merchantApply, applyRes)
		return applyRes, nil
	}
	return nil, nil
}

func convert(merchantApply *models.MerchantApply, applyRes *model2.ApplyRes) {
	baseInfo := &model2.BaseInfo{}
	hkLicenseInfo := &model2.HkLicenseInfo{}
	legalInfo := &model2.LegalInfo{}
	settleInfo := &model2.SettleInfo{}
	shareholder1 := &model2.ShareholderInfo{}
	shareholder2 := &model2.ShareholderInfo{}
	shareholder3 := &model2.ShareholderInfo{}
	shareholder4 := &model2.ShareholderInfo{}
	shopInfo := &model2.ShopInfo{}
	legalOcrFront := &model2.LegalOcrFront{}
	legalOcrEnd := &model2.LegalOcrEnd{}
	licenseOcr := &model2.LicenseOcr{}
	//records := &[]model.Record{}
	settleOcrInfo := &model2.SettleOcrInfo{}
	if merchantApply.BaseInfo != "" {
		_ = json.Unmarshal([]byte(merchantApply.BaseInfo), baseInfo)
		applyRes.BaseInfo = baseInfo
	}
	if merchantApply.HkLicenseInfo != "" {
		_ = json.Unmarshal([]byte(merchantApply.HkLicenseInfo), hkLicenseInfo)
		applyRes.HkLicenseInfo = hkLicenseInfo
	}
	if merchantApply.LegalInfo != "" {
		_ = json.Unmarshal([]byte(merchantApply.LegalInfo), legalInfo)
		applyRes.LegalInfo = legalInfo
	}
	if merchantApply.SettleInfo != "" {
		_ = json.Unmarshal([]byte(merchantApply.SettleInfo), settleInfo)
		applyRes.SettleInfo = settleInfo
	}
	if merchantApply.Shareholder1 != "" {
		_ = json.Unmarshal([]byte(merchantApply.Shareholder1), shareholder1)
		applyRes.Shareholder1 = shareholder1
	}
	if merchantApply.Shareholder2 != "" {
		_ = json.Unmarshal([]byte(merchantApply.Shareholder2), shareholder2)
		applyRes.Shareholder2 = shareholder2
	}
	if merchantApply.Shareholder3 != "" {
		_ = json.Unmarshal([]byte(merchantApply.Shareholder3), shareholder3)
		applyRes.Shareholder3 = shareholder3
	}
	if merchantApply.Shareholder4 != "" {
		_ = json.Unmarshal([]byte(merchantApply.Shareholder4), shareholder4)
		applyRes.Shareholder4 = shareholder4
	}
	if merchantApply.ShopInfos != "" {
		_ = json.Unmarshal([]byte(merchantApply.ShopInfos), shopInfo)
		applyRes.ShopInfo = shopInfo
	}
	if merchantApply.LegalOcrFront != "" {
		_ = json.Unmarshal([]byte(merchantApply.LegalOcrFront), legalOcrFront)
		applyRes.LegalOcrFront = legalOcrFront
	}
	if merchantApply.LegalOcrEnd != "" {
		_ = json.Unmarshal([]byte(merchantApply.LegalOcrEnd), legalOcrEnd)
		applyRes.LegalOcrEnd = legalOcrEnd
	}
	if merchantApply.LicenseOcr != "" {
		_ = json.Unmarshal([]byte(merchantApply.LicenseOcr), licenseOcr)
		applyRes.LicenseOcr = licenseOcr
	}
	if merchantApply.SettleOcrInfo != "" {
		_ = json.Unmarshal([]byte(merchantApply.SettleOcrInfo), settleOcrInfo)
		applyRes.SettleOcrInfo = settleOcrInfo
	}

}
