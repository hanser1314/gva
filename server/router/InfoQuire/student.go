package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentRouter struct {
}

// InitStudentRouter 初始化 student表 路由信息
func (s *StudentRouter) InitStudentRouter(Router *gin.RouterGroup) {
	studentRouter := Router.Group("student").Use(middleware.OperationRecord())
	studentRouterWithoutRecord := Router.Group("student")
	var studentApi = v1.ApiGroupApp.InfoQuireApiGroup.StudentApi
	{
		studentRouter.POST("createStudent", studentApi.CreateStudent)             // 新建student表
		studentRouter.DELETE("deleteStudent", studentApi.DeleteStudent)           // 删除student表
		studentRouter.DELETE("deleteStudentByIds", studentApi.DeleteStudentByIds) // 批量删除student表
		studentRouter.PUT("updateStudent", studentApi.UpdateStudent)              // 更新student表
	}
	{
		studentRouterWithoutRecord.GET("findStudent", studentApi.FindStudent)           // 根据ID获取student表
		studentRouterWithoutRecord.GET("getStudentList", studentApi.GetStudentList)     // 获取student表列表
		studentRouterWithoutRecord.GET("getStudentByName", studentApi.GetStudentByName) //根据姓名和年龄获取student表
		studentRouterWithoutRecord.GET("getStudentBySno", studentApi.GetStudentBySno)   //根据学号获取student表
		studentRouterWithoutRecord.GET("getStudentByDept", studentApi.GetStudentByDept) //根据学号获取student表
	}
}
