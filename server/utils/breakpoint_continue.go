package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成

const BreakpointDir = "./breakpointDir/"
const FinishDir = "./fileDir/"

//@author: [piexlmax](https://github.com/piexlmax)
//@function: BreakPointContinue
//@description: 断点续传
//@param: content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string
//@return: error, string

func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (error, string) {
	path := BreakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err, path
	}
	// filename 无实际用处，改为 fileMd5
	err, pathc := makeFileContent(content, fileMd5, path, contentNumber)
	return err, pathc

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CheckMd5
//@description: 检查Md5
//@param: content []byte, chunkMd5 string
//@return: CanUpload bool

func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true // 可以继续上传
	} else {
		return false // 切片不完整，废弃
	}
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: makeFileContent
//@description: 创建切片内容
//@param: content []byte, fileName string, FileDir string, contentNumber int
//@return: error, string

func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (error, string) {
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	if err != nil {
		return err, path
	} else {
		_, err = f.Write(content)
		if err != nil {
			return err, path
		}
	}
	defer f.Close()
	return nil, path
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: makeFileContent
//@description: 创建切片文件
//@param: fileName string, FileMd5 string
//@return: error, string

func MakeFile(fileName string, FileMd5 string) (error, string) {
	//rd, err := ioutil.ReadDir(BreakpointDir + FileMd5)
	//if err != nil {
	//	return err, FinishDir + FileMd5
	//}
	_ = os.MkdirAll(FinishDir, os.ModePerm)
	fd, err := os.OpenFile(FinishDir+FileMd5, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err, FinishDir + FileMd5
	}
	defer fd.Close()

	var contents string
	i := 1
	for {
		filepath := fmt.Sprintf("%s/%s/%s_%d", BreakpointDir, FileMd5, FileMd5, i)
		_, err := os.Stat(filepath)
		if err != nil {
			break
		}

		content, _ := ioutil.ReadFile(filepath)
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(FinishDir + fileName)
			return err, FinishDir + FileMd5
		}
		contents += string(content)

		i++
	}

	// 校验MD5
	md5 := MD5V([]byte(contents))
	if md5 != FileMd5 {
		_ = RemoveChunk(FileMd5)
		return errors.New("生成的文件MD5错误，请重新生成"), ""
	}

	return nil, FinishDir + FileMd5
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: RemoveChunk
//@description: 移除切片
//@param: FileMd5 string
//@return: error

func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(BreakpointDir + FileMd5)
	return err
}
