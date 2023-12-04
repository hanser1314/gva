package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SelectedCourseRouter struct {
}

// InitSelectedCourseRouter 初始化 已选课程表 路由信息
func (s *SelectedCourseRouter) InitSelectedCourseRouter(Router *gin.RouterGroup) {
	selectedCourseRouter := Router.Group("selectedCourse").Use(middleware.OperationRecord())
	selectedCourseRouterWithoutRecord := Router.Group("selectedCourse")
	var selectedCourseApi = v1.ApiGroupApp.InfoQuireApiGroup.SelectedCourseApi
	{
		selectedCourseRouter.POST("createSelectedCourse", selectedCourseApi.CreateSelectedCourse)   // 新建已选课程表
		selectedCourseRouter.DELETE("deleteSelectedCourse", selectedCourseApi.DeleteSelectedCourse) // 删除已选课程表
		selectedCourseRouter.DELETE("deleteSelectedCourseByIds", selectedCourseApi.DeleteSelectedCourseByIds) // 批量删除已选课程表
		selectedCourseRouter.PUT("updateSelectedCourse", selectedCourseApi.UpdateSelectedCourse)    // 更新已选课程表
	}
	{
		selectedCourseRouterWithoutRecord.GET("findSelectedCourse", selectedCourseApi.FindSelectedCourse)        // 根据ID获取已选课程表
		selectedCourseRouterWithoutRecord.GET("getSelectedCourseList", selectedCourseApi.GetSelectedCourseList)  // 获取已选课程表列表
	}
}
