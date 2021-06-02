package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitData2GoRouter(Router *gin.RouterGroup) {
	Data2GoRouter := Router.Group("d2g")
	{
		//AutoCodeRouter.POST("preview", v1.PreviewTemp)   // 获取自动创建代码预览
		//AutoCodeRouter.POST("createTemp", v1.CreateTemp) // 创建自动化代码
		//AutoCodeRouter.GET("getTables", v1.GetTables)    // 获取对应数据库的表
		//AutoCodeRouter.GET("getDB", v1.GetDB)            // 获取数据库
		//AutoCodeRouter.GET("getColumn", v1.GetColumn)    // 获取指定表所有字段信息
		Data2GoRouter.GET("sysbasemenu2go", v1.SysBaseMenue2go)
	}
}
