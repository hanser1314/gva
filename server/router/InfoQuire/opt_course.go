package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OptCourseRouter struct {
}

// InitOptCourseRouter 初始化 选课表 路由信息
func (s *OptCourseRouter) InitOptCourseRouter(Router *gin.RouterGroup) {
	optCourseRouter := Router.Group("optCourse").Use(middleware.OperationRecord())
	optCourseRouterWithoutRecord := Router.Group("optCourse")
	var optCourseApi = v1.ApiGroupApp.InfoQuireApiGroup.OptCourseApi
	{
		optCourseRouter.POST("createOptCourse", optCourseApi.CreateOptCourse)   // 新建选课表
		optCourseRouter.DELETE("deleteOptCourse", optCourseApi.DeleteOptCourse) // 删除选课表
		optCourseRouter.DELETE("deleteOptCourseByIds", optCourseApi.DeleteOptCourseByIds) // 批量删除选课表
		optCourseRouter.PUT("updateOptCourse", optCourseApi.UpdateOptCourse)    // 更新选课表
	}
	{
		optCourseRouterWithoutRecord.GET("findOptCourse", optCourseApi.FindOptCourse)        // 根据ID获取选课表
		optCourseRouterWithoutRecord.GET("getOptCourseList", optCourseApi.GetOptCourseList)  // 获取选课表列表
	}
}
