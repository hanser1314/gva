package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SctRouter struct {
}

// InitSctRouter 初始化 选课表 路由信息
func (s *SctRouter) InitSctRouter(Router *gin.RouterGroup) {
	sctRouter := Router.Group("sct").Use(middleware.OperationRecord())
	sctRouterWithoutRecord := Router.Group("sct")
	var sctApi = v1.ApiGroupApp.InfoQuireApiGroup.SctApi
	{
		sctRouter.POST("createSct", sctApi.CreateSct)             // 新建选课表
		sctRouter.DELETE("deleteSct", sctApi.DeleteSct)           // 删除选课表
		sctRouter.DELETE("deleteSctByIds", sctApi.DeleteSctByIds) // 批量删除选课表
		sctRouter.PUT("updateSct", sctApi.UpdateSct)              // 更新选课表
	}
	{
		sctRouterWithoutRecord.GET("findSct", sctApi.FindSct)                             // 根据ID获取选课表
		sctRouterWithoutRecord.GET("getSctList", sctApi.GetSctList)                       // 获取选课表列表
		sctRouterWithoutRecord.GET("getSctListBySno", sctApi.GetSctListBySno)             // 获取选课表列表
		sctRouterWithoutRecord.GET("getSctListBySnoAndCno", sctApi.GetSctListBySnoAndCno) // 更新选课表
		sctRouterWithoutRecord.GET("getSctListByTno", sctApi.GetSctListByTno)             // 更新选课表

	}
}
