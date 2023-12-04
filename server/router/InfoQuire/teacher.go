package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeacherRouter struct {
}

// InitTeacherRouter 初始化 教师表 路由信息
func (s *TeacherRouter) InitTeacherRouter(Router *gin.RouterGroup) {
	teacherRouter := Router.Group("teacher").Use(middleware.OperationRecord())
	teacherRouterWithoutRecord := Router.Group("teacher")
	var teacherApi = v1.ApiGroupApp.InfoQuireApiGroup.TeacherApi
	{
		teacherRouter.POST("createTeacher", teacherApi.CreateTeacher)             // 新建教师表
		teacherRouter.DELETE("deleteTeacher", teacherApi.DeleteTeacher)           // 删除教师表
		teacherRouter.DELETE("deleteTeacherByIds", teacherApi.DeleteTeacherByIds) // 批量删除教师表
		teacherRouter.PUT("updateTeacher", teacherApi.UpdateTeacher)              // 更新教师表
	}
	{
		teacherRouterWithoutRecord.GET("findTeacher", teacherApi.FindTeacher)       // 根据ID获取教师表
		teacherRouterWithoutRecord.GET("getTeacherList", teacherApi.GetTeacherList) // 获取教师表列表

		teacherRouterWithoutRecord.GET("getTeacherByNameOrTno", teacherApi.GetTeacherByNameOrTno) //根据教师姓名或者教师号搜索
	}
}
