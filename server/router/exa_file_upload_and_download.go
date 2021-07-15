package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	FileUploadAndDownloadRouter := Router.Group("fileUploadAndDownload")
	{
		FileUploadAndDownloadRouter.POST("/upload", v1.UploadFile)                                // 上传文件
		FileUploadAndDownloadRouter.POST("/getFileList", v1.GetFileList)                          // 获取上传文件列表
		FileUploadAndDownloadRouter.POST("/deleteFile", v1.DeleteFile)                            // 删除指定文件
		FileUploadAndDownloadRouter.POST("/breakpointContinue", v1.BreakpointContinue)            // 断点续传
		FileUploadAndDownloadRouter.GET("/findFile", v1.FindFile)                                 // 查询当前文件成功的切片
		FileUploadAndDownloadRouter.PUT("/breakpointContinueFinish", v1.BreakpointContinueFinish) // 查询当前文件成功的切片
		FileUploadAndDownloadRouter.DELETE("/removeChunk", v1.RemoveChunk)                        // 查询当前文件成功的切片
	}
}
