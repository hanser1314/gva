package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CourseRouter struct {
}

// InitCourseRouter 初始化 课程表 路由信息
func (s *CourseRouter) InitCourseRouter(Router *gin.RouterGroup) {
	courseRouter := Router.Group("course").Use(middleware.OperationRecord())
	courseRouterWithoutRecord := Router.Group("course")
	var courseApi = v1.ApiGroupApp.InfoQuireApiGroup.CourseApi
	{
		courseRouter.POST("createCourse", courseApi.CreateCourse)   // 新建课程表
		courseRouter.DELETE("deleteCourse", courseApi.DeleteCourse) // 删除课程表
		courseRouter.DELETE("deleteCourseByIds", courseApi.DeleteCourseByIds) // 批量删除课程表
		courseRouter.PUT("updateCourse", courseApi.UpdateCourse)    // 更新课程表
	}
	{
		courseRouterWithoutRecord.GET("findCourse", courseApi.FindCourse)        // 根据ID获取课程表
		courseRouterWithoutRecord.GET("getCourseList", courseApi.GetCourseList)  // 获取课程表列表
	}
}
