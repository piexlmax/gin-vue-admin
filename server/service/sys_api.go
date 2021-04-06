package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

func CreateApi(api model.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func DeleteApi(api model.SysApi) (err error) {
	err = global.GVA_DB.Delete(&api).Error
	ClearCasbin(1, api.Path, api.Method)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func GetAPIInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysApi{})
	var apiList []model.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func GetAllApis() (err error, apis []model.SysApi) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func GetApiById(id float64) (err error, api model.SysApi) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func UpdateApi(api model.SysApi) (err error) {
	var oldA model.SysApi
	err = global.GVA_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(&api).Error
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func DeleteApisByIds(ids request.IdsReq) error {
	return global.GVA_DB.Delete(&[]model.SysApi{}, "id in ?", ids.Ids).Error
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: AutoRegisterRouter
//@description: 自动把路由注册到sys_apis表
//@param: engine *gin.Engine
//@return: error

func AutoRegisterRouter(engine *gin.Engine) error {
	routes := engine.Routes()
	for i := 0; i < len(routes); i++ {
		path := routes[i].Path
		method := routes[i].Method
		entity := model.SysApi{Path: path, Method: method}
		switch { // 排除 不需要鉴权的路由 swagger, uploads/file 静态文件 表单生成器的静态文件托管路由
		case path == "/base/captcha" && method == "POST":
			continue
		case path == "/init/initdb" && method == "POST":
			continue
		case path == "/init/checkdb" && method == "POST":
			continue
		case path == "/swagger/*any" && method == "GET":
			continue
		case path == "/uploads/file/*filepath" && (method == "HEAD" || method == "GET"):
			continue
		case path == "/form-generator/*filepath" && (method == "HEAD" || method == "GET"):
			continue
		}
		if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", path, method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			//global.GVA_LOG.Info("sys_apis 表已经存在此记录!", zap.String("Path", path), zap.String("method", method))
			continue
		}
		if strings.Contains(path, "create") {
			if list := strings.Split(path, "create"); len(list) == 2 {
				entity.Description = "新建" + list[1]
			}
		} else if strings.Contains(path, "find") {
			if list := strings.Split(path, "find"); len(list) == 2 {
				entity.Description = "根据ID获取" + list[1]
			}
		} else if strings.Contains(path, "update") {
			if list := strings.Split(path, "update"); len(list) == 2 {
				entity.Description = "更新" + list[1]
			}
		} else if strings.Contains(path, "delete") {
			if list := strings.Split(path, "delete"); len(list) == 2 {
				if strings.Contains(list[1], "ByIds") {
					if deletes := strings.Split(list[1], "ByIds"); len(deletes) == 2 {
						entity.Description = "批量删除" + deletes[0]
					}
				} else {
					entity.Description = "删除" + list[1]
				}
			}
		} else if strings.Contains(path, "get") {
			if list := strings.Split(path, "get"); len(list) == 2 {
				if strings.Contains(list[1], "List") {
					if lists := strings.Split(list[1], "List"); len(lists) == 2 {
						entity.Description = "获取" + lists[0] + "列表"
					}
				}
				entity.Description = "获取" + list[1]
			}
		} else {
			entity.Description = path // 如果不符合代码生成器规则, api中文描述则为路由路径
		}
		if err := global.GVA_DB.Create(&entity).Error; err != nil { // 创建api记录
			global.GVA_LOG.Info("插入 sys_apis 表记录失败!", zap.String("Path", path), zap.String("method", method))
			return err
		}
	}
	return nil
}
