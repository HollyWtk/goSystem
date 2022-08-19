package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"merchantService/constant"
	"merchantService/server/common"
	"merchantService/server/web/apply/logic"
	"merchantService/server/web/apply/model"
	"net/http"
)

var DefaultApplyController = &ApplyController{}

type ApplyController struct {
}

func (a *ApplyController) AddApply(ctx *gin.Context) {
	applyReq := &model.ApplyReq{}
	err := ctx.BindJSON(applyReq)
	if err != nil {
		log.Println("参数格式不合法", err)
		ctx.JSON(http.StatusOK, common.Error(constant.InvalidParam, "参数不合法"))
		return
	}
	validate := validator.New()
	err = validate.Struct(applyReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error(constant.InvalidParam, "参数错误"))
		return
	}
	err = logic.DefaultApplyLogic.DoApply(applyReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Error(constant.OK, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, common.Success(constant.OK, nil))
}

func (a *ApplyController) QueryApply(ctx *gin.Context) {
	merchantId := ctx.Query("merchantId")
	if len(merchantId) == 0 {
		ctx.JSON(http.StatusInternalServerError, common.Error(constant.InvalidParam, "参数错误"))
		return
	}
	apply, err := logic.DefaultApplyLogic.QueryApply(merchantId)
	if err != nil {
		ctx.JSON(http.StatusOK, common.Error(constant.SystemError, "系统错误,请联系管理员"))
		return
	}
	ctx.JSON(http.StatusOK, common.Success(constant.OK, apply))
}
