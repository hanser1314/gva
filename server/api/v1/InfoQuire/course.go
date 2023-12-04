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

type CourseApi struct {
}

var courseService = service.ServiceGroupApp.InfoQuireServiceGroup.CourseService

// CreateCourse 创建课程表
// @Tags Course
// @Summary 创建课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Course true "创建课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /course/createCourse [post]
func (courseApi *CourseApi) CreateCourse(c *gin.Context) {
	var course InfoQuire.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseService.CreateCourse(&course); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCourse 删除课程表
// @Tags Course
// @Summary 删除课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Course true "删除课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /course/deleteCourse [delete]
func (courseApi *CourseApi) DeleteCourse(c *gin.Context) {
	var course InfoQuire.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseService.DeleteCourse(course); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCourseByIds 批量删除课程表
// @Tags Course
// @Summary 批量删除课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /course/deleteCourseByIds [delete]
func (courseApi *CourseApi) DeleteCourseByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseService.DeleteCourseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCourse 更新课程表
// @Tags Course
// @Summary 更新课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Course true "更新课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /course/updateCourse [put]
func (courseApi *CourseApi) UpdateCourse(c *gin.Context) {
	var course InfoQuire.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := courseService.UpdateCourse(course); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCourse 用id查询课程表
// @Tags Course
// @Summary 用id查询课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuire.Course true "用id查询课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /course/findCourse [get]
func (courseApi *CourseApi) FindCourse(c *gin.Context) {
	var course InfoQuire.Course
	err := c.ShouldBindQuery(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recourse, err := courseService.GetCourse(course.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recourse": recourse}, c)
	}
}

// GetCourseList 分页获取课程表列表
// @Tags Course
// @Summary 分页获取课程表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuireReq.CourseSearch true "分页获取课程表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /course/getCourseList [get]
func (courseApi *CourseApi) GetCourseList(c *gin.Context) {
	var pageInfo InfoQuireReq.CourseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := courseService.GetCourseInfoList(pageInfo); err != nil {
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
