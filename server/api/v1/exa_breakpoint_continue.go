package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"strconv"
)

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))

	// 模拟测试失败，继续上传的情况
	//if chunkNumber == 2 {
	//	response.FailWithMessage("接收文件失败", c)
	//	return
	//}

	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.GVA_LOG.Error("文件读取失败!", zap.Any("err", err))
		response.FailWithMessage("文件读取失败", c)
		return
	}
	defer f.Close()
	cen, _ := ioutil.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		global.GVA_LOG.Error("检查md5失败!", zap.Any("err", err))
		response.FailWithMessage("检查md5失败", c)
		return
	}
	//err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	//if err != nil {
	//	global.GVA_LOG.Error("查找或创建记录失败!", zap.Any("err", err))
	//	response.FailWithMessage("查找或创建记录失败", c)
	//	return
	//}
	err, _ = utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("断点续传失败!", zap.Any("err", err))
		response.FailWithMessage("断点续传失败", c)
		return
	}

	//if err = service.CreateFileChunk(file.ID, pathc, chunkNumber); err != nil {
	//	global.GVA_LOG.Error("创建文件记录失败!", zap.Any("err", err))
	//	response.FailWithMessage("创建文件记录失败", c)
	//	return
	//}
	response.OkWithMessage("切片创建成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [get]
func FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))

	//err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)

	file := model.ExaFile{
		FileName:   fileName,
		FileMd5:    fileMd5,
		ChunkTotal: chunkTotal,
	}

	filepath := fmt.Sprintf("%s/%s", utils.FinishDir, fileMd5)
	stat, _ := os.Stat(filepath)
	if stat != nil {
		file.IsFinish = true
	} else {
		// 检查分片文件是否已经存在
		for i := 1; i <= chunkTotal; i++ {
			filepath := fmt.Sprintf("%s/%s/%s_%d", utils.BreakpointDir, fileMd5, fileMd5, i)
			stat, _ := os.Stat(filepath)
			if stat != nil {
				file.ExaFileChunk = append(file.ExaFileChunk, model.ExaFileChunk{
					FileChunkNumber: i,
					FileChunkPath:   filepath,
				})
			}
		}

		if len(file.ExaFileChunk) == chunkTotal {
			file.IsFinish = true
		}
	}

	response.OkWithDetailed(response.FileResponse{File: file}, "ok", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 创建文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/breakpointContinueFinish [put]
func BreakpointContinueFinish(c *gin.Context) {
	/* TODO 云端存储时，将本地文件替换为一个内容为空的文件，名称为文件 MD5 即可 */
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	err, filePath := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("文件创建失败!", zap.Any("err", err))
		response.FailWithDetailed(response.FilePathResponse{FilePath: filePath}, "文件创建失败", c)
	} else {
		response.OkWithDetailed(response.FilePathResponse{FilePath: filePath}, "文件创建成功", c)
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"缓存切片删除成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func RemoveChunk(c *gin.Context) {
	/* TODO 云端存储时，将本地文件替换为一个内容为空的文件，名称为文件 MD5 即可 */
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath := c.Query("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = service.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		global.GVA_LOG.Error("缓存切片删除失败!", zap.Any("err", err))
		response.FailWithDetailed(response.FilePathResponse{FilePath: filePath}, "缓存切片删除失败", c)
	} else {
		response.OkWithDetailed(response.FilePathResponse{FilePath: filePath}, "缓存切片删除成功", c)
	}
}
