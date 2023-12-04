package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	InfoQuireReq "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SelectedCourseApi struct {
}

var selectedCourseService = service.ServiceGroupApp.InfoQuireServiceGroup.SelectedCourseService

// CreateSelectedCourse 创建已选课程表
// @Tags SelectedCourse
// @Summary 创建已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.SelectedCourse true "创建已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /selectedCourse/createSelectedCourse [post]
func (selectedCourseApi *SelectedCourseApi) CreateSelectedCourse(c *gin.Context) {
	var selectedCourse InfoQuire.SelectedCourse
	err := c.ShouldBindJSON(&selectedCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sno":   {utils.NotEmpty()},
		"Cno":   {utils.NotEmpty()},
		"Tno":   {utils.NotEmpty()},
		"Grade": {utils.NotEmpty()},
	}
	if err := utils.Verify(selectedCourse, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := selectedCourseService.CreateSelectedCourse(&selectedCourse); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSelectedCourse 删除已选课程表
// @Tags SelectedCourse
// @Summary 删除已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.SelectedCourse true "删除已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /selectedCourse/deleteSelectedCourse [delete]
func (selectedCourseApi *SelectedCourseApi) DeleteSelectedCourse(c *gin.Context) {
	var selectedCourse InfoQuire.SelectedCourse
	err := c.ShouldBindJSON(&selectedCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := selectedCourseService.DeleteSelectedCourse(selectedCourse); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSelectedCourseByIds 批量删除已选课程表
// @Tags SelectedCourse
// @Summary 批量删除已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /selectedCourse/deleteSelectedCourseByIds [delete]
func (selectedCourseApi *SelectedCourseApi) DeleteSelectedCourseByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := selectedCourseService.DeleteSelectedCourseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSelectedCourse 更新已选课程表
// @Tags SelectedCourse
// @Summary 更新已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.SelectedCourse true "更新已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /selectedCourse/updateSelectedCourse [put]
func (selectedCourseApi *SelectedCourseApi) UpdateSelectedCourse(c *gin.Context) {
	var selectedCourse InfoQuire.SelectedCourse
	err := c.ShouldBindJSON(&selectedCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Sno":   {utils.NotEmpty()},
		"Cno":   {utils.NotEmpty()},
		"Tno":   {utils.NotEmpty()},
		"Grade": {utils.NotEmpty()},
	}
	if err := utils.Verify(selectedCourse, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := selectedCourseService.UpdateSelectedCourse(selectedCourse); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSelectedCourse 用id查询已选课程表
// @Tags SelectedCourse
// @Summary 用id查询已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuire.SelectedCourse true "用id查询已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /selectedCourse/findSelectedCourse [get]
func (selectedCourseApi *SelectedCourseApi) FindSelectedCourse(c *gin.Context) {
	var selectedCourse InfoQuire.SelectedCourse
	err := c.ShouldBindQuery(&selectedCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reselectedCourse, err := selectedCourseService.GetSelectedCourse(selectedCourse.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reselectedCourse": reselectedCourse}, c)
	}
}

// GetSelectedCourseList 分页获取已选课程表列表
// @Tags SelectedCourse
// @Summary 分页获取已选课程表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuireReq.SelectedCourseSearch true "分页获取已选课程表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /selectedCourse/getSelectedCourseList [get]
func (selectedCourseApi *SelectedCourseApi) GetSelectedCourseList(c *gin.Context) {
	var pageInfo InfoQuireReq.SelectedCourseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := selectedCourseService.GetSelectedCourseInfoList(pageInfo); err != nil {
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
