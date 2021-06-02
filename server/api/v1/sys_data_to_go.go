package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"text/template"
)

const im = `
package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"
)`

func SysBaseMenue2go(c *gin.Context) {
	tmpl, err := template.ParseFiles("./resource/template/data2go/model.tpl")
	if err != nil {
		global.GVA_LOG.Error("获取模板失败", zap.Any("err", err))
		response.Fail(c)
		return
	}
	GetModelText( "SysApi", tmpl, []model.SysApi{})
	GetModelText( "SysBaseMenu", tmpl, []model.SysBaseMenu{})
	GetModelText( "SysBaseMenuParameter", tmpl, []model.SysBaseMenuParameter{})
	GetModelText( "SysDictionary", tmpl, []model.SysDictionary{})
	GetModelText( "SysDictionaryDetail", tmpl, []model.SysDictionaryDetail{})
	response.OkWithDetailed("","数据已导出",c)
}
func GetModelText( filename string, tmpl *template.Template, mo interface{}) (err error) {
	var mapdatas []map[string]interface{}
	if models, ok := mo.([]model.SysApi); ok {
		if global.GVA_DB.Find(&models).RowsAffected== 0 {
			return nil
		}
		for _, value := range models {
			mapdata := utils.Struct2Map(value)
			mapdatas = append(mapdatas, mapdata)
		}
	}else if models, ok := mo.([]model.SysBaseMenu); ok {
		if global.GVA_DB.Find(&models).RowsAffected== 0 {
			return nil
		}
		for _, value := range models {
			mapdata := utils.Struct2Map(value)
			mapdatas = append(mapdatas, mapdata)
		}
	}else if models, ok := mo.([]model.SysBaseMenuParameter); ok {
		if global.GVA_DB.Find(&models).RowsAffected== 0 {
			return nil
		}
		for _, value := range models {
			mapdata := utils.Struct2Map(value)
			mapdatas = append(mapdatas, mapdata)
		}
	}else if models, ok := mo.([]model.SysDictionary); ok {
		if global.GVA_DB.Find(&models).RowsAffected== 0 {
			return nil
		}
		for _, value := range models {
			mapdata := utils.Struct2Map(value)
			mapdatas = append(mapdatas, mapdata)
		}
	}else if models, ok := mo.([]model.SysDictionaryDetail); ok {
		if global.GVA_DB.Find(&models).RowsAffected== 0 {
			return nil
		}
		for _, value := range models {
			mapdata := utils.Struct2Map(value)
			mapdatas = append(mapdatas, mapdata)
		}
	}
	mypath := "./source/gen_" + filename + ".go"
	if _, e := os.Stat(mypath); e == nil {
		os.Remove(mypath)
	}
	file, err := os.Create(mypath)
	if err != nil {
		return
	}
	defer file.Close()
	if _, err = fmt.Fprintln(file, im); err != nil {
		return
	}
	d := struct {
		Data    interface{}
		Varname interface{}
	}{mapdatas, filename}
	err = tmpl.Execute(file, d)
	if err != nil {
		return
	}
	return nil
}
