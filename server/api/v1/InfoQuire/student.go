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

type StudentApi struct {
}

var studentService = service.ServiceGroupApp.InfoQuireServiceGroup.StudentService

// CreateStudent 创建student表
// @Tags Student
// @Summary 创建student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Student true "创建student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /student/createStudent [post]
func (studentApi *StudentApi) CreateStudent(c *gin.Context) {
	var student InfoQuire.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := studentService.CreateStudent(&student); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudent 删除student表
// @Tags Student
// @Summary 删除student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Student true "删除student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
func (studentApi *StudentApi) DeleteStudent(c *gin.Context) {
	var student InfoQuire.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := studentService.DeleteStudent(student); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentByIds 批量删除student表
// @Tags Student
// @Summary 批量删除student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /student/deleteStudentByIds [delete]
func (studentApi *StudentApi) DeleteStudentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := studentService.DeleteStudentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudent 更新student表
// @Tags Student
// @Summary 更新student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Student true "更新student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /student/updateStudent [put]
func (studentApi *StudentApi) UpdateStudent(c *gin.Context) {
	var student InfoQuire.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := studentService.UpdateStudent(student); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudent 用id查询student表
// @Tags Student
// @Summary 用id查询student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuire.Student true "用id查询student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /student/findStudent [get]
func (studentApi *StudentApi) FindStudent(c *gin.Context) {
	var student InfoQuire.Student
	err := c.ShouldBindQuery(&student)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if restudent, err := studentService.GetStudent(student.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restudent": restudent}, c)
	}
}

// GetStudentList 分页获取student表列表
// @Tags Student
// @Summary 分页获取student表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuireReq.StudentSearch true "分页获取student表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/getStudentList [get]
func (studentApi *StudentApi) GetStudentList(c *gin.Context) {
	var pageInfo InfoQuireReq.StudentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := studentService.GetStudentInfoList(pageInfo); err != nil {
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

func (studentApi *StudentApi) GetStudentByName(c *gin.Context) {
	var pageInfo InfoQuireReq.StudentSearchByName
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := studentService.GetStudentInfoListByName(pageInfo); err != nil {
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

func (studentApi *StudentApi) GetStudentBySno(c *gin.Context) {
	var pageInfo InfoQuireReq.StudentSearchBySno
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := studentService.GetStudentInfoListBySno(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
			//Page:     pageInfo.Page,
			//PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (studentApi *StudentApi) GetStudentByDept(c *gin.Context) {
	var pageInfo InfoQuireReq.StudentSearchByDept
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := studentService.GetStudentInfoListByDept(pageInfo); err != nil {
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
