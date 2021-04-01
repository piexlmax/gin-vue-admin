package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitExaSensitiveWordRouter(Router *gin.RouterGroup) {
	ExaSensitiveWordRouter := Router.Group("exaSensitiveWord").Use(middleware.OperationRecord())
	{
		ExaSensitiveWordRouter.POST("createExaSensitiveWord", v1.CreateExaSensitiveWord)   // 新建ExaSensitiveWord
		ExaSensitiveWordRouter.DELETE("deleteExaSensitiveWord", v1.DeleteExaSensitiveWord) // 删除ExaSensitiveWord
		ExaSensitiveWordRouter.DELETE("deleteExaSensitiveWordByIds", v1.DeleteExaSensitiveWordByIds) // 批量删除ExaSensitiveWord
		ExaSensitiveWordRouter.PUT("updateExaSensitiveWord", v1.UpdateExaSensitiveWord)    // 更新ExaSensitiveWord
		ExaSensitiveWordRouter.GET("findExaSensitiveWord", v1.FindExaSensitiveWord)        // 根据ID获取ExaSensitiveWord
		ExaSensitiveWordRouter.GET("getExaSensitiveWordList", v1.GetExaSensitiveWordList)  // 获取ExaSensitiveWord列表
	}
}
