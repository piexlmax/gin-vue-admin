package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ExaSensitiveWord
// @Summary 创建ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaSensitiveWord true "创建ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaSensitiveWord/createExaSensitiveWord [post]
func CreateExaSensitiveWord(c *gin.Context) {
	var exaSensitiveWord model.ExaSensitiveWord
	_ = c.ShouldBindJSON(&exaSensitiveWord)
	if err := service.CreateExaSensitiveWord(exaSensitiveWord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		middleware.Library.Build()
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ExaSensitiveWord
// @Summary 删除ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaSensitiveWord true "删除ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaSensitiveWord/deleteExaSensitiveWord [delete]
func DeleteExaSensitiveWord(c *gin.Context) {
	var exaSensitiveWord model.ExaSensitiveWord
	_ = c.ShouldBindJSON(&exaSensitiveWord)
	if err := service.DeleteExaSensitiveWord(exaSensitiveWord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		middleware.Library.Build()
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ExaSensitiveWord
// @Summary 批量删除ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exaSensitiveWord/deleteExaSensitiveWordByIds [delete]
func DeleteExaSensitiveWordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteExaSensitiveWordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		middleware.Library.Build()
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ExaSensitiveWord
// @Summary 更新ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaSensitiveWord true "更新ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exaSensitiveWord/updateExaSensitiveWord [put]
func UpdateExaSensitiveWord(c *gin.Context) {
	var exaSensitiveWord model.ExaSensitiveWord
	_ = c.ShouldBindJSON(&exaSensitiveWord)
	if err := service.UpdateExaSensitiveWord(exaSensitiveWord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		middleware.Library.Build()
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ExaSensitiveWord
// @Summary 用id查询ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaSensitiveWord true "用id查询ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exaSensitiveWord/findExaSensitiveWord [get]
func FindExaSensitiveWord(c *gin.Context) {
	var exaSensitiveWord model.ExaSensitiveWord
	_ = c.ShouldBindQuery(&exaSensitiveWord)
	if err, reexaSensitiveWord := service.GetExaSensitiveWord(exaSensitiveWord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexaSensitiveWord": reexaSensitiveWord}, c)
	}
}

// @Tags ExaSensitiveWord
// @Summary 分页获取ExaSensitiveWord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ExaSensitiveWordSearch true "分页获取ExaSensitiveWord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaSensitiveWord/getExaSensitiveWordList [get]
func GetExaSensitiveWordList(c *gin.Context) {
	var pageInfo request.ExaSensitiveWordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetExaSensitiveWordInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
