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

type TeacherApi struct {
}

var teacherService = service.ServiceGroupApp.InfoQuireServiceGroup.TeacherService

// CreateTeacher 创建教师表
// @Tags Teacher
// @Summary 创建教师表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Teacher true "创建教师表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /teacher/createTeacher [post]
func (teacherApi *TeacherApi) CreateTeacher(c *gin.Context) {
	var teacher InfoQuire.Teacher
	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Tno":   {utils.NotEmpty()},
		"Tname": {utils.NotEmpty()},
		"Tsex":  {utils.NotEmpty()},
		"Tage":  {utils.NotEmpty()},
		"Teb":   {utils.NotEmpty()},
		"Cno1":  {utils.NotEmpty()},
	}
	if err := utils.Verify(teacher, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teacherService.CreateTeacher(&teacher); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeacher 删除教师表
// @Tags Teacher
// @Summary 删除教师表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Teacher true "删除教师表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacher/deleteTeacher [delete]
func (teacherApi *TeacherApi) DeleteTeacher(c *gin.Context) {
	var teacher InfoQuire.Teacher
	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teacherService.DeleteTeacher(teacher); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeacherByIds 批量删除教师表
// @Tags Teacher
// @Summary 批量删除教师表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除教师表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teacher/deleteTeacherByIds [delete]
func (teacherApi *TeacherApi) DeleteTeacherByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teacherService.DeleteTeacherByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeacher 更新教师表
// @Tags Teacher
// @Summary 更新教师表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Teacher true "更新教师表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teacher/updateTeacher [put]
func (teacherApi *TeacherApi) UpdateTeacher(c *gin.Context) {
	var teacher InfoQuire.Teacher
	err := c.ShouldBindJSON(&teacher)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Tno":   {utils.NotEmpty()},
		"Tname": {utils.NotEmpty()},
		"Tsex":  {utils.NotEmpty()},
		"Tage":  {utils.NotEmpty()},
		"Teb":   {utils.NotEmpty()},
		"Cno1":  {utils.NotEmpty()},
	}
	if err := utils.Verify(teacher, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := teacherService.UpdateTeacher(teacher); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeacher 用id查询教师表
// @Tags Teacher
// @Summary 用id查询教师表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuire.Teacher true "用id查询教师表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teacher/findTeacher [get]
func (teacherApi *TeacherApi) FindTeacher(c *gin.Context) {
	var teacher InfoQuire.Teacher
	err := c.ShouldBindQuery(&teacher)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reteacher, err := teacherService.GetTeacher(teacher.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteacher": reteacher}, c)
	}
}

// GetTeacherList 分页获取教师表列表
// @Tags Teacher
// @Summary 分页获取教师表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuireReq.TeacherSearch true "分页获取教师表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacher/getTeacherList [get]
func (teacherApi *TeacherApi) GetTeacherList(c *gin.Context) {
	var pageInfo InfoQuireReq.TeacherSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teacherService.GetTeacherInfoList(pageInfo); err != nil {
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

func (teacherApi *TeacherApi) GetTeacherByNameOrTno(c *gin.Context) {
	var pageInfo InfoQuireReq.TeacherSearchByNameOrTno
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := teacherService.GetTeacherInfoListByNameOrTno(pageInfo); err != nil {
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
