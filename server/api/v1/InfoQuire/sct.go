package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	InfoQuireReq "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SctApi struct {
}

var sctService = service.ServiceGroupApp.InfoQuireServiceGroup.SctService

// CreateSct 创建选课表
// @Tags Sct
// @Summary 创建选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Sct true "创建选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sct/createSct [post]
func (sctApi *SctApi) CreateSct(c *gin.Context) {
	var sct InfoQuire.Sct
	err := c.ShouldBindJSON(&sct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sctService.CreateSct(&sct); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSct 删除选课表
// @Tags Sct
// @Summary 删除选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Sct true "删除选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sct/deleteSct [delete]
func (sctApi *SctApi) DeleteSct(c *gin.Context) {
	var sct InfoQuire.Sct
	err := c.ShouldBindJSON(&sct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sctService.DeleteSct(sct); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSctByIds 批量删除选课表
// @Tags Sct
// @Summary 批量删除选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sct/deleteSctByIds [delete]
func (sctApi *SctApi) DeleteSctByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sctService.DeleteSctByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSct 更新选课表
// @Tags Sct
// @Summary 更新选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Sct true "更新选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sct/updateSct [put]
func (sctApi *SctApi) UpdateSct(c *gin.Context) {
	var sct InfoQuire.Sct
	err := c.ShouldBindJSON(&sct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sctService.UpdateSct(sct); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSct 用id查询选课表
// @Tags Sct
// @Summary 用id查询选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuire.Sct true "用id查询选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sct/findSct [get]
func (sctApi *SctApi) FindSct(c *gin.Context) {
	var sct InfoQuire.Sct
	err := c.ShouldBindQuery(&sct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resct, err := sctService.GetSct(sct.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resct": resct}, c)
	}
}

// GetSctList 分页获取选课表列表
// @Tags Sct
// @Summary 分页获取选课表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuireReq.SctSearch true "分页获取选课表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sct/getSctList [get]
func (sctApi *SctApi) GetSctList(c *gin.Context) {
	var pageInfo InfoQuireReq.SctSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sctService.GetSctInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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

func (sctApi *SctApi) GetSctListBySno(c *gin.Context) {
	var pageInfo InfoQuireReq.SctSearchBySno
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sctService.GetSctInfoListBySno(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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

func (sctApi *SctApi) GetSctListBySnoAndCno(c *gin.Context) {
	var pageInfo InfoQuireReq.SctBySnoAndCno
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sctService.GetSctInfoListBySnoAndCno(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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

func (sctApi *SctApi) GetSctListByTno(c *gin.Context) {
	var pageInfo InfoQuireReq.SctByTno
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sctService.GetSctInfoListByTno(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
