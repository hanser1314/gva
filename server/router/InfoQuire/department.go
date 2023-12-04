package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DepartmentRouter struct {
}

// InitDepartmentRouter 初始化 院系信息表 路由信息
func (s *DepartmentRouter) InitDepartmentRouter(Router *gin.RouterGroup) {
	departmentRouter := Router.Group("department").Use(middleware.OperationRecord())
	departmentRouterWithoutRecord := Router.Group("department")
	var departmentApi = v1.ApiGroupApp.InfoQuireApiGroup.DepartmentApi
	{
		departmentRouter.POST("createDepartment", departmentApi.CreateDepartment)   // 新建院系信息表
		departmentRouter.DELETE("deleteDepartment", departmentApi.DeleteDepartment) // 删除院系信息表
		departmentRouter.DELETE("deleteDepartmentByIds", departmentApi.DeleteDepartmentByIds) // 批量删除院系信息表
		departmentRouter.PUT("updateDepartment", departmentApi.UpdateDepartment)    // 更新院系信息表
	}
	{
		departmentRouterWithoutRecord.GET("findDepartment", departmentApi.FindDepartment)        // 根据ID获取院系信息表
		departmentRouterWithoutRecord.GET("getDepartmentList", departmentApi.GetDepartmentList)  // 获取院系信息表列表
	}
}
