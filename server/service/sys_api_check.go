package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

func CheckApis(Router *gin.Engine) {
	if global.GVA_DB == nil {
		return
	}
	global.GVA_DB.Model(model.SysApi{}).Where("valid <> false").UpdateColumn("valid", false)
	rt := Router.Routes()
	for _, v := range rt {
		urls := strings.Split(v.Path, "/")
		group := urls[len(urls)-2]
		Description := urls[len(urls)-1]
		//fmt.Println(urls[len(urls)-1])
		var profile model.SysApi
		//global.GVA_DB.Deb
		//errs := global.GVA_DB.Debug().Model(model.SysApi{}).Where(model.SysApi{Path: v.Path}).Where("description is null").Assign(model.SysApi{Valid: true, ApiGroup: group,Method: v.Method,Description: Description}).FirstOrCreate(&profile).Error
		//if errs!=nil {
		//	global.GVA_LOG.Error("API判别1",zap.Any("err",errs))
		//}
		errs := global.GVA_DB.Model(model.SysApi{}).Where(model.SysApi{Path: v.Path}).Assign(model.SysApi{Valid: true, ApiGroup: group, Method: v.Method}).FirstOrCreate(&profile).Error
		if errs != nil {
			global.GVA_LOG.Error("API判别2", zap.Any("err", errs))
		}
		var sapi model.SysApi
		if e := global.GVA_DB.Model(model.SysApi{}).Where(model.SysApi{Path: v.Path}).Where("description = ''").First(&sapi).Error; e == nil {
			sapi.Description = Description
			global.GVA_DB.Save(sapi)
		}

	}
	var effective int64
	var all int64
	global.GVA_DB.Model(model.SysApi{}).Where(model.SysApi{Valid: true}).Count(&effective)
	global.GVA_DB.Model(model.SysApi{}).Where("1=1").Count(&all)
	global.GVA_LOG.Info(fmt.Sprintf("完成API校验，有效API/总API：%d/%d", effective, all))
}

