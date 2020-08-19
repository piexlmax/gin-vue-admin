package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		UserRouter.POST("changePassword", v1.ChangePassword)         // 修改密码
		UserRouter.POST("uploadHeaderImg", v1.UploadHeaderImg)       // 上传头像
		UserRouter.POST("getUserList", v1.GetUserList)               // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority)     // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteUser)               // 删除用户
		UserRouter.POST("checkusermfaisopen", v1.CheckUserMFAIsOpen) // 获取用户是否开启多因子认证
		UserRouter.POST("getusermfainfo", v1.GetUserMFAInfo)         // 获取用户mfa信息
		UserRouter.POST("bindmfa", v1.BindMfa)                       //绑定多因子认证
		UserRouter.POST("unbindmfa", v1.UnBindMfa)                   //解除多因子认证
	}
}
