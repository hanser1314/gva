package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    InfoQuireReq "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OptCourseApi struct {
}

var optCourseService = service.ServiceGroupApp.InfoQuireServiceGroup.OptCourseService


// CreateOptCourse 创建选课表
// @Tags OptCourse
// @Summary 创建选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.OptCourse true "创建选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /optCourse/createOptCourse [post]
func (optCourseApi *OptCourseApi) CreateOptCourse(c *gin.Context) {
	var optCourse InfoQuire.OptCourse
	err := c.ShouldBindJSON(&optCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Cno":{utils.NotEmpty()},
        "Cname":{utils.NotEmpty()},
        "Cpno":{utils.NotEmpty()},
        "Ccredit":{utils.NotEmpty()},
    }
	if err := utils.Verify(optCourse, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := optCourseService.CreateOptCourse(&optCourse); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOptCourse 删除选课表
// @Tags OptCourse
// @Summary 删除选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.OptCourse true "删除选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /optCourse/deleteOptCourse [delete]
func (optCourseApi *OptCourseApi) DeleteOptCourse(c *gin.Context) {
	var optCourse InfoQuire.OptCourse
	err := c.ShouldBindJSON(&optCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := optCourseService.DeleteOptCourse(optCourse); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOptCourseByIds 批量删除选课表
// @Tags OptCourse
// @Summary 批量删除选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /optCourse/deleteOptCourseByIds [delete]
func (optCourseApi *OptCourseApi) DeleteOptCourseByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := optCourseService.DeleteOptCourseByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOptCourse 更新选课表
// @Tags OptCourse
// @Summary 更新选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.OptCourse true "更新选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /optCourse/updateOptCourse [put]
func (optCourseApi *OptCourseApi) UpdateOptCourse(c *gin.Context) {
	var optCourse InfoQuire.OptCourse
	err := c.ShouldBindJSON(&optCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Cno":{utils.NotEmpty()},
          "Cname":{utils.NotEmpty()},
          "Cpno":{utils.NotEmpty()},
          "Ccredit":{utils.NotEmpty()},
      }
    if err := utils.Verify(optCourse, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := optCourseService.UpdateOptCourse(optCourse); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOptCourse 用id查询选课表
// @Tags OptCourse
// @Summary 用id查询选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuire.OptCourse true "用id查询选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /optCourse/findOptCourse [get]
func (optCourseApi *OptCourseApi) FindOptCourse(c *gin.Context) {
	var optCourse InfoQuire.OptCourse
	err := c.ShouldBindQuery(&optCourse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reoptCourse, err := optCourseService.GetOptCourse(optCourse.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoptCourse": reoptCourse}, c)
	}
}

// GetOptCourseList 分页获取选课表列表
// @Tags OptCourse
// @Summary 分页获取选课表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuireReq.OptCourseSearch true "分页获取选课表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /optCourse/getOptCourseList [get]
func (optCourseApi *OptCourseApi) GetOptCourseList(c *gin.Context) {
	var pageInfo InfoQuireReq.OptCourseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := optCourseService.GetOptCourseInfoList(pageInfo); err != nil {
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
