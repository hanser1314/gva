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

type DepartmentApi struct {
}

var departmentService = service.ServiceGroupApp.InfoQuireServiceGroup.DepartmentService


// CreateDepartment 创建院系信息表
// @Tags Department
// @Summary 创建院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Department true "创建院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /department/createDepartment [post]
func (departmentApi *DepartmentApi) CreateDepartment(c *gin.Context) {
	var department InfoQuire.Department
	err := c.ShouldBindJSON(&department)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Dno":{utils.NotEmpty()},
        "Dname":{utils.NotEmpty()},
        "Dmanagerno":{utils.NotEmpty()},
    }
	if err := utils.Verify(department, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := departmentService.CreateDepartment(&department); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDepartment 删除院系信息表
// @Tags Department
// @Summary 删除院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Department true "删除院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /department/deleteDepartment [delete]
func (departmentApi *DepartmentApi) DeleteDepartment(c *gin.Context) {
	var department InfoQuire.Department
	err := c.ShouldBindJSON(&department)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := departmentService.DeleteDepartment(department); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDepartmentByIds 批量删除院系信息表
// @Tags Department
// @Summary 批量删除院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /department/deleteDepartmentByIds [delete]
func (departmentApi *DepartmentApi) DeleteDepartmentByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := departmentService.DeleteDepartmentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDepartment 更新院系信息表
// @Tags Department
// @Summary 更新院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body InfoQuire.Department true "更新院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /department/updateDepartment [put]
func (departmentApi *DepartmentApi) UpdateDepartment(c *gin.Context) {
	var department InfoQuire.Department
	err := c.ShouldBindJSON(&department)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Dno":{utils.NotEmpty()},
          "Dname":{utils.NotEmpty()},
          "Dmanagerno":{utils.NotEmpty()},
      }
    if err := utils.Verify(department, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := departmentService.UpdateDepartment(department); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDepartment 用id查询院系信息表
// @Tags Department
// @Summary 用id查询院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuire.Department true "用id查询院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /department/findDepartment [get]
func (departmentApi *DepartmentApi) FindDepartment(c *gin.Context) {
	var department InfoQuire.Department
	err := c.ShouldBindQuery(&department)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redepartment, err := departmentService.GetDepartment(department.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redepartment": redepartment}, c)
	}
}

// GetDepartmentList 分页获取院系信息表列表
// @Tags Department
// @Summary 分页获取院系信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query InfoQuireReq.DepartmentSearch true "分页获取院系信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /department/getDepartmentList [get]
func (departmentApi *DepartmentApi) GetDepartmentList(c *gin.Context) {
	var pageInfo InfoQuireReq.DepartmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := departmentService.GetDepartmentInfoList(pageInfo); err != nil {
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
